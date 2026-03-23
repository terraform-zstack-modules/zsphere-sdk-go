// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryEventSubscription queries EventSubscription list
func (cli *ZSClient) QueryEventSubscription(ctx context.Context, params *param.QueryParam) ([]view.EventSubscriptionInventoryView, error) {
	var resp []view.EventSubscriptionInventoryView
	return resp, cli.List(ctx, "v1/zwatch/events/subscriptions", params, &resp)
}

func (cli *ZSClient) GetEventSubscription(ctx context.Context, uuid string) (*view.EventSubscriptionInventoryView, error) {
	var resp view.EventSubscriptionInventoryView
	if err := cli.Get(ctx, "v1/zwatch/events/subscriptions", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventSubscription Pagination
func (cli *ZSClient) PageEventSubscription(ctx context.Context, params *param.QueryParam) ([]view.EventSubscriptionInventoryView, int, error) {
	var eventSubscriptions []view.EventSubscriptionInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/events/subscriptions", params, &eventSubscriptions)
	return eventSubscriptions, total, err
}
