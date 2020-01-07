package client

const (
	PipelineSettingType                     = "pipelineSetting"
	PipelineSettingFieldCallbackScripts     = "callbackScripts"
	PipelineSettingFieldDefaultRegistry     = "defaultRegistry"
	PipelineSettingFieldEnableLocalRegistry = "enableLocalRegistry"
	PipelineSettingFieldLocalShare          = "localShare"
	PipelineSettingFieldNodeSelector        = "nodeSelector"
	PipelineSettingFieldNodeToleration      = "nodeToleration"
	PipelineSettingFieldRegistryInsecure    = "registryInsecure"
	PipelineSettingFieldSonarScanner        = "sonarScanner"
)

type PipelineSetting struct {
	CallbackScripts     []CallbackScript `json:"callbackScripts,omitempty" yaml:"callbackScripts,omitempty"`
	DefaultRegistry     string           `json:"defaultRegistry,omitempty" yaml:"defaultRegistry,omitempty"`
	EnableLocalRegistry bool             `json:"enableLocalRegistry,omitempty" yaml:"enableLocalRegistry,omitempty"`
	LocalShare          string           `json:"localShare,omitempty" yaml:"localShare,omitempty"`
	NodeSelector        string           `json:"nodeSelector,omitempty" yaml:"nodeSelector,omitempty"`
	NodeToleration      string           `json:"nodeToleration,omitempty" yaml:"nodeToleration,omitempty"`
	RegistryInsecure    bool             `json:"registryInsecure,omitempty" yaml:"registryInsecure,omitempty"`
	SonarScanner        *SonarScanner    `json:"sonarScanner,omitempty" yaml:"sonarScanner,omitempty"`
}
