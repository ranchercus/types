package client

const (
	LoggingSettingType               = "loggingSetting"
	LoggingSettingFieldEnforceDeploy = "enforceDeploy"
)

type LoggingSetting struct {
	EnforceDeploy bool `json:"enforceDeploy,omitempty" yaml:"enforceDeploy,omitempty"`
}
