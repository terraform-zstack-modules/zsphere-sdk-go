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

// AddEventRuleTemplate adds EventRuleTemplate
func (cli *ZSClient) AddEventRuleTemplate(ctx context.Context, params param.AddEventRuleTemplateParam) (*view.EventRuleTemplateInventoryView, error) {
	resp := view.EventRuleTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitortemplates/evenrules"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEventRuleTemplate deletes EventRuleTemplate
func (cli *ZSClient) DeleteEventRuleTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/monitortemplates/evenrules", uuid, string(deleteMode))
}
// UpdateEventRuleTemplate updates EventRuleTemplate
func (cli *ZSClient) UpdateEventRuleTemplate(ctx context.Context, uuid string, params param.UpdateEventRuleTemplateParam) (*view.EventRuleTemplateInventoryView, error) {
	resp := view.EventRuleTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/monitortemplates/evenrules", uuid, "actions", "inventory", map[string]interface{}{
		"updateEventRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryEventRuleTemplate queries EventRuleTemplate list
func (cli *ZSClient) QueryEventRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.EventRuleTemplateInventoryView, error) {
	var resp []view.EventRuleTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitortemplates/evenrules", params, &resp)
}

func (cli *ZSClient) GetEventRuleTemplate(ctx context.Context, uuid string) (*view.EventRuleTemplateInventoryView, error) {
	var resp view.EventRuleTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitortemplates/evenrules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventRuleTemplate Pagination
func (cli *ZSClient) PageEventRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.EventRuleTemplateInventoryView, int, error) {
	var eventRuleTemplates []view.EventRuleTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitortemplates/evenrules", params, &eventRuleTemplates)
	return eventRuleTemplates, total, err
}
