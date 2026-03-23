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

// UpdateHostNetworkInterface updates HostNetworkInterface
func (cli *ZSClient) UpdateHostNetworkInterface(ctx context.Context, interfaceUuid string, params param.UpdateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/nics/%s/actions", interfaceUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryHostNetworkInterface queries HostNetworkInterface list
func (cli *ZSClient) QueryHostNetworkInterface(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, error) {
	var resp []view.HostNetworkInterfaceInventoryView
	return resp, cli.List(ctx, "v1/hosts/nics", params, &resp)
}

func (cli *ZSClient) GetHostNetworkInterface(ctx context.Context, uuid string) (*view.HostNetworkInterfaceInventoryView, error) {
	var resp view.HostNetworkInterfaceInventoryView
	if err := cli.Get(ctx, "v1/hosts/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostNetworkInterface Pagination
func (cli *ZSClient) PageHostNetworkInterface(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkInterfaceInventoryView, int, error) {
	var hostNetworkInterfaces []view.HostNetworkInterfaceInventoryView
	total, err := cli.Page(ctx, "v1/hosts/nics", params, &hostNetworkInterfaces)
	return hostNetworkInterfaces, total, err
}
// LocateHostNetworkInterface operates on HostNetworkInterface
func (cli *ZSClient) LocateHostNetworkInterface(ctx context.Context, hostUuid string, params param.LocateHostNetworkInterfaceParam) (*view.HostNetworkInterfaceInventoryView, error) {
	resp := view.HostNetworkInterfaceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts", hostUuid, "locate/network-interface", "", map[string]interface{}{
		"locateHostNetworkInterface": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
