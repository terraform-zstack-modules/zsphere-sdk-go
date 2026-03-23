// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryActiveAlarm queries ActiveAlarm list
func (cli *ZSClient) QueryActiveAlarm(ctx context.Context, params *param.QueryParam) ([]view.ActiveAlarmInventoryView, error) {
	var resp []view.ActiveAlarmInventoryView
	return resp, cli.List(ctx, "v1/zwatch/activealarms/alarms", params, &resp)
}

func (cli *ZSClient) GetActiveAlarm(ctx context.Context, uuid string) (*view.ActiveAlarmInventoryView, error) {
	var resp view.ActiveAlarmInventoryView
	if err := cli.Get(ctx, "v1/zwatch/activealarms/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageActiveAlarm Pagination
func (cli *ZSClient) PageActiveAlarm(ctx context.Context, params *param.QueryParam) ([]view.ActiveAlarmInventoryView, int, error) {
	var activeAlarms []view.ActiveAlarmInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/activealarms/alarms", params, &activeAlarms)
	return activeAlarms, total, err
}
