package client

const (
	SonarScannerType                = "sonarScanner"
	SonarScannerFieldHostUrl        = "hostUrl"
	SonarScannerFieldLogin          = "login"
	SonarScannerFieldSourceEncoding = "sourceEncoding"
)

type SonarScanner struct {
	HostUrl        string `json:"hostUrl,omitempty" yaml:"hostUrl,omitempty"`
	Login          string `json:"login,omitempty" yaml:"login,omitempty"`
	SourceEncoding string `json:"sourceEncoding,omitempty" yaml:"sourceEncoding,omitempty"`
}
