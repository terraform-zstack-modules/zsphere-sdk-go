// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVmSchedHistory queries VmSchedHistory list
func (cli *ZSClient) QueryVmSchedHistory(ctx context.Context, params *param.QueryParam) ([]view.VmSchedHistoryInventoryView, error) {
	var resp []view.VmSchedHistoryInventoryView
	return resp, cli.List(ctx, "v1/vm/sched-history", params, &resp)
}

func (cli *ZSClient) GetVmSchedHistory(ctx context.Context, uuid string) (*view.VmSchedHistoryInventoryView, error) {
	var resp view.VmSchedHistoryInventoryView
	if err := cli.Get(ctx, "v1/vm/sched-history", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmSchedHistory Pagination
func (cli *ZSClient) PageVmSchedHistory(ctx context.Context, params *param.QueryParam) ([]view.VmSchedHistoryInventoryView, int, error) {
	var vmSchedHistories []view.VmSchedHistoryInventoryView
	total, err := cli.Page(ctx, "v1/vm/sched-history", params, &vmSchedHistories)
	return vmSchedHistories, total, err
}
