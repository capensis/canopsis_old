package saml

import (
	"bytes"
	"cmp"
	"compress/flate"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/auth"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/common"
	apisecurity "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/api/security"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/config"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/security"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/security/model"
	"git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/security/roleprovider"
	libsession "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/security/session"
	"github.com/beevik/etree"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	saml2 "github.com/russellhaering/gosaml2"
	samltypes "github.com/russellhaering/gosaml2/types"
	dsig "github.com/russellhaering/goxmldsig"
)

const MetadataReqTimeout = time.Second * 15

const (
	BindingRedirect = "redirect"
	BindingPOST     = "post"
)

type Provider interface {
	SamlMetadataHandler() gin.HandlerFunc
	SamlAuthHandler() gin.HandlerFunc
	SamlAcsHandler() gin.HandlerFunc
	SamlSloHandler() gin.HandlerFunc
}

type provider struct {
	userProvider       security.UserProvider
	roleProvider       security.RoleProvider
	sessionStore       libsession.Store
	enforcer           security.Enforcer
	config             security.SamlConfig
	tokenService       apisecurity.TokenService
	maintenanceAdapter config.MaintenanceAdapter
	logger             zerolog.Logger
	keyStore           dsig.X509KeyStore
	acsUrl             string
	sloUrl             string
	metadataUrl        string

	samlMx           sync.Mutex
	samlSP           *saml2.SAMLServiceProvider
	samlSPValidUntil time.Time
	samlSPAvailable  bool
}

func NewProvider(
	ctx context.Context,
	userProvider security.UserProvider,
	roleValidator security.RoleProvider,
	sessionStore libsession.Store,
	enforcer security.Enforcer,
	config security.SamlConfig,
	tokenService apisecurity.TokenService,
	maintenanceAdapter config.MaintenanceAdapter,
	logger zerolog.Logger,
) (Provider, error) {
	if config.IdPMetadataUrl != "" && config.IdPMetadataXml != "" {
		return nil, errors.New("should provide only IdP metadata url or xml, not both")
	}

	if config.CanopsisSSOBinding != BindingRedirect && config.CanopsisSSOBinding != BindingPOST {
		return nil, errors.New("wrong canopsis_sso_binding value, should be post or redirect")
	}

	if config.CanopsisACSBinding != BindingRedirect && config.CanopsisACSBinding != BindingPOST {
		return nil, errors.New("wrong canopsis_acs_binding value, should be post or redirect")
	}

	parsedSamlURL, err := url.Parse(config.CanopsisSamlUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid canopsis_saml_url: %w", err)
	}

	keyPair, err := tls.LoadX509KeyPair(config.X509Cert, config.X509Key)
	if err != nil {
		return nil, fmt.Errorf("failed to load x509 key pair: %w", err)
	}

	p := &provider{
		userProvider:       userProvider,
		roleProvider:       roleValidator,
		sessionStore:       sessionStore,
		enforcer:           enforcer,
		config:             config,
		tokenService:       tokenService,
		maintenanceAdapter: maintenanceAdapter,
		logger:             logger,
		keyStore:           dsig.TLSCertKeyStore(keyPair),
		acsUrl:             parsedSamlURL.JoinPath("acs").String(),
		sloUrl:             parsedSamlURL.JoinPath("slo").String(),
		metadataUrl:        parsedSamlURL.JoinPath("metadata").String(),
	}

	if config.IdPMetadataXml != "" {
		err = p.loadXmlMetadata()
		if err != nil {
			return nil, fmt.Errorf("failed to load IdP metadata by xml: %w", err)
		}
	} else {
		err = p.loadUrlMetadata(ctx)
		if err != nil {
			logger.Warn().Err(err).Msg("failed to load IdP metadata by url")
		}
	}

	return p, nil
}

