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

// CreateMonitorGroup creates MonitorGroup
func (cli *ZSClient) CreateMonitorGroup(ctx context.Context, params param.CreateMonitorGroupParam) (*view.MonitorGroupInventoryView, error) {
	resp := view.MonitorGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitorgroups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMonitorGroup deletes MonitorGroup
func (cli *ZSClient) DeleteMonitorGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/monitorgroups", uuid, string(deleteMode))
}
// QueryMonitorGroup queries MonitorGroup list
func (cli *ZSClient) QueryMonitorGroup(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupInventoryView, error) {
	var resp []view.MonitorGroupInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups", params, &resp)
}

func (cli *ZSClient) GetMonitorGroup(ctx context.Context, uuid string) (*view.MonitorGroupInventoryView, error) {
	var resp view.MonitorGroupInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitorgroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroup Pagination
func (cli *ZSClient) PageMonitorGroup(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupInventoryView, int, error) {
	var monitorGroups []view.MonitorGroupInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups", params, &monitorGroups)
	return monitorGroups, total, err
}
// UpdateMonitorGroup updates MonitorGroup
func (cli *ZSClient) UpdateMonitorGroup(ctx context.Context, uuid string, params param.UpdateMonitorGroupParam) (*view.MonitorGroupInventoryView, error) {
	resp := view.MonitorGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/monitorgroups", uuid, "actions", "inventory", map[string]interface{}{
		"updateMonitorGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
