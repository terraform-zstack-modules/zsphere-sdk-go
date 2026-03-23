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

// DeleteHostKernelInterface deletes HostKernelInterface
func (cli *ZSClient) DeleteHostKernelInterface(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/kernel-interfaces", uuid, string(deleteMode))
}
// CreateHostKernelInterface creates HostKernelInterface
func (cli *ZSClient) CreateHostKernelInterface(ctx context.Context, params param.CreateHostKernelInterfaceParam) (*view.HostKernelInterfaceInventoryView, error) {
	resp := view.HostKernelInterfaceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/kernel-interfaces"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostKernelInterface queries HostKernelInterface list
func (cli *ZSClient) QueryHostKernelInterface(ctx context.Context, params *param.QueryParam) ([]view.HostKernelInterfaceInventoryView, error) {
	var resp []view.HostKernelInterfaceInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/kernel-interfaces", params, &resp)
}

func (cli *ZSClient) GetHostKernelInterface(ctx context.Context, uuid string) (*view.HostKernelInterfaceInventoryView, error) {
	var resp view.HostKernelInterfaceInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/kernel-interfaces", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostKernelInterface Pagination
func (cli *ZSClient) PageHostKernelInterface(ctx context.Context, params *param.QueryParam) ([]view.HostKernelInterfaceInventoryView, int, error) {
	var hostKernelInterfaces []view.HostKernelInterfaceInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/kernel-interfaces", params, &hostKernelInterfaces)
	return hostKernelInterfaces, total, err
}
// UpdateHostKernelInterface updates HostKernelInterface
func (cli *ZSClient) UpdateHostKernelInterface(ctx context.Context, uuid string, params param.UpdateHostKernelInterfaceParam) (*view.HostKernelInterfaceInventoryView, error) {
	resp := view.HostKernelInterfaceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks/kernel-interfaces", uuid, "actions", "inventory", map[string]interface{}{
		"updateHostKernelInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
