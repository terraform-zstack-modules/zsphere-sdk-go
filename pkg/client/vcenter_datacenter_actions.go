// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVCenterDatacenter queries VCenterDatacenter list
func (cli *ZSClient) QueryVCenterDatacenter(ctx context.Context, params *param.QueryParam) ([]view.VCenterDatacenterInventoryView, error) {
	var resp []view.VCenterDatacenterInventoryView
	return resp, cli.List(ctx, "v1/vcenters/datacenters", params, &resp)
}

func (cli *ZSClient) GetVCenterDatacenter(ctx context.Context, uuid string) (*view.VCenterDatacenterInventoryView, error) {
	var resp view.VCenterDatacenterInventoryView
	if err := cli.Get(ctx, "v1/vcenters/datacenters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterDatacenter Pagination
func (cli *ZSClient) PageVCenterDatacenter(ctx context.Context, params *param.QueryParam) ([]view.VCenterDatacenterInventoryView, int, error) {
	var vCenterDatacenters []view.VCenterDatacenterInventoryView
	total, err := cli.Page(ctx, "v1/vcenters/datacenters", params, &vCenterDatacenters)
	return vCenterDatacenters, total, err
}
