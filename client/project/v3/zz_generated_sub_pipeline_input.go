package client

const (
	SubPipelineInputType             = "subPipelineInput"
	SubPipelineInputFieldContextPath = "contextPath"
	SubPipelineInputFieldSubPath     = "subPath"
)

type SubPipelineInput struct {
	ContextPath string `json:"contextPath,omitempty" yaml:"contextPath,omitempty"`
	SubPath     string `json:"subPath,omitempty" yaml:"subPath,omitempty"`
}
