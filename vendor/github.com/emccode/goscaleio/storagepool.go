package goscaleio

import (
	"errors"
	"fmt"

	types "github.com/emccode/goscaleio/types/v1"
)

type StoragePool struct {
	StoragePool *types.StoragePool
	client      *Client
}

func NewStoragePool(client *Client) *StoragePool {
	return &StoragePool{
		StoragePool: new(types.StoragePool),
		client:      client,
	}
}

func (protectionDomain *ProtectionDomain) GetStoragePool(storagepoolhref string) (storagePools []*types.StoragePool, err error) {

	endpoint := protectionDomain.client.SIOEndpoint

	if storagepoolhref == "" {
		link, err := GetLink(protectionDomain.ProtectionDomain.Links, "/api/ProtectionDomain/relationship/StoragePool")
		if err != nil {
			return []*types.StoragePool{}, errors.New("Error: problem finding link")
		}
		endpoint.Path = link.HREF
	} else {
		endpoint.Path = storagepoolhref
	}

	req := protectionDomain.client.NewRequest(map[string]string{}, "GET", endpoint, nil)
	req.SetBasicAuth("", protectionDomain.client.Token)
	req.Header.Add("Accept", "application/json;version="+protectionDomain.client.configConnect.Version)

	resp, err := protectionDomain.client.retryCheckResp(&protectionDomain.client.Http, req)
	if err != nil {
		return []*types.StoragePool{}, fmt.Errorf("problem getting response: %v", err)
	}
	defer resp.Body.Close()

	if storagepoolhref == "" {
		if err = protectionDomain.client.decodeBody(resp, &storagePools); err != nil {
			return []*types.StoragePool{}, fmt.Errorf("error decoding storage pool response: %s", err)
		}
	} else {
		storagePool := &types.StoragePool{}
		if err = protectionDomain.client.decodeBody(resp, &storagePool); err != nil {
			return []*types.StoragePool{}, fmt.Errorf("error decoding instances response: %s", err)
		}
		storagePools = append(storagePools, storagePool)
	}
	return storagePools, nil
}

func (protectionDomain *ProtectionDomain) FindStoragePool(id, name, href string) (storagePool *types.StoragePool, err error) {
	storagePools, err := protectionDomain.GetStoragePool(href)
	if err != nil {
		return &types.StoragePool{}, fmt.Errorf("Error getting protection domains %s", err)
	}

	for _, storagePool = range storagePools {
		if storagePool.ID == id || storagePool.Name == name || href != "" {
			return storagePool, nil
		}
	}

	return &types.StoragePool{}, errors.New("Couldn't find protection domain")

}
