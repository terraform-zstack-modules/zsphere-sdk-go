// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"fmt"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// CreateAccessKey creates AccessKey
func (cli *ZSClient) CreateAccessKey(ctx context.Context, params param.CreateAccessKeyParam) (*view.AccessKeyInventoryView, error) {
	resp := view.AccessKeyInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/accesskeys"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccessKey queries AccessKey list
func (cli *ZSClient) QueryAccessKey(ctx context.Context, params *param.QueryParam) ([]view.AccessKeyInventoryView, error) {
	var resp []view.AccessKeyInventoryView
	return resp, cli.List(ctx, "v1/accesskeys", params, &resp)
}

func (cli *ZSClient) GetAccessKey(ctx context.Context, uuid string) (*view.AccessKeyInventoryView, error) {
	var resp view.AccessKeyInventoryView
	if err := cli.Get(ctx, "v1/accesskeys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccessKey Pagination
func (cli *ZSClient) PageAccessKey(ctx context.Context, params *param.QueryParam) ([]view.AccessKeyInventoryView, int, error) {
	var accessKeys []view.AccessKeyInventoryView
	total, err := cli.Page(ctx, "v1/accesskeys", params, &accessKeys)
	return accessKeys, total, err
}
// DeleteAccessKey deletes AccessKey
func (cli *ZSClient) DeleteAccessKey(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/accesskeys", uuid, string(deleteMode))
}
