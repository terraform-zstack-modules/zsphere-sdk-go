// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryUplinkGroup queries UplinkGroup list
func (cli *ZSClient) QueryUplinkGroup(ctx context.Context, params *param.QueryParam) ([]view.UplinkGroupInventoryView, error) {
	var resp []view.UplinkGroupInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/virtual-switch/uplink-group", params, &resp)
}

func (cli *ZSClient) GetUplinkGroup(ctx context.Context, uuid string) (*view.UplinkGroupInventoryView, error) {
	var resp view.UplinkGroupInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/virtual-switch/uplink-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUplinkGroup Pagination
func (cli *ZSClient) PageUplinkGroup(ctx context.Context, params *param.QueryParam) ([]view.UplinkGroupInventoryView, int, error) {
	var uplinkGroups []view.UplinkGroupInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/virtual-switch/uplink-group", params, &uplinkGroups)
	return uplinkGroups, total, err
}
