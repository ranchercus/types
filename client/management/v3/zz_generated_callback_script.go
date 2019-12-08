package client

const (
	CallbackScriptType       = "callbackScript"
	CallbackScriptFieldLabel = "label"
	CallbackScriptFieldValue = "value"
)

type CallbackScript struct {
	Label string `json:"label,omitempty" yaml:"label,omitempty"`
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
