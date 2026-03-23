// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// CreateL2VxlanNetworkPool creates L2VxlanNetworkPool
func (cli *ZSClient) CreateL2VxlanNetworkPool(ctx context.Context) (*view.L2VxlanNetworkPoolInventoryView, error) {
	resp := view.L2VxlanNetworkPoolInventoryView{}
	if err := cli.Post(ctx, "v1/l2-networks/vxlan-pool", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2VxlanNetworkPool queries L2VxlanNetworkPool list
func (cli *ZSClient) QueryL2VxlanNetworkPool(ctx context.Context, params *param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, error) {
	var resp []view.L2VxlanNetworkPoolInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/vxlan-pool", params, &resp)
}

func (cli *ZSClient) GetL2VxlanNetworkPool(ctx context.Context, uuid string) (*view.L2VxlanNetworkPoolInventoryView, error) {
	var resp view.L2VxlanNetworkPoolInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/vxlan-pool", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2VxlanNetworkPool Pagination
func (cli *ZSClient) PageL2VxlanNetworkPool(ctx context.Context, params *param.QueryParam) ([]view.L2VxlanNetworkPoolInventoryView, int, error) {
	var l2VxlanNetworkPools []view.L2VxlanNetworkPoolInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/vxlan-pool", params, &l2VxlanNetworkPools)
	return l2VxlanNetworkPools, total, err
}
