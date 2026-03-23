// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryGpuDevice queries GpuDevice list
func (cli *ZSClient) QueryGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.GpuDeviceInventoryView, error) {
	var resp []view.GpuDeviceInventoryView
	return resp, cli.List(ctx, "v1/gpu-device/gpu-devices", params, &resp)
}

func (cli *ZSClient) GetGpuDevice(ctx context.Context, uuid string) (*view.GpuDeviceInventoryView, error) {
	var resp view.GpuDeviceInventoryView
	if err := cli.Get(ctx, "v1/gpu-device/gpu-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGpuDevice Pagination
func (cli *ZSClient) PageGpuDevice(ctx context.Context, params *param.QueryParam) ([]view.GpuDeviceInventoryView, int, error) {
	var gpuDevices []view.GpuDeviceInventoryView
	total, err := cli.Page(ctx, "v1/gpu-device/gpu-devices", params, &gpuDevices)
	return gpuDevices, total, err
}
