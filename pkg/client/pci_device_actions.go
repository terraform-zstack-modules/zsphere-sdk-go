// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeletePciDevice deletes PciDevice
func (cli *ZSClient) DeletePciDevice(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/pci-device/pci-devices", uuid, string(deleteMode))
}
// UpdatePciDevice updates PciDevice
func (cli *ZSClient) UpdatePciDevice(ctx context.Context, uuid string, params param.UpdatePciDeviceParam) (*view.PciDeviceInventoryView, error) {
	resp := view.PciDeviceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-device/pci-devices", uuid, "actions", "inventory", map[string]interface{}{
		"updatePciDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPciDevice queries PciDevice list
func (cli *ZSClient) QueryPciDevice(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceInventoryView, error) {
	var resp []view.PciDeviceInventoryView
	return resp, cli.List(ctx, "v1/pci-device/pci-devices", params, &resp)
}

func (cli *ZSClient) GetPciDevice(ctx context.Context, uuid string) (*view.PciDeviceInventoryView, error) {
	var resp view.PciDeviceInventoryView
	if err := cli.Get(ctx, "v1/pci-device/pci-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDevice Pagination
func (cli *ZSClient) PagePciDevice(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceInventoryView, int, error) {
	var pciDevices []view.PciDeviceInventoryView
	total, err := cli.Page(ctx, "v1/pci-device/pci-devices", params, &pciDevices)
	return pciDevices, total, err
}
