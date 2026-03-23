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

// AddZStone adds ZStone
func (cli *ZSClient) AddZStone(ctx context.Context, params param.AddZStoneParam) (*view.ZStoneInventoryView, error) {
	resp := view.ZStoneInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zstone-plugin"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZStone queries ZStone list
func (cli *ZSClient) QueryZStone(ctx context.Context, params *param.QueryParam) ([]view.ZStoneInventoryView, error) {
	var resp []view.ZStoneInventoryView
	return resp, cli.List(ctx, "v1/zstone-plugin", params, &resp)
}

func (cli *ZSClient) GetZStone(ctx context.Context, uuid string) (*view.ZStoneInventoryView, error) {
	var resp view.ZStoneInventoryView
	if err := cli.Get(ctx, "v1/zstone-plugin", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZStone Pagination
func (cli *ZSClient) PageZStone(ctx context.Context, params *param.QueryParam) ([]view.ZStoneInventoryView, int, error) {
	var zStones []view.ZStoneInventoryView
	total, err := cli.Page(ctx, "v1/zstone-plugin", params, &zStones)
	return zStones, total, err
}
// RemoveZStone removes ZStone
func (cli *ZSClient) RemoveZStone(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zstone-plugin", uuid, string(deleteMode))
}
