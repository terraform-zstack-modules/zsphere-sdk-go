// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryMonitorGroupInstance queries MonitorGroupInstance list
func (cli *ZSClient) QueryMonitorGroupInstance(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, error) {
	var resp []view.MonitorGroupInstanceInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups/instances", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupInstance(ctx context.Context, uuid string) (*view.MonitorGroupInstanceInventoryView, error) {
	var resp view.MonitorGroupInstanceInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitorgroups/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupInstance Pagination
func (cli *ZSClient) PageMonitorGroupInstance(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, int, error) {
	var monitorGroupInstances []view.MonitorGroupInstanceInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups/instances", params, &monitorGroupInstances)
	return monitorGroupInstances, total, err
}