func (p *provider) loadXmlMetadata() error {
	f, err := os.Open(p.config.IdPMetadataXml)
	if err != nil {
		return fmt.Errorf("failed to open IdP metadata xml file: %w", err)
	}

	defer f.Close()

	idpMetadata := samltypes.EntityDescriptor{}
	err = xml.NewDecoder(f).Decode(&idpMetadata)
	if err != nil {
		return fmt.Errorf("failed to unmarshal IdP metadata xml file: %w", err)
	}

	ssoLocation, sloLocation, certStore, err := processIdPMetadata(idpMetadata)
	if err != nil {
		return fmt.Errorf("failed to process IdP metadata xml file: %w", err)
	}

	p.samlSP = &saml2.SAMLServiceProvider{
		IdentityProviderSSOURL:         ssoLocation,
		IdentityProviderSLOURL:         sloLocation,
		IdentityProviderIssuer:         idpMetadata.EntityID,
		IDPCertificateStore:            &certStore,
		NameIdFormat:                   p.config.NameIdFormat,
		AssertionConsumerServiceURL:    p.acsUrl,
		ServiceProviderSLOURL:          p.sloUrl,
		ServiceProviderIssuer:          p.metadataUrl,
		AudienceURI:                    p.metadataUrl,
		SignAuthnRequests:              p.config.SignAuthRequest,
		SPKeyStore:                     p.keyStore,
		SkipSignatureValidation:        p.config.SkipSignatureValidation,
		SignAuthnRequestsCanonicalizer: dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList(""),
	}
	p.samlSPAvailable = true

	return nil
}

func (p *provider) loadUrlMetadata(ctx context.Context) error {
	// recreate service provider on load without an IdP data for canopsis metadata endpoint.
	// the canopsis saml metadata endpoint should work even if an IdP is unavailable.
	p.samlSP = &saml2.SAMLServiceProvider{
		AssertionConsumerServiceURL:    p.acsUrl,
		ServiceProviderSLOURL:          p.sloUrl,
		ServiceProviderIssuer:          p.metadataUrl,
		AudienceURI:                    p.metadataUrl,
		SignAuthnRequests:              p.config.SignAuthRequest,
		SPKeyStore:                     p.keyStore,
		NameIdFormat:                   p.config.NameIdFormat,
		SkipSignatureValidation:        p.config.SkipSignatureValidation,
		SignAuthnRequestsCanonicalizer: dsig.MakeC14N10ExclusiveCanonicalizerWithPrefixList(""),
	}

	dt, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return errors.New("unknown type of http.DefaultTransport")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, p.config.IdPMetadataUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to create IdP metadata request: %w", err)
	}

	tr := dt.Clone()
	tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: p.config.InsecureSkipVerify} //nolint:gosec

	hc := http.Client{Timeout: MetadataReqTimeout, Transport: tr}

	res, err := hc.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send IdP metadata request: %w", err)
	}

	defer res.Body.Close()

	idpMetadata := samltypes.EntityDescriptor{}
	err = xml.NewDecoder(res.Body).Decode(&idpMetadata)
	if err != nil {
		return fmt.Errorf("failed to unmarshal IdP metadata response: %w", err)
	}

	ssoLocation, sloLocation, certStore, err := processIdPMetadata(idpMetadata)
	if err != nil {
		return fmt.Errorf("failed to process IdP metadata response: %w", err)
	}

	p.samlSP.IdentityProviderSSOURL = ssoLocation
	p.samlSP.IdentityProviderSLOURL = sloLocation
	p.samlSP.IdentityProviderIssuer = idpMetadata.EntityID
	p.samlSP.IDPCertificateStore = &certStore
	p.samlSPAvailable = true

	now := time.Now()
	if idpMetadata.ValidUntil.After(now) {
		p.samlSPValidUntil = idpMetadata.ValidUntil
	} else {
		p.samlSPValidUntil = now.Add(auth.DefaultMetaValidDuration)
	}

	return nil
}

