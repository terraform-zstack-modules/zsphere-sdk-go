// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryExponBlockVolume queries ExponBlockVolume list
func (cli *ZSClient) QueryExponBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.ExponBlockVolumeInventoryView, error) {
	var resp []view.ExponBlockVolumeInventoryView
	return resp, cli.List(ctx, "v1/expon/block-volumes", params, &resp)
}

func (cli *ZSClient) GetExponBlockVolume(ctx context.Context, uuid string) (*view.ExponBlockVolumeInventoryView, error) {
	var resp view.ExponBlockVolumeInventoryView
	if err := cli.Get(ctx, "v1/expon/block-volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageExponBlockVolume Pagination
func (cli *ZSClient) PageExponBlockVolume(ctx context.Context, params *param.QueryParam) ([]view.ExponBlockVolumeInventoryView, int, error) {
	var exponBlockVolumes []view.ExponBlockVolumeInventoryView
	total, err := cli.Page(ctx, "v1/expon/block-volumes", params, &exponBlockVolumes)
	return exponBlockVolumes, total, err
}
