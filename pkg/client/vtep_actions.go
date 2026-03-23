// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVtep queries Vtep list
func (cli *ZSClient) QueryVtep(ctx context.Context, params *param.QueryParam) ([]view.VtepInventoryView, error) {
	var resp []view.VtepInventoryView
	return resp, cli.List(ctx, "v1/l2-networks/vteps", params, &resp)
}

func (cli *ZSClient) GetVtep(ctx context.Context, uuid string) (*view.VtepInventoryView, error) {
	var resp view.VtepInventoryView
	if err := cli.Get(ctx, "v1/l2-networks/vteps", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVtep Pagination
func (cli *ZSClient) PageVtep(ctx context.Context, params *param.QueryParam) ([]view.VtepInventoryView, int, error) {
	var vteps []view.VtepInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks/vteps", params, &vteps)
	return vteps, total, err
}
