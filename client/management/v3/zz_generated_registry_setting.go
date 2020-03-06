package client

const (
	RegistrySettingType          = "registrySetting"
	RegistrySettingFieldCert     = "cert"
	RegistrySettingFieldHost     = "host"
	RegistrySettingFieldInsecure = "insecure"
)

type RegistrySetting struct {
	Cert     string `json:"cert,omitempty" yaml:"cert,omitempty"`
	Host     string `json:"host,omitempty" yaml:"host,omitempty"`
	Insecure bool   `json:"insecure,omitempty" yaml:"insecure,omitempty"`
}
