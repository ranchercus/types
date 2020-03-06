package client

import (
	"github.com/rancher/norman/types"
)

const (
	HarborTagType                 = "harborTag"
	HarborTagFieldAnnotations     = "annotations"
	HarborTagFieldAuthor          = "author"
	HarborTagFieldCreated         = "created"
	HarborTagFieldCreatorID       = "creatorId"
	HarborTagFieldDockerVersion   = "docker_version"
	HarborTagFieldFullName        = "full_name"
	HarborTagFieldLabels          = "labels"
	HarborTagFieldName            = "name"
	HarborTagFieldOwnerReferences = "ownerReferences"
	HarborTagFieldRemoved         = "removed"
	HarborTagFieldRepo            = "repo"
	HarborTagFieldSize            = "size"
	HarborTagFieldTagLabels       = "tag_labels"
	HarborTagFieldUUID            = "uuid"
)

type HarborTag struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Author          string            `json:"author,omitempty" yaml:"author,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	DockerVersion   string            `json:"docker_version,omitempty" yaml:"docker_version,omitempty"`
	FullName        string            `json:"full_name,omitempty" yaml:"full_name,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	Repo            string            `json:"repo,omitempty" yaml:"repo,omitempty"`
	Size            int64             `json:"size,omitempty" yaml:"size,omitempty"`
	TagLabels       []string          `json:"tag_labels,omitempty" yaml:"tag_labels,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
}

type HarborTagCollection struct {
	types.Collection
	Data   []HarborTag `json:"data,omitempty"`
	client *HarborTagClient
}

type HarborTagClient struct {
	apiClient *Client
}

type HarborTagOperations interface {
	List(opts *types.ListOpts) (*HarborTagCollection, error)
	Create(opts *HarborTag) (*HarborTag, error)
	Update(existing *HarborTag, updates interface{}) (*HarborTag, error)
	Replace(existing *HarborTag) (*HarborTag, error)
	ByID(id string) (*HarborTag, error)
	Delete(container *HarborTag) error
}

func newHarborTagClient(apiClient *Client) *HarborTagClient {
	return &HarborTagClient{
		apiClient: apiClient,
	}
}

func (c *HarborTagClient) Create(container *HarborTag) (*HarborTag, error) {
	resp := &HarborTag{}
	err := c.apiClient.Ops.DoCreate(HarborTagType, container, resp)
	return resp, err
}

func (c *HarborTagClient) Update(existing *HarborTag, updates interface{}) (*HarborTag, error) {
	resp := &HarborTag{}
	err := c.apiClient.Ops.DoUpdate(HarborTagType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *HarborTagClient) Replace(obj *HarborTag) (*HarborTag, error) {
	resp := &HarborTag{}
	err := c.apiClient.Ops.DoReplace(HarborTagType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *HarborTagClient) List(opts *types.ListOpts) (*HarborTagCollection, error) {
	resp := &HarborTagCollection{}
	err := c.apiClient.Ops.DoList(HarborTagType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *HarborTagCollection) Next() (*HarborTagCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &HarborTagCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *HarborTagClient) ByID(id string) (*HarborTag, error) {
	resp := &HarborTag{}
	err := c.apiClient.Ops.DoByID(HarborTagType, id, resp)
	return resp, err
}

func (c *HarborTagClient) Delete(container *HarborTag) error {
	return c.apiClient.Ops.DoResourceDelete(HarborTagType, &container.Resource)
}