func processIdPMetadata(idpMetadata samltypes.EntityDescriptor) (string, string, dsig.MemoryX509CertificateStore, error) {
	certStore := dsig.MemoryX509CertificateStore{Roots: []*x509.Certificate{}}

	for _, kd := range idpMetadata.IDPSSODescriptor.KeyDescriptors {
		for idx, xCert := range kd.KeyInfo.X509Data.X509Certificates {
			if xCert.Data == "" {
				panic(fmt.Errorf("metadata certificate(%d) must not be empty", idx))
			}
			certData, err := base64.StdEncoding.DecodeString(xCert.Data)
			if err != nil {
				return "", "", dsig.MemoryX509CertificateStore{}, fmt.Errorf("failed to decode metadata certificate(%d) data: %w", idx, err)
			}

			idpCert, err := x509.ParseCertificate(certData)
			if err != nil {
				return "", "", dsig.MemoryX509CertificateStore{}, fmt.Errorf("failed to parse metadata certificate(%d) data: %w", idx, err)
			}

			certStore.Roots = append(certStore.Roots, idpCert)
		}
	}

	ssoLocation := ""
	sloLocation := ""

	if len(idpMetadata.IDPSSODescriptor.SingleSignOnServices) > 0 {
		ssoLocation = idpMetadata.IDPSSODescriptor.SingleSignOnServices[0].Location
	}

	if len(idpMetadata.IDPSSODescriptor.SingleLogoutServices) > 0 {
		sloLocation = idpMetadata.IDPSSODescriptor.SingleLogoutServices[0].Location
	}

	return ssoLocation, sloLocation, certStore, nil
}

func (p *provider) isServiceProviderAvailable(c *gin.Context) bool {
	if p.config.IdPMetadataUrl != "" && (!p.samlSPAvailable || p.isSamlSPExpired()) {
		err := p.loadUrlMetadata(c)
		if err != nil {
			p.logger.Warn().Str("error", err.Error()).Msg("failed to load IdP metadata")
			c.AbortWithStatus(http.StatusServiceUnavailable)

			return false
		}
	}

	return true
}

func (p *provider) isSamlSPExpired() bool {
	return time.Now().After(p.samlSPValidUntil)
}

func (p *provider) SamlMetadataHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		meta, err := p.samlSP.MetadataWithSLO(0)
		if err != nil {
			p.logger.Err(err).Msg("samlMetadataHandler: MetadataWithSLO error")
			panic(err)
		}

		if len(meta.SPSSODescriptor.SingleLogoutServices) > 0 && p.config.CanopsisSSOBinding == BindingRedirect {
			meta.SPSSODescriptor.SingleLogoutServices[0].Binding = saml2.BindingHttpRedirect
		}

		if len(meta.SPSSODescriptor.AssertionConsumerServices) > 0 {
			if p.config.CanopsisACSBinding == BindingRedirect {
				meta.SPSSODescriptor.AssertionConsumerServices[0].Binding = saml2.BindingHttpRedirect
			}

			if p.config.ACSIndex != nil {
				meta.SPSSODescriptor.AssertionConsumerServices[0].Index = *p.config.ACSIndex
			} else {
				meta.SPSSODescriptor.AssertionConsumerServices[0].Index = 1
			}
		}

		c.XML(http.StatusOK, meta)
	}
}

type samlLoginRequest struct {
	RelayState string `form:"relayState" binding:"required,url"`
}

