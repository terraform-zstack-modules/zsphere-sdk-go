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

// CreateWebhook creates Webhook
func (cli *ZSClient) CreateWebhook(ctx context.Context, params param.CreateWebhookParam) (*view.WebhookInventoryView, error) {
	resp := view.WebhookInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/web-hooks"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryWebhook queries Webhook list
func (cli *ZSClient) QueryWebhook(ctx context.Context, params *param.QueryParam) ([]view.WebhookInventoryView, error) {
	var resp []view.WebhookInventoryView
	return resp, cli.List(ctx, "v1/web-hooks", params, &resp)
}

func (cli *ZSClient) GetWebhook(ctx context.Context, uuid string) (*view.WebhookInventoryView, error) {
	var resp view.WebhookInventoryView
	if err := cli.Get(ctx, "v1/web-hooks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageWebhook Pagination
func (cli *ZSClient) PageWebhook(ctx context.Context, params *param.QueryParam) ([]view.WebhookInventoryView, int, error) {
	var webhooks []view.WebhookInventoryView
	total, err := cli.Page(ctx, "v1/web-hooks", params, &webhooks)
	return webhooks, total, err
}
// UpdateWebhook updates Webhook
func (cli *ZSClient) UpdateWebhook(ctx context.Context, uuid string, params param.UpdateWebhookParam) (*view.WebhookInventoryView, error) {
	resp := view.WebhookInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/web-hooks", uuid, "", "inventory", map[string]interface{}{
		"updateWebhook": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteWebhook deletes Webhook
func (cli *ZSClient) DeleteWebhook(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/web-hooks", uuid, string(deleteMode))
}
