package client

const (
	HarborMetaType        = "harborMeta"
	HarborMetaFieldPublic = "public"
)

type HarborMeta struct {
	Public bool `json:"public,omitempty" yaml:"public,omitempty"`
}
