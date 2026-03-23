// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateXskyBlockVolume updates XskyBlockVolume
func (cli *ZSClient) UpdateXskyBlockVolume(ctx context.Context, uuid string, params param.UpdateXskyBlockVolumeParam) (*view.BlockVolumeInventoryView, error) {
	resp := view.BlockVolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/xsky/block-volumes", uuid, "actions", "inventory", map[string]interface{}{
		"updateXskyBlockVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryXskyBlockVolume queries XskyBlockVolume list
func (cli *ZSClient) QueryXskyBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.XskyBlockVolumeInventoryView, error) {
	var resp []view.XskyBlockVolumeInventoryView
	return resp, cli.List(ctx, "v1/xksy/block-volumes", params, &resp)
}

func (cli *ZSClient) GetXskyBlockVolume(ctx context.Context, uuid string) (*view.XskyBlockVolumeInventoryView, error) {
	var resp view.XskyBlockVolumeInventoryView
	if err := cli.Get(ctx, "v1/xsky/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageXskyBlockVolume Pagination
func (cli *ZSClient) PageXskyBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.XskyBlockVolumeInventoryView, int, error) {
	var xskyBlockVolumes []view.XskyBlockVolumeInventoryView
	total, err := cli.Page(ctx, "v1/xksy/block-volumes", params, &xskyBlockVolumes)
	return xskyBlockVolumes, total, err
}