func (p *provider) SamlAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.samlMx.Lock()
		defer p.samlMx.Unlock()

		ok := p.isServiceProviderAvailable(c)
		if !ok {
			return
		}

		if p.samlSP.IdentityProviderSSOURL == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		request := samlLoginRequest{}

		if err := c.ShouldBindQuery(&request); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewValidationErrorResponse(err, request))
			return
		}

		if p.config.CanopsisSSOBinding == BindingRedirect {
			var authRequest *etree.Document
			var err error

			if p.config.SignAuthRequest {
				authRequest, err = p.samlSP.BuildAuthRequestDocument()
			} else {
				authRequest, err = p.samlSP.BuildAuthRequestDocumentNoSig()
			}

			if err != nil {
				p.logger.Err(err).Msg("samlAuthHandler: BuildAuthRequestDocument error")
				panic(err)
			}

			el := authRequest.SelectElement("AuthnRequest")
			attr := el.SelectAttr("ProtocolBinding")
			attr.Value = saml2.BindingHttpRedirect

			authUrl, err := p.samlSP.BuildAuthURLRedirect(request.RelayState, authRequest)
			if err != nil {
				p.logger.Err(err).Msg("samlAuthHandler: parse IdentityProviderSSOURL error")
				panic(err)
			}

			c.Redirect(http.StatusPermanentRedirect, authUrl)
		}

		if p.config.CanopsisSSOBinding == BindingPOST {
			b, err := p.samlSP.BuildAuthBodyPost(request.RelayState)
			if err != nil {
				p.logger.Err(err).Msg("samlAuthHandler: BuildAuthBodyPost error")
				panic(err)
			}

			c.Data(http.StatusOK, gin.MIMEHTML, b)
		}
	}
}

