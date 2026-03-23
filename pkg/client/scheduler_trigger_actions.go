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

// DeleteSchedulerTrigger deletes SchedulerTrigger
func (cli *ZSClient) DeleteSchedulerTrigger(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/scheduler/triggers", uuid, string(deleteMode))
}
// CreateSchedulerTrigger creates SchedulerTrigger
func (cli *ZSClient) CreateSchedulerTrigger(ctx context.Context, params param.CreateSchedulerTriggerParam) (*view.SchedulerTriggerInventoryView, error) {
	resp := view.SchedulerTriggerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/scheduler/triggers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySchedulerTrigger queries SchedulerTrigger list
func (cli *ZSClient) QuerySchedulerTrigger(ctx context.Context, params *param.QueryParam) ([]view.SchedulerTriggerInventoryView, error) {
	var resp []view.SchedulerTriggerInventoryView
	return resp, cli.List(ctx, "v1/scheduler/triggers", params, &resp)
}

func (cli *ZSClient) GetSchedulerTrigger(ctx context.Context, uuid string) (*view.SchedulerTriggerInventoryView, error) {
	var resp view.SchedulerTriggerInventoryView
	if err := cli.Get(ctx, "v1/scheduler/triggers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSchedulerTrigger Pagination
func (cli *ZSClient) PageSchedulerTrigger(ctx context.Context, params *param.QueryParam) ([]view.SchedulerTriggerInventoryView, int, error) {
	var schedulerTriggers []view.SchedulerTriggerInventoryView
	total, err := cli.Page(ctx, "v1/scheduler/triggers", params, &schedulerTriggers)
	return schedulerTriggers, total, err
}
// UpdateSchedulerTrigger updates SchedulerTrigger
func (cli *ZSClient) UpdateSchedulerTrigger(ctx context.Context, uuid string, params param.UpdateSchedulerTriggerParam) (*view.SchedulerTriggerInventoryView, error) {
	resp := view.SchedulerTriggerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/scheduler/triggers", uuid, "actions", "inventory", map[string]interface{}{
		"updateSchedulerTrigger": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
