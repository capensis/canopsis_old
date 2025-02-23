package security

const (
	SourceLdap   = "ldap"
	SourceCas    = "cas"
	SourceSaml   = "saml"
	SourceOauth2 = "oauth2"
)

// User field constants to unify values to map fields from external identity providers.
const (
	UserName      = "name"
	UserFirstName = "firstname"
	UserLastName  = "lastname"
	UserEmail     = "email"
	UserRole      = "role"
)

// User represents user model.
type User struct {
	ID             string   `bson:"_id"`
	Name           string   `bson:"name"`
	DisplayName    string   `bson:"display_name,omitempty"`
	Firstname      string   `bson:"firstname"`
	Lastname       string   `bson:"lastname"`
	Email          string   `bson:"email"`
	HashedPassword string   `bson:"password,omitempty"`
	AuthApiKey     string   `bson:"authkey"`
	Roles          []string `bson:"roles"`
	Contact        struct {
		Name    string `bson:"name"`
		Address string `bson:"address"`
	} `bson:"contact"`
	IsEnabled  bool   `bson:"enable"`
	ExternalID string `bson:"external_id"`
	Source     string `bson:"source"`

	// IdPRoles field show roles from idp, and they should be used ONLY in idp/canopsis role merging, see SetRolesFromIdP.
	IdPRoles []string `bson:"idp_roles"`
}

func (u *User) SetRolesFromIdP(newIdpRoles []string, mergeRoles bool) {
	if mergeRoles {
		idpRolesMap := make(map[string]bool, len(u.IdPRoles)+len(newIdpRoles))

		for _, role := range u.IdPRoles {
			idpRolesMap[role] = true
		}

		for _, role := range newIdpRoles {
			idpRolesMap[role] = true
		}

		u.IdPRoles = make([]string, len(newIdpRoles))
		copy(u.IdPRoles, newIdpRoles)

		for _, role := range u.Roles {
			if !idpRolesMap[role] {
				newIdpRoles = append(newIdpRoles, role)
			}
		}

		u.Roles = newIdpRoles
	} else {
		u.IdPRoles = newIdpRoles
		u.Roles = newIdpRoles
	}
}