func (p *provider) SamlAcsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.samlMx.Lock()
		defer p.samlMx.Unlock()

		ok := p.isServiceProviderAvailable(c)
		if !ok {
			return
		}

		samlResponse, exists := c.GetPostForm("SAMLResponse")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(errors.New("SAMLResponse doesn't exist")))
			return
		}

		relayState, exists := c.GetPostForm("RelayState")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(errors.New("RelayState doesn't exist")))
			return
		}

		relayUrl, err := url.Parse(relayState)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(errors.New("RelayState is not a valid url")))
			return
		}

		assertionInfo, err := p.samlSP.RetrieveAssertionInfo(samlResponse)
		if err != nil {
			p.logger.Err(err).Msg("assertion is not valid")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if assertionInfo.WarningInfo.InvalidTime {
			p.logger.Err(errors.New("invalid time")).Msg("assertion is not valid")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		if assertionInfo.WarningInfo.NotInAudience {
			p.logger.Err(errors.New("not in audience")).Msg("assertion is not valid")
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		user, err := p.userProvider.FindByExternalSource(c, assertionInfo.NameID, security.SourceSaml)
		if err != nil {
			p.logger.Err(err).Msg("samlAcsHandler: userProvider FindByExternalSource error")
			panic(err)
		}

		if user == nil {
			var ok bool
			user, ok = p.createUser(c, relayUrl, assertionInfo)
			if !ok {
				return
			}
		} else if !user.IsEnabled {
			c.AbortWithStatus(http.StatusForbidden)
			return
		} else if !p.updateUser(c, relayUrl, assertionInfo, user) {
			return
		}

		maintenanceConf, err := p.maintenanceAdapter.GetConfig(c)
		if err != nil {
			panic(err)
		}

		if maintenanceConf.Enabled {
			ok, err := p.enforcer.Enforce(user.ID, apisecurity.PermMaintenance, model.PermissionCan)
			if err != nil {
				panic(err)
			}

			if !ok {
				c.AbortWithStatusJSON(http.StatusServiceUnavailable, common.CanopsisUnderMaintenanceResponse)
				return
			}
		}

		err = p.enforcer.LoadPolicy()
		if err != nil {
			panic(fmt.Errorf("reload enforcer error: %w", err))
		}

		var accessToken string
		if p.config.ExpirationInterval != "" || assertionInfo.SessionNotOnOrAfter == nil {
			accessToken, err = p.tokenService.Create(c, *user, security.AuthMethodSaml)
		} else {
			accessToken, err = p.tokenService.CreateWithExpiration(c, *user, security.AuthMethodSaml, *assertionInfo.SessionNotOnOrAfter)
		}
		if err != nil {
			panic(err)
		}

		query := relayUrl.Query()
		query.Set("access_token", accessToken)
		relayUrl.RawQuery = query.Encode()

		c.Redirect(http.StatusSeeOther, relayUrl.String())
	}
}

func (p *provider) getAssocAttribute(attrs saml2.Values, canopsisName, defaultValue string) string {
	return cmp.Or(attrs.Get(p.config.IdPAttributesMap[canopsisName]), defaultValue)
}

func (p *provider) getAssocArrayAttribute(attrs saml2.Values, canopsisName string, defaultValue []string) []string {
	v := attrs.GetAll(p.config.IdPAttributesMap[canopsisName])
	if len(v) != 0 {
		return v
	}

	return defaultValue
}

func (p *provider) SamlSloHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		p.samlMx.Lock()
		defer p.samlMx.Unlock()

		ok := p.isServiceProviderAvailable(c)
		if !ok {
			return
		}

		if p.samlSP.IdentityProviderSLOURL == "" {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		samlRequest, exists := c.GetQuery("SAMLRequest")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(errors.New("SAMLRequest doesn't exist")))
			return
		}

		relayState, exists := c.GetQuery("RelayState")
		if !exists {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(errors.New("RelayState doesn't exist")))
			return
		}

		request, err := p.samlSP.ValidateEncodedLogoutRequestPOST(samlRequest)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, common.NewErrorResponse(err))
			return
		}

		user, err := p.userProvider.FindByExternalSource(c, request.NameID.Value, security.SourceSaml)
		if err != nil {
			p.logger.Err(err).Msg("samlSloHandler: userProvider FindByExternalSource error")
			panic(err)
		}

		if user == nil {
			responseUrl, err := p.buildLogoutResponseUrl(saml2.StatusCodeUnknownPrincipal, request.ID, relayState)
			if err != nil {
				p.logger.Err(err).Msg("samlSloHandler: buildLogoutResponseUrl error")
				panic(err)
			}

			c.Redirect(http.StatusPermanentRedirect, responseUrl.String())
			return
		}

		err = p.sessionStore.ExpireSessions(c, user.ID, security.AuthMethodSaml)
		if err != nil {
			responseUrl, err := p.buildLogoutResponseUrl(saml2.StatusCodeUnknownPrincipal, request.ID, relayState)
			if err != nil {
				p.logger.Err(err).Msg("samlSloHandler: buildLogoutResponseUrl error")
				panic(err)
			}

			c.Redirect(http.StatusPermanentRedirect, responseUrl.String())
			return
		}

		err = p.tokenService.DeleteBy(c, user.ID, security.AuthMethodSaml)
		if err != nil {
			responseUrl, err := p.buildLogoutResponseUrl(saml2.StatusCodeUnknownPrincipal, request.ID, relayState)
			if err != nil {
				p.logger.Err(err).Msg("samlSloHandler: buildLogoutResponseUrl error")
				panic(err)
			}

			c.Redirect(http.StatusPermanentRedirect, responseUrl.String())
			return
		}

		responseUrl, err := p.buildLogoutResponseUrl(saml2.StatusCodeSuccess, request.ID, relayState)
		if err != nil {
			p.logger.Err(err).Msg("samlSloHandler: buildLogoutResponseUrl error")
			panic(err)
		}

		c.Redirect(http.StatusPermanentRedirect, responseUrl.String())
	}
}

func (p *provider) buildLogoutResponseUrl(status, reqID, relayState string) (*url.URL, error) {
	responseDoc, err := p.samlSP.BuildLogoutResponseDocument(status, reqID)
	if err != nil {
		return nil, err
	}

	buffer, err := p.encodeAndCompress(responseDoc)
	if err != nil {
		return nil, err
	}

	responseUrl, err := url.Parse(p.samlSP.IdentityProviderSLOURL)
	if err != nil {
		return nil, err
	}

	query := responseUrl.Query()
	query.Set("SAMLResponse", buffer.String())
	query.Set("RelayState", relayState)

	responseUrl.RawQuery = query.Encode()

	return responseUrl, nil
}

