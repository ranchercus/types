package client

const (
	PipelineConfigType              = "pipelineConfig"
	PipelineConfigFieldBranch       = "branch"
	PipelineConfigFieldEnv          = "env"
	PipelineConfigFieldNotification = "notification"
	PipelineConfigFieldStages       = "stages"
	PipelineConfigFieldTimeout      = "timeout"
)

type PipelineConfig struct {
	Branch       *Constraint           `json:"branch,omitempty" yaml:"branch,omitempty"`
	Env          map[string]string     `json:"env,omitempty" yaml:"env,omitempty"`
	Notification *PipelineNotification `json:"notification,omitempty" yaml:"notification,omitempty"`
	Stages       []Stage               `json:"stages,omitempty" yaml:"stages,omitempty"`
	Timeout      int64                 `json:"timeout,omitempty" yaml:"timeout,omitempty"`
}
