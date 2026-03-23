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

// DeletePciDeviceOffering deletes PciDeviceOffering
func (cli *ZSClient) DeletePciDeviceOffering(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/pci-device/pci-device-offerings", uuid, string(deleteMode))
}
// CreatePciDeviceOffering creates PciDeviceOffering
func (cli *ZSClient) CreatePciDeviceOffering(ctx context.Context, params param.CreatePciDeviceOfferingParam) (*view.PciDeviceOfferingInventoryView, error) {
	resp := view.PciDeviceOfferingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/pci-device/pci-device-offerings"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPciDeviceOffering queries PciDeviceOffering list
func (cli *ZSClient) QueryPciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, error) {
	var resp []view.PciDeviceOfferingInventoryView
	return resp, cli.List(ctx, "v1/pci-device/pci-device-offerings", params, &resp)
}

func (cli *ZSClient) GetPciDeviceOffering(ctx context.Context, uuid string) (*view.PciDeviceOfferingInventoryView, error) {
	var resp view.PciDeviceOfferingInventoryView
	if err := cli.Get(ctx, "v1/pci-device/pci-device-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDeviceOffering Pagination
func (cli *ZSClient) PagePciDeviceOffering(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceOfferingInventoryView, int, error) {
	var pciDeviceOfferings []view.PciDeviceOfferingInventoryView
	total, err := cli.Page(ctx, "v1/pci-device/pci-device-offerings", params, &pciDeviceOfferings)
	return pciDeviceOfferings, total, err
}
