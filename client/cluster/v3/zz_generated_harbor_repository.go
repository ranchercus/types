package client

import (
	"github.com/rancher/norman/types"
)

const (
	HarborRepositoryType                 = "harborRepository"
	HarborRepositoryFieldAnnotations     = "annotations"
	HarborRepositoryFieldCreated         = "created"
	HarborRepositoryFieldCreationTime    = "creation_time"
	HarborRepositoryFieldCreatorID       = "creatorId"
	HarborRepositoryFieldLabels          = "labels"
	HarborRepositoryFieldName            = "name"
	HarborRepositoryFieldOwnerReferences = "ownerReferences"
	HarborRepositoryFieldProjectId       = "project_id"
	HarborRepositoryFieldPullCount       = "pull_count"
	HarborRepositoryFieldRemoved         = "removed"
	HarborRepositoryFieldTagsCount       = "tags_count"
	HarborRepositoryFieldUUID            = "uuid"
)

type HarborRepository struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreationTime    string            `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectId       int64             `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	PullCount       int64             `json:"pull_count,omitempty" yaml:"pull_count,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	TagsCount       int64             `json:"tags_count,omitempty" yaml:"tags_count,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type HarborRepositoryCollection struct {
	types.Collection
	Data   []HarborRepository `json:"data,omitempty"`
	client *HarborRepositoryClient
}

type HarborRepositoryClient struct {
	apiClient *Client
}

type HarborRepositoryOperations interface {
	List(opts *types.ListOpts) (*HarborRepositoryCollection, error)
	Create(opts *HarborRepository) (*HarborRepository, error)
	Update(existing *HarborRepository, updates interface{}) (*HarborRepository, error)
	Replace(existing *HarborRepository) (*HarborRepository, error)
	ByID(id string) (*HarborRepository, error)
	Delete(container *HarborRepository) error
}

func newHarborRepositoryClient(apiClient *Client) *HarborRepositoryClient {
	return &HarborRepositoryClient{
		apiClient: apiClient,
	}
}

func (c *HarborRepositoryClient) Create(container *HarborRepository) (*HarborRepository, error) {
	resp := &HarborRepository{}
	err := c.apiClient.Ops.DoCreate(HarborRepositoryType, container, resp)
	return resp, err
}

func (c *HarborRepositoryClient) Update(existing *HarborRepository, updates interface{}) (*HarborRepository, error) {
	resp := &HarborRepository{}
	err := c.apiClient.Ops.DoUpdate(HarborRepositoryType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HarborRepositoryClient) Replace(obj *HarborRepository) (*HarborRepository, error) {
	resp := &HarborRepository{}
	err := c.apiClient.Ops.DoReplace(HarborRepositoryType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *HarborRepositoryClient) List(opts *types.ListOpts) (*HarborRepositoryCollection, error) {
	resp := &HarborRepositoryCollection{}
	err := c.apiClient.Ops.DoList(HarborRepositoryType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HarborRepositoryCollection) Next() (*HarborRepositoryCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HarborRepositoryCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HarborRepositoryClient) ByID(id string) (*HarborRepository, error) {
	resp := &HarborRepository{}
	err := c.apiClient.Ops.DoByID(HarborRepositoryType, id, resp)
	return resp, err
}

func (c *HarborRepositoryClient) Delete(container *HarborRepository) error {
	return c.apiClient.Ops.DoResourceDelete(HarborRepositoryType, &container.Resource)
}
