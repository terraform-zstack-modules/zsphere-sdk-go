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

// QuerySharedBlock queries SharedBlock list
func (cli *ZSClient) QuerySharedBlock(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockInventoryView, error) {
	var resp []view.SharedBlockInventoryView
	return resp, cli.List(ctx, "v1/sharedblock-group/sharedblocks", params, &resp)
}

func (cli *ZSClient) GetSharedBlock(ctx context.Context, uuid string) (*view.SharedBlockInventoryView, error) {
	var resp view.SharedBlockInventoryView
	if err := cli.Get(ctx, "v1/sharedblock-group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSharedBlock Pagination
func (cli *ZSClient) PageSharedBlock(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockInventoryView, int, error) {
	var sharedBlocks []view.SharedBlockInventoryView
	total, err := cli.Page(ctx, "v1/sharedblock-group/sharedblocks", params, &sharedBlocks)
	return sharedBlocks, total, err
}
// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(ctx context.Context, sharedBlockGroupUuid string, uuid string, params param.UpdateSharedBlockParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/primary-storage/sharedblockgroup", sharedBlockGroupUuid, fmt.Sprintf("sharedblocks/%s/actions", uuid), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
