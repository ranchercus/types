package client

import (
	"github.com/rancher/norman/types"
)

const (
	HarborProjectType                   = "harborProject"
	HarborProjectFieldAnnotations       = "annotations"
	HarborProjectFieldCreated           = "created"
	HarborProjectFieldCreationTime      = "creation_time"
	HarborProjectFieldCreatorID         = "creatorId"
	HarborProjectFieldCurrentUserRoleId = "current_user_role_id"
	HarborProjectFieldLabels            = "labels"
	HarborProjectFieldMetadata          = "harborMeta"
	HarborProjectFieldName              = "name"
	HarborProjectFieldOwnerReferences   = "ownerReferences"
	HarborProjectFieldProjectId         = "project_id"
	HarborProjectFieldRemoved           = "removed"
	HarborProjectFieldRepoCount         = "repo_count"
	HarborProjectFieldUUID              = "uuid"
)

type HarborProject struct {
	types.Resource
	Annotations       map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created           string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreationTime      string            `json:"creation_time,omitempty" yaml:"creation_time,omitempty"`
	CreatorID         string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	CurrentUserRoleId int64             `json:"current_user_role_id,omitempty" yaml:"current_user_role_id,omitempty"`
	Labels            map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Metadata          *HarborMeta       `json:"harborMeta,omitempty" yaml:"harborMeta,omitempty"`
	Name              string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences   []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	ProjectId         int64             `json:"project_id,omitempty" yaml:"project_id,omitempty"`
	Removed           string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	RepoCount         int64             `json:"repo_count,omitempty" yaml:"repo_count,omitempty"`
	UUID              string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type HarborProjectCollection struct {
	types.Collection
	Data   []HarborProject `json:"data,omitempty"`
	client *HarborProjectClient
}

type HarborProjectClient struct {
	apiClient *Client
}

type HarborProjectOperations interface {
	List(opts *types.ListOpts) (*HarborProjectCollection, error)
	Create(opts *HarborProject) (*HarborProject, error)
	Update(existing *HarborProject, updates interface{}) (*HarborProject, error)
	Replace(existing *HarborProject) (*HarborProject, error)
	ByID(id string) (*HarborProject, error)
	Delete(container *HarborProject) error
}

func newHarborProjectClient(apiClient *Client) *HarborProjectClient {
	return &HarborProjectClient{
		apiClient: apiClient,
	}
}

func (c *HarborProjectClient) Create(container *HarborProject) (*HarborProject, error) {
	resp := &HarborProject{}
	err := c.apiClient.Ops.DoCreate(HarborProjectType, container, resp)
	return resp, err
}

func (c *HarborProjectClient) Update(existing *HarborProject, updates interface{}) (*HarborProject, error) {
	resp := &HarborProject{}
	err := c.apiClient.Ops.DoUpdate(HarborProjectType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HarborProjectClient) Replace(obj *HarborProject) (*HarborProject, error) {
	resp := &HarborProject{}
	err := c.apiClient.Ops.DoReplace(HarborProjectType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *HarborProjectClient) List(opts *types.ListOpts) (*HarborProjectCollection, error) {
	resp := &HarborProjectCollection{}
	err := c.apiClient.Ops.DoList(HarborProjectType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HarborProjectCollection) Next() (*HarborProjectCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HarborProjectCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HarborProjectClient) ByID(id string) (*HarborProject, error) {
	resp := &HarborProject{}
	err := c.apiClient.Ops.DoByID(HarborProjectType, id, resp)
	return resp, err
}

func (c *HarborProjectClient) Delete(container *HarborProject) error {
	return c.apiClient.Ops.DoResourceDelete(HarborProjectType, &container.Resource)
}
