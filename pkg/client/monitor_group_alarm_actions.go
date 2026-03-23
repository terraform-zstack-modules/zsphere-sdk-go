// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryMonitorGroupAlarm queries MonitorGroupAlarm list
func (cli *ZSClient) QueryMonitorGroupAlarm(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, error) {
	var resp []view.MonitorGroupAlarmInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups/alarms", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupAlarm(ctx context.Context, uuid string) (*view.MonitorGroupAlarmInventoryView, error) {
	var resp view.MonitorGroupAlarmInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitorgroups/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupAlarm Pagination
func (cli *ZSClient) PageMonitorGroupAlarm(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupAlarmInventoryView, int, error) {
	var monitorGroupAlarms []view.MonitorGroupAlarmInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups/alarms", params, &monitorGroupAlarms)
	return monitorGroupAlarms, total, err
}
