package client

const (
	RunPipelineInputType                   = "runPipelineInput"
	RunPipelineInputFieldBranch            = "branch"
	RunPipelineInputFieldRunCallbackScript = "runCallbackScript"
)

type RunPipelineInput struct {
	Branch            string `json:"branch,omitempty" yaml:"branch,omitempty"`
	RunCallbackScript bool   `json:"runCallbackScript,omitempty" yaml:"runCallbackScript,omitempty"`
}
