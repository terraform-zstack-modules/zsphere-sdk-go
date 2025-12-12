package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryAccessKeys Query access keys
func (cli *ZSClient) QueryAccessKeys(params param.QueryParam) ([]view.AccessKeyInventoryView, error) {
	var resp []view.AccessKeyInventoryView
	return resp, cli.List("v1/accesskeys", &params, &resp)
}

// GetAccessKey retrieves an access key by UUID.
func (cli *ZSClient) GetAccessKey(uuid string) (*view.AccessKeyInventoryView, error) {
	var resp view.AccessKeyInventoryView
	if err := cli.Get("v1/accesskeys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAccessKey creates an access key
func (cli *ZSClient) CreateAccessKey(params *param.CreateAccessKeyParam) (*view.AccessKeyInventoryView, error) {
	var resp view.AccessKeyInventoryView
	if err := cli.Post("v1/accesskeys", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAccessKey deletes an access key
func (cli *ZSClient) DeleteAccessKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accesskeys", uuid, string(deleteMode))
}
