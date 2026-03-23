// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryPciDeviceSpec queries PciDeviceSpec list
func (cli *ZSClient) QueryPciDeviceSpec(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, error) {
	var resp []view.PciDeviceSpecInventoryView
	return resp, cli.List(ctx, "v1/pci-device-specs", params, &resp)
}

func (cli *ZSClient) GetPciDeviceSpec(ctx context.Context, uuid string) (*view.PciDeviceSpecInventoryView, error) {
	var resp view.PciDeviceSpecInventoryView
	if err := cli.Get(ctx, "v1/pci-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePciDeviceSpec Pagination
func (cli *ZSClient) PagePciDeviceSpec(ctx context.Context, params *param.QueryParam) ([]view.PciDeviceSpecInventoryView, int, error) {
	var pciDeviceSpecs []view.PciDeviceSpecInventoryView
	total, err := cli.Page(ctx, "v1/pci-device-specs", params, &pciDeviceSpecs)
	return pciDeviceSpecs, total, err
}
// UpdatePciDeviceSpec updates PciDeviceSpec
func (cli *ZSClient) UpdatePciDeviceSpec(ctx context.Context, uuid string, params param.UpdatePciDeviceSpecParam) (*view.PciDeviceSpecInventoryView, error) {
	resp := view.PciDeviceSpecInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/pci-device-specs", uuid, "actions", "inventory", map[string]interface{}{
		"updatePciDeviceSpec": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
