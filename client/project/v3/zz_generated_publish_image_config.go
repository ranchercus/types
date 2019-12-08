package client

const (
	PublishImageConfigType                      = "publishImageConfig"
	PublishImageConfigFieldBuildContext         = "buildContext"
	PublishImageConfigFieldCallbackScript       = "callbackScript"
	PublishImageConfigFieldCallbackScriptParams = "callbackScriptParams"
	PublishImageConfigFieldContainerIndex       = "containerIndex"
	PublishImageConfigFieldDeploy               = "deploy"
	PublishImageConfigFieldDeployService        = "deployService"
	PublishImageConfigFieldDockerfilePath       = "dockerfilePath"
	PublishImageConfigFieldPort                 = "port"
	PublishImageConfigFieldPushRemote           = "pushRemote"
	PublishImageConfigFieldRegistry             = "registry"
	PublishImageConfigFieldTag                  = "tag"
	PublishImageConfigFieldTargetNamespace      = "targetNamespace"
	PublishImageConfigFieldWorkloadId           = "workloadId"
)

type PublishImageConfig struct {
	BuildContext         string `json:"buildContext,omitempty" yaml:"buildContext,omitempty"`
	CallbackScript       string `json:"callbackScript,omitempty" yaml:"callbackScript,omitempty"`
	CallbackScriptParams string `json:"callbackScriptParams,omitempty" yaml:"callbackScriptParams,omitempty"`
	ContainerIndex       int64  `json:"containerIndex,omitempty" yaml:"containerIndex,omitempty"`
	Deploy               bool   `json:"deploy,omitempty" yaml:"deploy,omitempty"`
	DeployService        bool   `json:"deployService,omitempty" yaml:"deployService,omitempty"`
	DockerfilePath       string `json:"dockerfilePath,omitempty" yaml:"dockerfilePath,omitempty"`
	Port                 string `json:"port,omitempty" yaml:"port,omitempty"`
	PushRemote           bool   `json:"pushRemote,omitempty" yaml:"pushRemote,omitempty"`
	Registry             string `json:"registry,omitempty" yaml:"registry,omitempty"`
	Tag                  string `json:"tag,omitempty" yaml:"tag,omitempty"`
	TargetNamespace      string `json:"targetNamespace,omitempty" yaml:"targetNamespace,omitempty"`
	WorkloadId           string `json:"workloadId,omitempty" yaml:"workloadId,omitempty"`
}
