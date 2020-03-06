package client

import (
	"github.com/rancher/norman/types"
)

const (
	ClusterSettingType                 = "clusterSetting"
	ClusterSettingFieldAnnotations     = "annotations"
	ClusterSettingFieldCreated         = "created"
	ClusterSettingFieldCreatorID       = "creatorId"
	ClusterSettingFieldLabels          = "labels"
	ClusterSettingFieldName            = "name"
	ClusterSettingFieldOwnerReferences = "ownerReferences"
	ClusterSettingFieldPipelineSetting = "pipelineSetting"
	ClusterSettingFieldRegistrySetting = "registrySetting"
	ClusterSettingFieldRemoved         = "removed"
	ClusterSettingFieldUUID            = "uuid"
	ClusterSettingFieldWorkloadSetting = "workloadSetting"
)

type ClusterSetting struct {
	types.Resource
	Annotations     map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Created         string            `json:"created,omitempty" yaml:"created,omitempty"`
	CreatorID       string            `json:"creatorId,omitempty" yaml:"creatorId,omitempty"`
	Labels          map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Name            string            `json:"name,omitempty" yaml:"name,omitempty"`
	OwnerReferences []OwnerReference  `json:"ownerReferences,omitempty" yaml:"ownerReferences,omitempty"`
	PipelineSetting *PipelineSetting  `json:"pipelineSetting,omitempty" yaml:"pipelineSetting,omitempty"`
	RegistrySetting *RegistrySetting  `json:"registrySetting,omitempty" yaml:"registrySetting,omitempty"`
	Removed         string            `json:"removed,omitempty" yaml:"removed,omitempty"`
	UUID            string            `json:"uuid,omitempty" yaml:"uuid,omitempty"`
	WorkloadSetting *WorkloadSetting  `json:"workloadSetting,omitempty" yaml:"workloadSetting,omitempty"`
}

type ClusterSettingCollection struct {
	types.Collection
	Data   []ClusterSetting `json:"data,omitempty"`
	client *ClusterSettingClient
}

type ClusterSettingClient struct {
	apiClient *Client
}

type ClusterSettingOperations interface {
	List(opts *types.ListOpts) (*ClusterSettingCollection, error)
	Create(opts *ClusterSetting) (*ClusterSetting, error)
	Update(existing *ClusterSetting, updates interface{}) (*ClusterSetting, error)
	Replace(existing *ClusterSetting) (*ClusterSetting, error)
	ByID(id string) (*ClusterSetting, error)
	Delete(container *ClusterSetting) error
}

func newClusterSettingClient(apiClient *Client) *ClusterSettingClient {
	return &ClusterSettingClient{
		apiClient: apiClient,
	}
}

func (c *ClusterSettingClient) Create(container *ClusterSetting) (*ClusterSetting, error) {
	resp := &ClusterSetting{}
	err := c.apiClient.Ops.DoCreate(ClusterSettingType, container, resp)
	return resp, err
}

func (c *ClusterSettingClient) Update(existing *ClusterSetting, updates interface{}) (*ClusterSetting, error) {
	resp := &ClusterSetting{}
	err := c.apiClient.Ops.DoUpdate(ClusterSettingType, &existing.Resource, updates, resp)
	return resp, err
}

func (c *ClusterSettingClient) Replace(obj *ClusterSetting) (*ClusterSetting, error) {
	resp := &ClusterSetting{}
	err := c.apiClient.Ops.DoReplace(ClusterSettingType, &obj.Resource, obj, resp)
	return resp, err
}

func (c *ClusterSettingClient) List(opts *types.ListOpts) (*ClusterSettingCollection, error) {
	resp := &ClusterSettingCollection{}
	err := c.apiClient.Ops.DoList(ClusterSettingType, opts, resp)
	resp.client = c
	return resp, err
}

func (cc *ClusterSettingCollection) Next() (*ClusterSettingCollection, error) {
	if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
		resp := &ClusterSettingCollection{}
		err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
		resp.client = cc.client
		return resp, err
	}
	return nil, nil
}

func (c *ClusterSettingClient) ByID(id string) (*ClusterSetting, error) {
	resp := &ClusterSetting{}
	err := c.apiClient.Ops.DoByID(ClusterSettingType, id, resp)
	return resp, err
}

func (c *ClusterSettingClient) Delete(container *ClusterSetting) error {
	return c.apiClient.Ops.DoResourceDelete(ClusterSettingType, &container.Resource)
}
