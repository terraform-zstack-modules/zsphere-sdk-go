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

// QueryVmInstancePciDeviceSpecRef queries VmInstancePciDeviceSpecRef list
func (cli *ZSClient) QueryVmInstancePciDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstancePciDeviceSpecRefInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/{vmInstanceUuid}/pci-device-specs", params, &resp)
}

func (cli *ZSClient) GetVmInstancePciDeviceSpecRef(ctx context.Context, vmInstanceUuid string, pciSpecUuid string) (*view.VmInstancePciDeviceSpecRefInventoryView, error) {
	var resp view.VmInstancePciDeviceSpecRefInventoryView
	err := cli.GetWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("pci-device-specs", pciSpecUuid), "inventories", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstancePciDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstancePciDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstancePciDeviceSpecRefInventoryView, int, error) {
	var vmInstancePciDeviceSpecRefs []view.VmInstancePciDeviceSpecRefInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/{vmInstanceUuid}/pci-device-specs", params, &vmInstancePciDeviceSpecRefs)
	return vmInstancePciDeviceSpecRefs, total, err
}
