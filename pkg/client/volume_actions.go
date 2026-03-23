// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVolume queries Volume list
func (cli *ZSClient) QueryVolume(ctx context.Context, params *param.QueryParam) ([]view.VolumeInventoryView, error) {
	var resp []view.VolumeInventoryView
	return resp, cli.List(ctx, "v1/volumes", params, &resp)
}

func (cli *ZSClient) GetVolume(ctx context.Context, uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.Get(ctx, "v1/volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolume Pagination
func (cli *ZSClient) PageVolume(ctx context.Context, params *param.QueryParam) ([]view.VolumeInventoryView, int, error) {
	var volumes []view.VolumeInventoryView
	total, err := cli.Page(ctx, "v1/volumes", params, &volumes)
	return volumes, total, err
}
// UpdateVolume updates Volume
func (cli *ZSClient) UpdateVolume(ctx context.Context, uuid string, params param.UpdateVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volumes", uuid, "actions", "inventory", map[string]interface{}{
		"updateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
