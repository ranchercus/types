package client

const (
	PipelineSpecType                        = "pipelineSpec"
	PipelineSpecFieldContextPath            = "contextPath"
	PipelineSpecFieldDisplayName            = "displayName"
	PipelineSpecFieldProjectID              = "projectId"
	PipelineSpecFieldRepositoryURL          = "repositoryUrl"
	PipelineSpecFieldSourceCodeCredentialID = "sourceCodeCredentialId"
	PipelineSpecFieldSubPath                = "subPath"
	PipelineSpecFieldTriggerWebhookPr       = "triggerWebhookPr"
	PipelineSpecFieldTriggerWebhookPush     = "triggerWebhookPush"
	PipelineSpecFieldTriggerWebhookTag      = "triggerWebhookTag"
)

type PipelineSpec struct {
	ContextPath            string `json:"contextPath,omitempty" yaml:"contextPath,omitempty"`
	DisplayName            string `json:"displayName,omitempty" yaml:"displayName,omitempty"`
	ProjectID              string `json:"projectId,omitempty" yaml:"projectId,omitempty"`
	RepositoryURL          string `json:"repositoryUrl,omitempty" yaml:"repositoryUrl,omitempty"`
	SourceCodeCredentialID string `json:"sourceCodeCredentialId,omitempty" yaml:"sourceCodeCredentialId,omitempty"`
	SubPath                string `json:"subPath,omitempty" yaml:"subPath,omitempty"`
	TriggerWebhookPr       bool   `json:"triggerWebhookPr,omitempty" yaml:"triggerWebhookPr,omitempty"`
	TriggerWebhookPush     bool   `json:"triggerWebhookPush,omitempty" yaml:"triggerWebhookPush,omitempty"`
	TriggerWebhookTag      bool   `json:"triggerWebhookTag,omitempty" yaml:"triggerWebhookTag,omitempty"`
}
