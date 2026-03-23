// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVirtualRouterVRouterRouteTableRef queries VirtualRouterVRouterRouteTableRef list
func (cli *ZSClient) QueryVirtualRouterVRouterRouteTableRef(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, error) {
	var resp []view.VirtualRouterVRouterRouteTableRefInventoryView
	return resp, cli.List(ctx, "v1/vrouter-route-tables/virtual-router-refs", params, &resp)
}

// PageVirtualRouterVRouterRouteTableRef Pagination
func (cli *ZSClient) PageVirtualRouterVRouterRouteTableRef(ctx context.Context, params *param.QueryParam) ([]view.VirtualRouterVRouterRouteTableRefInventoryView, int, error) {
	var virtualRouterVRouterRouteTableRefs []view.VirtualRouterVRouterRouteTableRefInventoryView
	total, err := cli.Page(ctx, "v1/vrouter-route-tables/virtual-router-refs", params, &virtualRouterVRouterRouteTableRefs)
	return virtualRouterVRouterRouteTableRefs, total, err
}
