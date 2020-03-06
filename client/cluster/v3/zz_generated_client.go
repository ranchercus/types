package client

import (
	"github.com/rancher/norman/clientbase"
)

type Client struct {
	clientbase.APIBaseClient

	Namespace        NamespaceOperations
	PersistentVolume PersistentVolumeOperations
	StorageClass     StorageClassOperations
	APIService       APIServiceOperations
	HarborProject    HarborProjectOperations
	HarborRepository HarborRepositoryOperations
	HarborTag        HarborTagOperations
}

func NewClient(opts *clientbase.ClientOpts) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		APIBaseClient: baseClient,
	}

	client.Namespace = newNamespaceClient(client)
	client.PersistentVolume = newPersistentVolumeClient(client)
	client.StorageClass = newStorageClassClient(client)
	client.APIService = newAPIServiceClient(client)
	client.HarborProject = newHarborProjectClient(client)
	client.HarborRepository = newHarborRepositoryClient(client)
	client.HarborTag = newHarborTagClient(client)

	return client, nil
}
