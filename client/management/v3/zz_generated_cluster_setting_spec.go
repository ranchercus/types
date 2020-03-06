package client

const (
	ClusterSettingSpecType                 = "clusterSettingSpec"
	ClusterSettingSpecFieldPipelineSetting = "pipelineSetting"
	ClusterSettingSpecFieldRegistrySetting = "registrySetting"
	ClusterSettingSpecFieldWorkloadSetting = "workloadSetting"
)

type ClusterSettingSpec struct {
	PipelineSetting *PipelineSetting `json:"pipelineSetting,omitempty" yaml:"pipelineSetting,omitempty"`
	RegistrySetting *RegistrySetting `json:"registrySetting,omitempty" yaml:"registrySetting,omitempty"`
	WorkloadSetting *WorkloadSetting `json:"workloadSetting,omitempty" yaml:"workloadSetting,omitempty"`
}
