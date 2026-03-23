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

// CreateMulticastRouter creates MulticastRouter
func (cli *ZSClient) CreateMulticastRouter(ctx context.Context, params param.CreateMulticastRouterParam) (*view.MulticastRouterInventoryView, error) {
	resp := view.MulticastRouterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/multicast/virtual-routers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMulticastRouter queries MulticastRouter list
func (cli *ZSClient) QueryMulticastRouter(ctx context.Context, params *param.QueryParam) ([]view.MulticastRouterInventoryView, error) {
	var resp []view.MulticastRouterInventoryView
	return resp, cli.List(ctx, "v1/multicast/virtual-routers", params, &resp)
}

func (cli *ZSClient) GetMulticastRouter(ctx context.Context, uuid string) (*view.MulticastRouterInventoryView, error) {
	var resp view.MulticastRouterInventoryView
	if err := cli.Get(ctx, "v1/multicast/virtual-routers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMulticastRouter Pagination
func (cli *ZSClient) PageMulticastRouter(ctx context.Context, params *param.QueryParam) ([]view.MulticastRouterInventoryView, int, error) {
	var multicastRouters []view.MulticastRouterInventoryView
	total, err := cli.Page(ctx, "v1/multicast/virtual-routers", params, &multicastRouters)
	return multicastRouters, total, err
}
// DeleteMulticastRouter deletes MulticastRouter
func (cli *ZSClient) DeleteMulticastRouter(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/multicast/virtual-routers", uuid, string(deleteMode))
}
