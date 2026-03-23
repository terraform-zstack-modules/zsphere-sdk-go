// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVmPriorityConfig queries VmPriorityConfig list
func (cli *ZSClient) QueryVmPriorityConfig(ctx context.Context, params *param.QueryParam) ([]view.VmPriorityConfigInventoryView, error) {
	var resp []view.VmPriorityConfigInventoryView
	return resp, cli.List(ctx, "v1/vm-priority-config", params, &resp)
}

func (cli *ZSClient) GetVmPriorityConfig(ctx context.Context, uuid string) (*view.VmPriorityConfigInventoryView, error) {
	var resp view.VmPriorityConfigInventoryView
	if err := cli.Get(ctx, "v1/vm-priority-config", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmPriorityConfig Pagination
func (cli *ZSClient) PageVmPriorityConfig(ctx context.Context, params *param.QueryParam) ([]view.VmPriorityConfigInventoryView, int, error) {
	var vmPriorityConfigs []view.VmPriorityConfigInventoryView
	total, err := cli.Page(ctx, "v1/vm-priority-config", params, &vmPriorityConfigs)
	return vmPriorityConfigs, total, err
}
