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

// QueryL2VlanNetwork queries L2VlanNetwork list
func (cli *ZSClient) QueryL2VlanNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VlanNetworkInventoryView, error) {
	var resp []view.L2VlanNetworkInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/vlan", params, &resp)
}

func (cli *ZSClient) GetL2VlanNetwork(ctx context.Context, uuid string) (*view.L2VlanNetworkInventoryView, error) {
	var resp view.L2VlanNetworkInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/vlan", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2VlanNetwork Pagination
func (cli *ZSClient) PageL2VlanNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VlanNetworkInventoryView, int, error) {
	var l2VlanNetworks []view.L2VlanNetworkInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/vlan", params, &l2VlanNetworks)
	return l2VlanNetworks, total, err
}
// CreateL2VlanNetwork creates L2VlanNetwork
func (cli *ZSClient) CreateL2VlanNetwork(ctx context.Context, params param.CreateL2VlanNetworkParam) (*view.L2VlanNetworkInventoryView, error) {
	resp := view.L2VlanNetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/vlan"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
