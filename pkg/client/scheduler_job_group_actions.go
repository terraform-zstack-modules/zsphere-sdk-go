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

// CreateSchedulerJobGroup creates SchedulerJobGroup
func (cli *ZSClient) CreateSchedulerJobGroup(ctx context.Context, params param.CreateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	resp := view.SchedulerJobGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/jobgroups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSchedulerJobGroup deletes SchedulerJobGroup
func (cli *ZSClient) DeleteSchedulerJobGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scheduler/jobgroups", uuid, string(deleteMode))
}
// UpdateSchedulerJobGroup updates SchedulerJobGroup
func (cli *ZSClient) UpdateSchedulerJobGroup(ctx context.Context, uuid string, params param.UpdateSchedulerJobGroupParam) (*view.SchedulerJobGroupInventoryView, error) {
	resp := view.SchedulerJobGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/scheduler/jobgroups", uuid, "actions", "inventory", map[string]interface{}{
		"updateSchedulerJobGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySchedulerJobGroup queries SchedulerJobGroup list
func (cli *ZSClient) QuerySchedulerJobGroup(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, error) {
	var resp []view.SchedulerJobGroupInventoryView
	return resp, cli.List(ctx, "v1/scheduler/jobgroups", params, &resp)
}

func (cli *ZSClient) GetSchedulerJobGroup(ctx context.Context, uuid string) (*view.SchedulerJobGroupInventoryView, error) {
	var resp view.SchedulerJobGroupInventoryView
	if err := cli.Get(ctx, "v1/scheduler/jobgroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerJobGroup Pagination
func (cli *ZSClient) PageSchedulerJobGroup(ctx context.Context, params *param.QueryParam) ([]view.SchedulerJobGroupInventoryView, int, error) {
	var schedulerJobGroups []view.SchedulerJobGroupInventoryView
	total, err := cli.Page(ctx, "v1/scheduler/jobgroups", params, &schedulerJobGroups)
	return schedulerJobGroups, total, err
}
