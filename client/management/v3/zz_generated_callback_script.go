package client

const (
	CallbackScriptType        = "callbackScript"
	CallbackScriptFieldLabel  = "label"
	CallbackScriptFieldScript = "script"
	CallbackScriptFieldValue  = "value"
)

type CallbackScript struct {
	Label  string `json:"label,omitempty" yaml:"label,omitempty"`
	Script string `json:"script,omitempty" yaml:"script,omitempty"`
	Value  string `json:"value,omitempty" yaml:"value,omitempty"`
}
