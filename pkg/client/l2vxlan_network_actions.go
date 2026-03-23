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

// CreateL2VxlanNetwork creates L2VxlanNetwork
func (cli *ZSClient) CreateL2VxlanNetwork(ctx context.Context, params param.CreateL2VxlanNetworkParam) (*view.L2VxlanNetworkInventoryView, error) {
	resp := view.L2VxlanNetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l2-networks/vxlan"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2VxlanNetwork queries L2VxlanNetwork list
func (cli *ZSClient) QueryL2VxlanNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VxlanNetworkInventoryView, error) {
	var resp []view.L2VxlanNetworkInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/vxlan", params, &resp)
}

func (cli *ZSClient) GetL2VxlanNetwork(ctx context.Context, uuid string) (*view.L2VxlanNetworkInventoryView, error) {
	var resp view.L2VxlanNetworkInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/vxlan", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2VxlanNetwork Pagination
func (cli *ZSClient) PageL2VxlanNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VxlanNetworkInventoryView, int, error) {
	var l2VxlanNetworks []view.L2VxlanNetworkInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/vxlan", params, &l2VxlanNetworks)
	return l2VxlanNetworks, total, err
}
