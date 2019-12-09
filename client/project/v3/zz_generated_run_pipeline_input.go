package client

const (
	RunPipelineInputType                   = "runPipelineInput"
	RunPipelineInputFieldBranch            = "branch"
	RunPipelineInputFieldRunCallbackScript = "runCallbackScript"
	RunPipelineInputFieldRunCodeScanner    = "runCodeScanner"
)

type RunPipelineInput struct {
	Branch            string `json:"branch,omitempty" yaml:"branch,omitempty"`
	RunCallbackScript bool   `json:"runCallbackScript,omitempty" yaml:"runCallbackScript,omitempty"`
	RunCodeScanner    bool   `json:"runCodeScanner,omitempty" yaml:"runCodeScanner,omitempty"`
}
