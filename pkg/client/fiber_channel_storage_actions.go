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

// RefreshFiberChannelStorage operates on FiberChannelStorage
func (cli *ZSClient) RefreshFiberChannelStorage(ctx context.Context, params param.RefreshFiberChannelStorageParam) (*view.FiberChannelStorageInventoryView, error) {
	resp := view.FiberChannelStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/storage-devices/fiber-channel/controllers"), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryFiberChannelStorage queries FiberChannelStorage list
func (cli *ZSClient) QueryFiberChannelStorage(ctx context.Context, params *param.QueryParam) ([]view.FiberChannelStorageInventoryView, error) {
	var resp []view.FiberChannelStorageInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/fiber-channel/controllers", params, &resp)
}

func (cli *ZSClient) GetFiberChannelStorage(ctx context.Context, uuid string) (*view.FiberChannelStorageInventoryView, error) {
	var resp view.FiberChannelStorageInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/fiber-channel", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFiberChannelStorage Pagination
func (cli *ZSClient) PageFiberChannelStorage(ctx context.Context, params *param.QueryParam) ([]view.FiberChannelStorageInventoryView, int, error) {
	var fiberChannelStorages []view.FiberChannelStorageInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/fiber-channel/controllers", params, &fiberChannelStorages)
	return fiberChannelStorages, total, err
}
