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

// QueryNvmeTarget queries NvmeTarget list
func (cli *ZSClient) QueryNvmeTarget(ctx context.Context, params *param.QueryParam) ([]view.NvmeTargetInventoryView, error) {
	var resp []view.NvmeTargetInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/nvme/controllers", params, &resp)
}

func (cli *ZSClient) GetNvmeTarget(ctx context.Context, uuid string) (*view.NvmeTargetInventoryView, error) {
	var resp view.NvmeTargetInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/nvme", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNvmeTarget Pagination
func (cli *ZSClient) PageNvmeTarget(ctx context.Context, params *param.QueryParam) ([]view.NvmeTargetInventoryView, int, error) {
	var nvmeTargets []view.NvmeTargetInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/nvme/controllers", params, &nvmeTargets)
	return nvmeTargets, total, err
}
// RefreshNvmeTarget operates on NvmeTarget
func (cli *ZSClient) RefreshNvmeTarget(ctx context.Context, params param.RefreshNvmeTargetParam) (*view.NvmeTargetInventoryView, error) {
	resp := view.NvmeTargetInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/storage-devices/nvme/controllers"), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
