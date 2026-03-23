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

// CreateSchedulerJob creates SchedulerJob
func (cli *ZSClient) CreateSchedulerJob(ctx context.Context, params param.CreateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobs"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSchedulerJob updates SchedulerJob
func (cli *ZSClient) UpdateSchedulerJob(ctx context.Context, uuid string, params param.UpdateSchedulerJobParam) (*view.SchedulerJobInventoryView, error) {
	resp := view.SchedulerJobInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/scheduler/jobs", uuid, "actions", "inventory", map[string]interface{}{
		"updateSchedulerJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSchedulerJob deletes SchedulerJob
func (cli *ZSClient) DeleteSchedulerJob(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scheduler/jobs", uuid, string(deleteMode))
}
// QuerySchedulerJob queries SchedulerJob list
func (cli *ZSClient) QuerySchedulerJob(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobInventoryView, error) {
	var resp []view.SchedulerJobInventoryView
	return resp, cli.List(ctx, "v1/scheduler/jobs", params, &resp)
}

func (cli *ZSClient) GetSchedulerJob(ctx context.Context, uuid string) (*view.SchedulerJobInventoryView, error) {
	var resp view.SchedulerJobInventoryView
	if err := cli.Get(ctx, "v1/scheduler/jobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJob Pagination
func (cli *ZSClient) PageSchedulerJob(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobInventoryView, int, error) {
	var schedulerJobs []view.SchedulerJobInventoryView
	total, err := cli.Page(ctx, "v1/scheduler/jobs", params, &schedulerJobs)
	return schedulerJobs, total, err
}
