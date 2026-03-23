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

// UpdateBlockPrimaryStorage updates BlockPrimaryStorage
func (cli *ZSClient) UpdateBlockPrimaryStorage(ctx context.Context, uuid string, params param.UpdateBlockPrimaryStorageParam) (*view.BlockPrimaryStorageInventoryView, error) {
	resp := view.BlockPrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/block", uuid, "actions", "inventory", map[string]interface{}{
		"updateBlockPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddBlockPrimaryStorage adds BlockPrimaryStorage
func (cli *ZSClient) AddBlockPrimaryStorage(ctx context.Context, params param.AddBlockPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/block"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBlockPrimaryStorage queries BlockPrimaryStorage list
func (cli *ZSClient) QueryBlockPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, error) {
	var resp []view.BlockPrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/block", params, &resp)
}

func (cli *ZSClient) GetBlockPrimaryStorage(ctx context.Context, uuid string) (*view.BlockPrimaryStorageInventoryView, error) {
	var resp view.BlockPrimaryStorageInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/block", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBlockPrimaryStorage Pagination
func (cli *ZSClient) PageBlockPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.BlockPrimaryStorageInventoryView, int, error) {
	var blockPrimaryStorages []view.BlockPrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/block", params, &blockPrimaryStorages)
	return blockPrimaryStorages, total, err
}
