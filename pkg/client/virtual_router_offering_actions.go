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

// QueryVirtualRouterOffering queries VirtualRouterOffering list
func (cli *ZSClient) QueryVirtualRouterOffering(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, error) {
	var resp []view.VirtualRouterOfferingInventoryView
	return resp, cli.List(ctx, "v1/instance-offerings/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetVirtualRouterOffering(ctx context.Context, uuid string) (*view.VirtualRouterOfferingInventoryView, error) {
	var resp view.VirtualRouterOfferingInventoryView
	if err := cli.Get(ctx, "v1/instance-offerings/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVirtualRouterOffering Pagination
func (cli *ZSClient) PageVirtualRouterOffering(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterOfferingInventoryView, int, error) {
	var virtualRouterOfferings []view.VirtualRouterOfferingInventoryView
	total, err := cli.Page(ctx, "v1/instance-offerings/virtual-routers", params, &virtualRouterOfferings)
	return virtualRouterOfferings, total, err
}
// CreateVirtualRouterOffering creates VirtualRouterOffering
func (cli *ZSClient) CreateVirtualRouterOffering(ctx context.Context, params param.CreateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/instance-offerings/virtual-routers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVirtualRouterOffering updates VirtualRouterOffering
func (cli *ZSClient) UpdateVirtualRouterOffering(ctx context.Context, uuid string, params param.UpdateVirtualRouterOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/instance-offerings/virtual-routers", uuid, "actions", "inventory", map[string]interface{}{
		"updateVirtualRouterOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
