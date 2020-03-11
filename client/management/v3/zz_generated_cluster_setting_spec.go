package client

const (
	ClusterSettingSpecType                 = "clusterSettingSpec"
	ClusterSettingSpecFieldLoggingSetting  = "loggingSetting"
	ClusterSettingSpecFieldPipelineSetting = "pipelineSetting"
	ClusterSettingSpecFieldRegistrySetting = "registrySetting"
	ClusterSettingSpecFieldWorkloadSetting = "workloadSetting"
)

type ClusterSettingSpec struct {
	LoggingSetting  *LoggingSetting  `json:"loggingSetting,omitempty" yaml:"loggingSetting,omitempty"`
	PipelineSetting *PipelineSetting `json:"pipelineSetting,omitempty" yaml:"pipelineSetting,omitempty"`
	RegistrySetting *RegistrySetting `json:"registrySetting,omitempty" yaml:"registrySetting,omitempty"`
	WorkloadSetting *WorkloadSetting `json:"workloadSetting,omitempty" yaml:"workloadSetting,omitempty"`
}
