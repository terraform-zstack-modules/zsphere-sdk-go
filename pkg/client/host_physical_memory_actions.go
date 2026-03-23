// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryHostPhysicalMemory queries HostPhysicalMemory list
func (cli *ZSClient) QueryHostPhysicalMemory(ctx context.Context, params *param.QueryParam) ([]view.HostPhysicalMemoryInventoryView, error) {
	var resp []view.HostPhysicalMemoryInventoryView
	return resp, cli.List(ctx, "v1/hosts/physicalmemory", params, &resp)
}

func (cli *ZSClient) GetHostPhysicalMemory(ctx context.Context, uuid string) (*view.HostPhysicalMemoryInventoryView, error) {
	var resp view.HostPhysicalMemoryInventoryView
	if err := cli.Get(ctx, "v1/hosts/physicalmemory", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostPhysicalMemory Pagination
func (cli *ZSClient) PageHostPhysicalMemory(ctx context.Context, params *param.QueryParam) ([]view.HostPhysicalMemoryInventoryView, int, error) {
	var hostPhysicalMemories []view.HostPhysicalMemoryInventoryView
	total, err := cli.Page(ctx, "v1/hosts/physicalmemory", params, &hostPhysicalMemories)
	return hostPhysicalMemories, total, err
}
