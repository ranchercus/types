package client

const (
	ClusterSettingSpecType                 = "clusterSettingSpec"
	ClusterSettingSpecFieldPipelineSetting = "pipelineSetting"
	ClusterSettingSpecFieldWorkloadSetting = "workloadSetting"
)

type ClusterSettingSpec struct {
	PipelineSetting *PipelineSetting `json:"pipelineSetting,omitempty" yaml:"pipelineSetting,omitempty"`
	WorkloadSetting *WorkloadSetting `json:"workloadSetting,omitempty" yaml:"workloadSetting,omitempty"`
}
