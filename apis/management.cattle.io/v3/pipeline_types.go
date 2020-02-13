package v3

import (
	"github.com/rancher/norman/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PipelineTemplate struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName string   `json:"displayName,omitempty" norman:"required"`
	ClusterName string   `json:"clusterName,omitempty" norman:"required,type=reference[cluster]"`
	Categories  []string `json:"categories,omitempty"`
	Template    string   `json:"template,omitempty"`
	Questions   []string `json:"questions,omitempty"`
}
