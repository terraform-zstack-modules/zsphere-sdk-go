// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateEventSubscriptionLabel updates EventSubscriptionLabel
func (cli *ZSClient) UpdateEventSubscriptionLabel(ctx context.Context, uuid string, params param.UpdateEventSubscriptionLabelParam) (*view.EventSubscriptionLabelInventoryView, error) {
	resp := view.EventSubscriptionLabelInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/events/subscriptions/labels", uuid, "actions", "inventory", map[string]interface{}{
		"updateEventSubscriptionLabel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
