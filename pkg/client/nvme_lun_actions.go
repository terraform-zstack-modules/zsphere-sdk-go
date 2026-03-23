// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryNvmeLun queries NvmeLun list
func (cli *ZSClient) QueryNvmeLun(ctx context.Context, params *param.QueryParam) ([]view.NvmeLunInventoryView, error) {
	var resp []view.NvmeLunInventoryView
	return resp, cli.List(ctx, "v1/storage-devices/nvme/luns", params, &resp)
}

func (cli *ZSClient) GetNvmeLun(ctx context.Context, uuid string) (*view.NvmeLunInventoryView, error) {
	var resp view.NvmeLunInventoryView
	if err := cli.Get(ctx, "v1/storage-devices/nvme/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNvmeLun Pagination
func (cli *ZSClient) PageNvmeLun(ctx context.Context, params *param.QueryParam) ([]view.NvmeLunInventoryView, int, error) {
	var nvmeLuns []view.NvmeLunInventoryView
	total, err := cli.Page(ctx, "v1/storage-devices/nvme/luns", params, &nvmeLuns)
	return nvmeLuns, total, err
}
