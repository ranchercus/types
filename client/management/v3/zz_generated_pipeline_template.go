package client

import (
	"github.com/rancher/norman/types"
)

const (
	PipelineTemplateType                 = "pipelineTemplate"
	PipelineTemplateFieldAnnotations     = "annotations"
	PipelineTemplateFieldCategories      = "categories"
	PipelineTemplateFieldClusterID       = "clusterId"
	PipelineTemplateFieldCreated         = "created"
	PipelineTemplateFieldCreatorID       = "creatorId"
	PipelineTemplateFieldLabels          = "labels"
	PipelineTemplateFieldName            = "name"
	PipelineTemplateFieldNamespaceId     = "namespaceId"
	PipelineTemplateFieldOwnerReferences = "ownerReferences"
	PipelineTemplateFieldQuestions       = "questions"
	PipelineTemplateFieldRemoved         = "removed"
	PipelineTemplateFieldTemplate        = "template"
	PipelineTemplateFieldUUID            = "uuid"
)

type PipelineTemplate struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Categories      []string          `json:"categories,omitempty" yaml:"categories,omitempty"`
	ClusterID       string            `json:"clusterId,omitempty" yaml:"clusterId,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	NamespaceId     string            `json:"namespaceId,omitempty" yaml:"namespaceId,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Questions       []string          `json:"questions,omitempty" yaml:"questions,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Template        string            `json:"template,omitempty" yaml:"template,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type PipelineTemplateCollection struct {
	types.Collection
	Data   []PipelineTemplate `json:"data,omitempty"`
	client *PipelineTemplateClient
}

type PipelineTemplateClient struct {
	apiClient *Client
}

type PipelineTemplateOperations interface {
	List(opts *types.ListOpts) (*PipelineTemplateCollection, error)
	Create(opts *PipelineTemplate) (*PipelineTemplate, error)
	Update(existing *PipelineTemplate, updates interface{}) (*PipelineTemplate, error)
	Replace(existing *PipelineTemplate) (*PipelineTemplate, error)
	ByID(id string) (*PipelineTemplate, error)
	Delete(container *PipelineTemplate) error
}

func newPipelineTemplateClient(apiClient *Client) *PipelineTemplateClient {
	return &PipelineTemplateClient{
		apiClient: apiClient,
	}
}

func (c *PipelineTemplateClient) Create(container *PipelineTemplate) (*PipelineTemplate, error) {
	resp := &PipelineTemplate{}
	err := c.apiClient.Ops.DoCreate(PipelineTemplateType, container, resp)
	return resp, err
}

func (c *PipelineTemplateClient) Update(existing *PipelineTemplate, updates interface{}) (*PipelineTemplate, error) {
	resp := &PipelineTemplate{}
	err := c.apiClient.Ops.DoUpdate(PipelineTemplateType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *PipelineTemplateClient) Replace(obj *PipelineTemplate) (*PipelineTemplate, error) {
	resp := &PipelineTemplate{}
	err := c.apiClient.Ops.DoReplace(PipelineTemplateType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *PipelineTemplateClient) List(opts *types.ListOpts) (*PipelineTemplateCollection, error) {
	resp := &PipelineTemplateCollection{}
	err := c.apiClient.Ops.DoList(PipelineTemplateType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *PipelineTemplateCollection) Next() (*PipelineTemplateCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &PipelineTemplateCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *PipelineTemplateClient) ByID(id string) (*PipelineTemplate, error) {
	resp := &PipelineTemplate{}
	err := c.apiClient.Ops.DoByID(PipelineTemplateType, id, resp)
	return resp, err
}

func (c *PipelineTemplateClient) Delete(container *PipelineTemplate) error {
	return c.apiClient.Ops.DoResourceDelete(PipelineTemplateType, &container.Resource)
}
