// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryFcHbaDevice queries FcHbaDevice list
func (cli *ZSClient) QueryFcHbaDevice(ctx context.Context, params *param.QueryParam) ([]view.HbaDeviceInventoryView, error) {
	var resp []view.HbaDeviceInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/hba", params, &resp)
}

func (cli *ZSClient) GetFcHbaDevice(ctx context.Context, uuid string) (*view.HbaDeviceInventoryView, error) {
	var resp view.HbaDeviceInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/hba", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFcHbaDevice Pagination
func (cli *ZSClient) PageFcHbaDevice(ctx context.Context, params *param.QueryParam) ([]view.HbaDeviceInventoryView, int, error) {
	var fcHbaDevices []view.HbaDeviceInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/hba", params, &fcHbaDevices)
	return fcHbaDevices, total, err
}
