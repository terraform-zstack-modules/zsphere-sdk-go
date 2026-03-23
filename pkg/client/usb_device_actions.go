// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateUsbDevice updates UsbDevice
func (cli *ZSClient) UpdateUsbDevice(ctx context.Context, uuid string, params param.UpdateUsbDeviceParam) (*view.UsbDeviceInventoryView, error) {
	resp := view.UsbDeviceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/usb-device/usb-devices", uuid, "actions", "inventory", map[string]interface{}{
		"updateUsbDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryUsbDevice queries UsbDevice list
func (cli *ZSClient) QueryUsbDevice(ctx context.Context, params *param.QueryParam) ([]view.UsbDeviceInventoryView, error) {
	var resp []view.UsbDeviceInventoryView
	return resp, cli.List(ctx, "v1/usb-device/usb-devices", params, &resp)
}

func (cli *ZSClient) GetUsbDevice(ctx context.Context, uuid string) (*view.UsbDeviceInventoryView, error) {
	var resp view.UsbDeviceInventoryView
	if err := cli.Get(ctx, "v1/usb-device/usb-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUsbDevice Pagination
func (cli *ZSClient) PageUsbDevice(ctx context.Context, params *param.QueryParam) ([]view.UsbDeviceInventoryView, int, error) {
	var usbDevices []view.UsbDeviceInventoryView
	total, err := cli.Page(ctx, "v1/usb-device/usb-devices", params, &usbDevices)
	return usbDevices, total, err
}
