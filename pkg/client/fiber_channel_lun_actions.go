// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryFiberChannelLun queries FiberChannelLun list
func (cli *ZSClient) QueryFiberChannelLun(ctx context.Context, params *param.QueryParam) ([]view.FiberChannelLunInventoryView, error) {
	var resp []view.FiberChannelLunInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/fiber-channel/luns", params, &resp)
}

func (cli *ZSClient) GetFiberChannelLun(ctx context.Context, uuid string) (*view.FiberChannelLunInventoryView, error) {
	var resp view.FiberChannelLunInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/fiber-channel/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFiberChannelLun Pagination
func (cli *ZSClient) PageFiberChannelLun(ctx context.Context, params *param.QueryParam) ([]view.FiberChannelLunInventoryView, int, error) {
	var fiberChannelLuns []view.FiberChannelLunInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/fiber-channel/luns", params, &fiberChannelLuns)
	return fiberChannelLuns, total, err
}
