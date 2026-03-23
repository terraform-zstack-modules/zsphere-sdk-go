// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVCenterPrimaryStorage queries VCenterPrimaryStorage list
func (cli *ZSClient) QueryVCenterPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.VCenterPrimaryStorageInventoryView, error) {
	var resp []view.VCenterPrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/vcenters/primary-storage", params, &resp)
}

func (cli *ZSClient) GetVCenterPrimaryStorage(ctx context.Context, uuid string) (*view.VCenterPrimaryStorageInventoryView, error) {
	var resp view.VCenterPrimaryStorageInventoryView
	if err := cli.Get(ctx, "v1/vcenters/primary-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterPrimaryStorage Pagination
func (cli *ZSClient) PageVCenterPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.VCenterPrimaryStorageInventoryView, int, error) {
	var vCenterPrimaryStorages []view.VCenterPrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/vcenters/primary-storage", params, &vCenterPrimaryStorages)
	return vCenterPrimaryStorages, total, err
}
