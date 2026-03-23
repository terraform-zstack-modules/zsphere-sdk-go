// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVCenterResourcePool queries VCenterResourcePool list
func (cli *ZSClient) QueryVCenterResourcePool(ctx context.Context, params *param.QueryParam) ([]view.VCenterResourcePoolInventoryView, error) {
	var resp []view.VCenterResourcePoolInventoryView
	return resp, cli.List(ctx, "v1/vcenters/clusters/resourcepools", params, &resp)
}

func (cli *ZSClient) GetVCenterResourcePool(ctx context.Context, uuid string) (*view.VCenterResourcePoolInventoryView, error) {
	var resp view.VCenterResourcePoolInventoryView
	if err := cli.Get(ctx, "v1/vcenters/clusters/resourcepools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterResourcePool Pagination
func (cli *ZSClient) PageVCenterResourcePool(ctx context.Context, params *param.QueryParam) ([]view.VCenterResourcePoolInventoryView, int, error) {
	var vCenterResourcePools []view.VCenterResourcePoolInventoryView
	total, err := cli.Page(ctx, "v1/vcenters/clusters/resourcepools", params, &vCenterResourcePools)
	return vCenterResourcePools, total, err
}
