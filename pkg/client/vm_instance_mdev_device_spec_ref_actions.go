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

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List(ctx, "v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetVmInstanceMdevDeviceSpecRef(ctx context.Context, vmInstanceUuid string, mdevSpecUuid string) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp view.VmInstanceMdevDeviceSpecRefInventoryView
	err := cli.GetWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, fmt.Sprintf("mdev-device-specs", mdevSpecUuid), "inventories", nil, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstanceMdevDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstanceMdevDeviceSpecRef(ctx context.Context, params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, int, error) {
	var vmInstanceMdevDeviceSpecRefs []view.VmInstanceMdevDeviceSpecRefInventoryView
	total, err := cli.Page(ctx, "v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &vmInstanceMdevDeviceSpecRefs)
	return vmInstanceMdevDeviceSpecRefs, total, err
}
