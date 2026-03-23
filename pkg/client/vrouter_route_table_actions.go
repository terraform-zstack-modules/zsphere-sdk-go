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

// QueryVRouterRouteTable queries VRouterRouteTable list
func (cli *ZSClient) QueryVRouterRouteTable(ctx context.Context, params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, error) {
	var resp []view.VRouterRouteTableInventoryView
	return resp, cli.List(ctx, "v1/vrouter-route-tables", params, &resp)
}

func (cli *ZSClient) GetVRouterRouteTable(ctx context.Context, uuid string) (*view.VRouterRouteTableInventoryView, error) {
	var resp view.VRouterRouteTableInventoryView
	if err := cli.Get(ctx, "v1/vrouter-route-tables", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVRouterRouteTable Pagination
func (cli *ZSClient) PageVRouterRouteTable(ctx context.Context, params *param.QueryParam) ([]view.VRouterRouteTableInventoryView, int, error) {
	var vRouterRouteTables []view.VRouterRouteTableInventoryView
	total, err := cli.Page(ctx, "v1/vrouter-route-tables", params, &vRouterRouteTables)
	return vRouterRouteTables, total, err
}
// CreateVRouterRouteTable creates VRouterRouteTable
func (cli *ZSClient) CreateVRouterRouteTable(ctx context.Context, params param.CreateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vrouter-route-tables"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVRouterRouteTable updates VRouterRouteTable
func (cli *ZSClient) UpdateVRouterRouteTable(ctx context.Context, uuid string, params param.UpdateVRouterRouteTableParam) (*view.VRouterRouteTableInventoryView, error) {
	resp := view.VRouterRouteTableInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vrouter-route-tables", uuid, "actions", "inventory", map[string]interface{}{
		"updateVRouterRouteTable": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVRouterRouteTable deletes VRouterRouteTable
func (cli *ZSClient) DeleteVRouterRouteTable(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vrouter-route-tables", uuid, string(deleteMode))
}
