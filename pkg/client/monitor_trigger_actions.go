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

// QueryMonitorTrigger queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTrigger(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerInventoryView, error) {
	var resp []view.MonitorTriggerInventoryView
	return resp, cli.List(ctx, "v1/monitoring/triggers", params, &resp)
}

func (cli *ZSClient) GetMonitorTrigger(ctx context.Context, uuid string) (*view.MonitorTriggerInventoryView, error) {
	var resp view.MonitorTriggerInventoryView
	if err := cli.Get(ctx, "v1/monitoring/triggers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorTrigger Pagination
func (cli *ZSClient) PageMonitorTrigger(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerInventoryView, int, error) {
	var monitorTriggers []view.MonitorTriggerInventoryView
	total, err := cli.Page(ctx, "v1/monitoring/triggers", params, &monitorTriggers)
	return monitorTriggers, total, err
}
// QueryMonitorTriggerAction queries MonitorTrigger list
func (cli *ZSClient) QueryMonitorTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, error) {
	var resp []view.MonitorTriggerActionInventoryView
	return resp, cli.List(ctx, "v1/monitoring/trigger-actions", params, &resp)
}

func (cli *ZSClient) GetMonitorTriggerAction(ctx context.Context, uuid string) (*view.MonitorTriggerActionInventoryView, error) {
	var resp view.MonitorTriggerActionInventoryView
	if err := cli.Get(ctx, "v1/monitoring/trigger-actions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorTriggerAction Pagination
func (cli *ZSClient) PageMonitorTriggerAction(ctx context.Context, params *param.QueryParam) ([]view.MonitorTriggerActionInventoryView, int, error) {
	var monitorTriggers []view.MonitorTriggerActionInventoryView
	total, err := cli.Page(ctx, "v1/monitoring/trigger-actions", params, &monitorTriggers)
	return monitorTriggers, total, err
}
// DeleteMonitorTrigger deletes MonitorTrigger
func (cli *ZSClient) DeleteMonitorTrigger(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/monitoring/triggers", uuid, string(deleteMode))
}
// UpdateMonitorTrigger updates MonitorTrigger
func (cli *ZSClient) UpdateMonitorTrigger(ctx context.Context, uuid string, params param.UpdateMonitorTriggerParam) (*view.MonitorTriggerInventoryView, error) {
	resp := view.MonitorTriggerInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/monitoring/triggers", uuid, "actions", "inventory", map[string]interface{}{
		"updateMonitorTrigger": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateMonitorTrigger creates MonitorTrigger
func (cli *ZSClient) CreateMonitorTrigger(ctx context.Context, params param.CreateMonitorTriggerParam) (*view.MonitorTriggerInventoryView, error) {
	resp := view.MonitorTriggerInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/monitoring/triggers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMonitorTriggerAction deletes MonitorTrigger
func (cli *ZSClient) DeleteMonitorTriggerAction(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/monitoring/trigger-actions", uuid, string(deleteMode))
}
