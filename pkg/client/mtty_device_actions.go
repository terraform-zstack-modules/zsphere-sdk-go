// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryMttyDevice queries MttyDevice list
func (cli *ZSClient) QueryMttyDevice(ctx context.Context, params *param.QueryParam) ([]view.MttyDeviceInventoryView, error) {
	var resp []view.MttyDeviceInventoryView
	return resp, cli.List(ctx, "v1/mtty-devices", params, &resp)
}

func (cli *ZSClient) GetMttyDevice(ctx context.Context, uuid string) (*view.MttyDeviceInventoryView, error) {
	var resp view.MttyDeviceInventoryView
	if err := cli.Get(ctx, "v1/mtty-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMttyDevice Pagination
func (cli *ZSClient) PageMttyDevice(ctx context.Context, params *param.QueryParam) ([]view.MttyDeviceInventoryView, int, error) {
	var mttyDevices []view.MttyDeviceInventoryView
	total, err := cli.Page(ctx, "v1/mtty-devices", params, &mttyDevices)
	return mttyDevices, total, err
}
