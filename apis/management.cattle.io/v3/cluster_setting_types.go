package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterSettingSpec `json:"spec"`
}

type ClusterSettingSpec struct {
	PipelineSetting PipelineSetting `json:"pipelineSetting"`
	WorkloadSetting WorkloadSetting `json:"workloadSetting"`
}

type PipelineSetting struct {
	RegistryInsecure bool             `json:"registryInsecure"`
	DefaultRegistry  string           `json:"defaultRegistry"`
	NodeToleration   string           `json:"nodeToleration"`
	NodeSelector     string           `json:"nodeSelector"`
	LocalShare       string           `json:"localShare"`
	CallbackScripts  []CallbackScript `json:"callbackScripts"`
	SonarScanner     *SonarScanner    `json:"sonarScanner"`
}

type WorkloadSetting struct {
}

type CallbackScript struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

type SonarScanner struct {
	Login          string `json:"login"`
	SourceEncoding string `json:"sourceEncoding"`
	HostUrl        string `json:"hostUrl"`
}
