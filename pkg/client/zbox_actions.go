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

// AddZBox adds ZBox
func (cli *ZSClient) AddZBox(ctx context.Context, params param.AddZBoxParam) (*view.ZBoxInventoryView, error) {
	resp := view.ZBoxInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zbox"), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZBox queries ZBox list
func (cli *ZSClient) QueryZBox(ctx context.Context, params *param.QueryParam) ([]view.ZBoxInventoryView, error) {
	var resp []view.ZBoxInventoryView
	return resp, cli.List(ctx, "v1/zbox", params, &resp)
}

func (cli *ZSClient) GetZBox(ctx context.Context, uuid string) (*view.ZBoxInventoryView, error) {
	var resp view.ZBoxInventoryView
	if err := cli.Get(ctx, "v1/zbox", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZBox Pagination
func (cli *ZSClient) PageZBox(ctx context.Context, params *param.QueryParam) ([]view.ZBoxInventoryView, int, error) {
	var zBoxs []view.ZBoxInventoryView
	total, err := cli.Page(ctx, "v1/zbox", params, &zBoxs)
	return zBoxs, total, err
}
