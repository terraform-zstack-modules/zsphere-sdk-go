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

// UpdateAlarm updates Alarm
func (cli *ZSClient) UpdateAlarm(ctx context.Context, uuid string, params param.UpdateAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/alarms", uuid, "actions", "inventory", map[string]interface{}{
		"updateAlarm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAlarm deletes Alarm
func (cli *ZSClient) DeleteAlarm(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/alarms", uuid, string(deleteMode))
}
// QueryAlarm queries Alarm list
func (cli *ZSClient) QueryAlarm(ctx context.Context, params *param.QueryParam) ([]view.AlarmInventoryView, error) {
	var resp []view.AlarmInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alarms", params, &resp)
}

func (cli *ZSClient) GetAlarm(ctx context.Context, uuid string) (*view.AlarmInventoryView, error) {
	var resp view.AlarmInventoryView
	if err := cli.Get(ctx, "v1/zwatch/alarms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlarm Pagination
func (cli *ZSClient) PageAlarm(ctx context.Context, params *param.QueryParam) ([]view.AlarmInventoryView, int, error) {
	var alarms []view.AlarmInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alarms", params, &alarms)
	return alarms, total, err
}
// CreateAlarm creates Alarm
func (cli *ZSClient) CreateAlarm(ctx context.Context, params param.CreateAlarmParam) (*view.AlarmInventoryView, error) {
	resp := view.AlarmInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/alarms"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
