package appinfo

import "git.canopsis.net/canopsis/canopsis-community/community/go-engines-community/lib/canopsis/types"

type VersionConf struct {
	Edition string `json:"edition" bson:"edition"`
	Stack   string `json:"stack" bson:"stack"`
	Version string `json:"version" bson:"version"`
}

type IntervalUnit struct {
	Interval uint   `json:"interval" bson:"interval"`
	Unit     string `json:"unit" bson:"unit" binding:"oneof=s h m"`
}

type PopupTimeout struct {
	Info  *IntervalUnit `json:"info,omitempty" bson:"info,omitempty"`
	Error *IntervalUnit `json:"error,omitempty" bson:"error,omitempty"`
}

type UserInterfaceConf struct {
	AppTitle                  string        `json:"app_title,omitempty" bson:"app_title,omitempty"`
	Footer                    string        `json:"footer,omitempty" bson:"footer,omitempty"`
	LoginPageDescription      string        `json:"login_page_description,omitempty" bson:"login_page_description,omitempty"`
	Logo                      string        `json:"logo,omitempty" bson:"logo,omitempty"`
	Language                  string        `json:"language,omitempty" bson:"language,omitempty" binding:"oneoforempty=fr en"`
	PopupTimeout              *PopupTimeout `json:"popup_timeout,omitempty" bson:"popup_timeout,omitempty"`
	AllowChangeSeverityToInfo bool          `json:"allow_change_severity_to_info" bson:"allow_change_severity_to_info"`
	MaxMatchedItems           int64         `json:"max_matched_items" bson:"max_matched_items" binding:"gt=0"`
	CheckCountRequestTimeout  int64         `json:"check_count_request_timeout" bson:"check_count_request_timeout" binding:"gt=0"`
}

type GlobalConf struct {
	Timezone          string `json:"timezone,omitempty"`
	FileUploadMaxSize int64  `json:"file_upload_max_size"`
}

type RemediationConf struct {
	PauseManualInstructionInterval types.DurationWithUnit `json:"pause_manual_instruction_interval"`
	JobConfigTypes                 []JobConfigType        `json:"job_config_types"`
}

type JobConfigType struct {
	Name     string `json:"name"`
	AuthType string `json:"auth_type"`
}

type AppInfoResponse struct {
	UserInterfaceConf
	GlobalConf
	VersionConf
	Remediation RemediationConf `json:"remediation"`
}

type LoginConfig struct {
	CasConfig  LoginConfigMethod `json:"casconfig,omitempty"`
	LdapConfig LoginConfigMethod `json:"ldapconfig,omitempty"`
	SamlConfig LoginConfigMethod `json:"saml2config,omitempty"`
}

type LoginConfigMethod struct {
	Title  string `json:"title,omitempty"`
	Enable bool   `json:"enable"`
}

type LoginConfigResponse struct {
	LoginConfig       LoginConfig       `json:"login_config"`
	UserInterfaceConf UserInterfaceConf `json:"user_interface"`
	VersionConf
}
