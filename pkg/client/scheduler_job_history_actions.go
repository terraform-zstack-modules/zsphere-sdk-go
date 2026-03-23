// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySchedulerJobHistory queries SchedulerJobHistory list
func (cli *ZSClient) QuerySchedulerJobHistory(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, error) {
	var resp []view.SchedulerJobHistoryInventoryView
	return resp, cli.List(ctx, "v1/scheduler/job/history", params, &resp)
}

// PageSchedulerJobHistory Pagination
func (cli *ZSClient) PageSchedulerJobHistory(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobHistoryInventoryView, int, error) {
	var schedulerJobHistories []view.SchedulerJobHistoryInventoryView
	total, err := cli.Page(ctx, "v1/scheduler/job/history", params, &schedulerJobHistories)
	return schedulerJobHistories, total, err
}
