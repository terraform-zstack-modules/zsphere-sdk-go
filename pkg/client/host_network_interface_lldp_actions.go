// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryHostNetworkInterfaceLldp queries HostNetworkInterfaceLldp list
func (cli *ZSClient) QueryHostNetworkInterfaceLldp(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp []view.HostNetworkInterfaceLldpInventoryView
	return resp, cli.List(ctx, "v1/hostNetworkInterface/lldp/all", params, &resp)
}

func (cli *ZSClient) GetHostNetworkInterfaceLldp(ctx context.Context, uuid string) (*view.HostNetworkInterfaceLldpInventoryView, error) {
	var resp view.HostNetworkInterfaceLldpInventoryView
	if err := cli.Get(ctx, "v1/hostNetworkInterface/lldp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostNetworkInterfaceLldp Pagination
func (cli *ZSClient) PageHostNetworkInterfaceLldp(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceLldpInventoryView, int, error) {
	var hostNetworkInterfaceLldps []view.HostNetworkInterfaceLldpInventoryView
	total, err := cli.Page(ctx, "v1/hostNetworkInterface/lldp/all", params, &hostNetworkInterfaceLldps)
	return hostNetworkInterfaceLldps, total, err
}
