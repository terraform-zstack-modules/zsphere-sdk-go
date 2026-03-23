// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateMdevDevice updates MdevDevice
func (cli *ZSClient) UpdateMdevDevice(ctx context.Context, uuid string, params param.UpdateMdevDeviceParam) (*view.MdevDeviceInventoryView, error) {
	resp := view.MdevDeviceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/mdev-devices", uuid, "actions", "inventory", map[string]interface{}{
		"updateMdevDevice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMdevDevice deletes MdevDevice
func (cli *ZSClient) DeleteMdevDevice(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/mdev-devices", uuid, string(deleteMode))
}
// QueryMdevDevice queries MdevDevice list
func (cli *ZSClient) QueryMdevDevice(ctx context.Context, params *param.QueryParam) ([]view.MdevDeviceInventoryView, error) {
	var resp []view.MdevDeviceInventoryView
	return resp, cli.List(ctx, "v1/mdev-devices", params, &resp)
}

func (cli *ZSClient) GetMdevDevice(ctx context.Context, uuid string) (*view.MdevDeviceInventoryView, error) {
	var resp view.MdevDeviceInventoryView
	if err := cli.Get(ctx, "v1/mdev-devices", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMdevDevice Pagination
func (cli *ZSClient) PageMdevDevice(ctx context.Context, params *param.QueryParam) ([]view.MdevDeviceInventoryView, int, error) {
	var mdevDevices []view.MdevDeviceInventoryView
	total, err := cli.Page(ctx, "v1/mdev-devices", params, &mdevDevices)
	return mdevDevices, total, err
}
