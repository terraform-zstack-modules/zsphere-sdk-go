// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVirtualRouterVm queries VirtualRouterVm list
func (cli *ZSClient) QueryVirtualRouterVm(ctx context.Context, params *param.QueryParam) ([]view.ApplianceVmInventoryView, error) {
	var resp []view.ApplianceVmInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/appliances/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterVm(ctx context.Context, uuid string) (*view.ApplianceVmInventoryView, error) {
	var resp view.ApplianceVmInventoryView
	if err := cli.Get(ctx, "v1/vm-instances/appliances/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualRouterVm Pagination
func (cli *ZSClient) PageVirtualRouterVm(ctx context.Context, params *param.QueryParam) ([]view.ApplianceVmInventoryView, int, error) {
	var virtualRouterVms []view.ApplianceVmInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/appliances/virtual-routers", params, &virtualRouterVms)
	return virtualRouterVms, total, err
}