func (p *provider) encodeAndCompress(doc io.WriterTo) (_ *bytes.Buffer, resErr error) {
	buffer := &bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, buffer)

	defer func() {
		err := encoder.Close()
		if err != nil && resErr == nil {
			resErr = err
		}
	}()

	compressor, err := flate.NewWriter(encoder, flate.BestCompression)
	if err != nil {
		return nil, err
	}

	defer func() {
		err = compressor.Close()
		if err != nil && resErr == nil {
			resErr = err
		}
	}()

	_, err = doc.WriteTo(compressor)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (p *provider) createUser(c *gin.Context, relayUrl *url.URL, assertionInfo *saml2.AssertionInfo) (*security.User, bool) {
	if !p.config.AutoUserRegistration {
		p.logger.Err(fmt.Errorf("user with external_id = %s is not found", assertionInfo.NameID)).Msg("autoUserRegistration is disabled")
		p.errorRedirect(c, relayUrl, "This user is not allowed to log into Canopsis")

		return nil, false
	}

	roles, err := p.roleProvider.GetValidRoleIDs(c, p.getAssocArrayAttribute(assertionInfo.Values, security.UserRole, []string{}), p.config.DefaultRole)
	if err != nil {
		roleNotFoundError := roleprovider.ProviderError{}
		if errors.As(err, &roleNotFoundError) {
			p.logger.Err(roleNotFoundError).Str("external_id", assertionInfo.NameID).Msg("failed to get user roles from saml assertion")
			p.errorRedirect(c, relayUrl, roleNotFoundError.Error())

			return nil, false
		}

		panic(err)
	}

	user := &security.User{
		Name:       p.getAssocAttribute(assertionInfo.Values, security.UserName, assertionInfo.NameID),
		Roles:      roles,
		IsEnabled:  true,
		ExternalID: assertionInfo.NameID,
		Source:     security.SourceSaml,
		Firstname:  p.getAssocAttribute(assertionInfo.Values, security.UserFirstName, ""),
		Lastname:   p.getAssocAttribute(assertionInfo.Values, security.UserLastName, ""),
		Email:      p.getAssocAttribute(assertionInfo.Values, security.UserEmail, ""),
		IdPRoles:   roles,
	}

	err = p.userProvider.Save(c, user)
	if err != nil {
		panic(fmt.Errorf("failed to save saml user with external_id = %s: %w", user.ExternalID, err))
	}

	return user, true
}

func (p *provider) updateUser(c *gin.Context, relayUrl *url.URL, assertionInfo *saml2.AssertionInfo, user *security.User) bool {
	roles, err := p.roleProvider.GetValidRoleIDs(c, p.getAssocArrayAttribute(assertionInfo.Values, security.UserRole, []string{}), p.config.DefaultRole)
	if err != nil {
		roleNotFoundError := roleprovider.ProviderError{}
		if errors.As(err, &roleNotFoundError) {
			p.logger.Err(roleNotFoundError).Str("external_id", user.ExternalID).Msg("failed to get user roles from saml assertion")
			p.errorRedirect(c, relayUrl, roleNotFoundError.Error())

			return false
		}

		panic(err)
	}

	user.Name = p.getAssocAttribute(assertionInfo.Values, security.UserName, user.Name)
	user.Firstname = p.getAssocAttribute(assertionInfo.Values, security.UserFirstName, user.Firstname)
	user.Lastname = p.getAssocAttribute(assertionInfo.Values, security.UserLastName, user.Lastname)
	user.Email = p.getAssocAttribute(assertionInfo.Values, security.UserEmail, user.Email)
	user.SetRolesFromIdP(roles, p.config.AllowExtraRoles)

	err = p.userProvider.Save(c, user)
	if err != nil {
		panic(fmt.Errorf("failed to update saml user with external_id = %s: %w", user.ExternalID, err))
	}

	return true
}

func (p *provider) errorRedirect(c *gin.Context, relayUrl *url.URL, errorMessage string) {
	query := relayUrl.Query()
	query.Set("errorMessage", errorMessage)
	relayUrl.RawQuery = query.Encode()

	c.Redirect(http.StatusPermanentRedirect, relayUrl.String())
}
