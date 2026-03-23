// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryL2PortGroupNetwork queries L2PortGroupNetwork list
func (cli *ZSClient) QueryL2PortGroupNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, error) {
	var resp []view.L2PortGroupNetworkInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/port-group", params, &resp)
}

func (cli *ZSClient) GetL2PortGroupNetwork(ctx context.Context, uuid string) (*view.L2PortGroupNetworkInventoryView, error) {
	var resp view.L2PortGroupNetworkInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/port-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2PortGroupNetwork Pagination
func (cli *ZSClient) PageL2PortGroupNetwork(ctx context.Context, params *param.QueryParam) ([]view.L2PortGroupNetworkInventoryView, int, error) {
	var l2PortGroupNetworks []view.L2PortGroupNetworkInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/port-group", params, &l2PortGroupNetworks)
	return l2PortGroupNetworks, total, err
}
