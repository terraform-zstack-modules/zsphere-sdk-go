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

// RemoveZceX removes ZceX
func (cli *ZSClient) RemoveZceX(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zce-x-plugin", uuid, string(deleteMode))
}
// AddZceX adds ZceX
func (cli *ZSClient) AddZceX(ctx context.Context, params param.AddZceXParam) (*view.ZceXInventoryView, error) {
	resp := view.ZceXInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zce-x-plugin"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZceX queries ZceX list
func (cli *ZSClient) QueryZceX(ctx context.Context, params *param.QueryParam) ([]view.ZceXInventoryView, error) {
	var resp []view.ZceXInventoryView
	return resp, cli.List(ctx, "v1/zce-x-plugin", params, &resp)
}

func (cli *ZSClient) GetZceX(ctx context.Context, uuid string) (*view.ZceXInventoryView, error) {
	var resp view.ZceXInventoryView
	if err := cli.Get(ctx, "v1/zce-x-plugin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZceX Pagination
func (cli *ZSClient) PageZceX(ctx context.Context, params *param.QueryParam) ([]view.ZceXInventoryView, int, error) {
	var zceXs []view.ZceXInventoryView
	total, err := cli.Page(ctx, "v1/zce-x-plugin", params, &zceXs)
	return zceXs, total, err
}
