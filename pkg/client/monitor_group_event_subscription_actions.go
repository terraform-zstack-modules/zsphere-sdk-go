// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryMonitorGroupEventSubscription queries MonitorGroupEventSubscription list
func (cli *ZSClient) QueryMonitorGroupEventSubscription(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp []view.MonitorGroupEventSubscriptionInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitorgroups/subscriptions", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupEventSubscription(ctx context.Context, uuid string) (*view.MonitorGroupEventSubscriptionInventoryView, error) {
	var resp view.MonitorGroupEventSubscriptionInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitorgroups/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupEventSubscription Pagination
func (cli *ZSClient) PageMonitorGroupEventSubscription(ctx context.Context, params *param.QueryParam) ([]view.MonitorGroupEventSubscriptionInventoryView, int, error) {
	var monitorGroupEventSubscriptions []view.MonitorGroupEventSubscriptionInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitorgroups/subscriptions", params, &monitorGroupEventSubscriptions)
	return monitorGroupEventSubscriptions, total, err
}
