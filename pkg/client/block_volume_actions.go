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

// UpdateBlockVolume updates BlockVolume
func (cli *ZSClient) UpdateBlockVolume(ctx context.Context, uuid string, params param.UpdateBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	resp := view.BlockVolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/block-volumes", uuid, "actions", "inventory", map[string]interface{}{
		"updateBlockVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateBlockVolume creates BlockVolume
func (cli *ZSClient) CreateBlockVolume(ctx context.Context, params param.CreateBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	resp := view.BlockVolumeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/block-volumes"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBlockVolume queries BlockVolume list
func (cli *ZSClient) QueryBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.BlockVolumeInventoryView, error) {
	var resp []view.BlockVolumeInventoryView
	return resp, cli.List(ctx, "v1/block-volumes", params, &resp)
}

func (cli *ZSClient) GetBlockVolume(ctx context.Context, uuid string) (*view.BlockVolumeInventoryView, error) {
	var resp view.BlockVolumeInventoryView
	if err := cli.Get(ctx, "v1/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBlockVolume Pagination
func (cli *ZSClient) PageBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.BlockVolumeInventoryView, int, error) {
	var blockVolumes []view.BlockVolumeInventoryView
	total, err := cli.Page(ctx, "v1/block-volumes", params, &blockVolumes)
	return blockVolumes, total, err
}
