// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryL2VirtualSwitchNetwork queries L2VirtualSwitchNetwork list
func (cli *ZSClient) QueryL2VirtualSwitchNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, error) {
	var resp []view.L2VirtualSwitchNetworkInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/virtual-switch", params, &resp)
}

func (cli *ZSClient) GetL2VirtualSwitchNetwork(ctx context.Context, uuid string) (*view.L2VirtualSwitchNetworkInventoryView, error) {
	var resp view.L2VirtualSwitchNetworkInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/virtual-switch", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2VirtualSwitchNetwork Pagination
func (cli *ZSClient) PageL2VirtualSwitchNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2VirtualSwitchNetworkInventoryView, int, error) {
	var l2VirtualSwitchNetworks []view.L2VirtualSwitchNetworkInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/virtual-switch", params, &l2VirtualSwitchNetworks)
	return l2VirtualSwitchNetworks, total, err
}
