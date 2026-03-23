// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryActiveAlarmTemplate queries ActiveAlarmTemplate list
func (cli *ZSClient) QueryActiveAlarmTemplate(ctx context.Context, params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, error) {
	var resp []view.ActiveAlarmTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/activealarms/templates", params, &resp)
}

func (cli *ZSClient) GetActiveAlarmTemplate(ctx context.Context, uuid string) (*view.ActiveAlarmTemplateInventoryView, error) {
	var resp view.ActiveAlarmTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/activealarms/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageActiveAlarmTemplate Pagination
func (cli *ZSClient) PageActiveAlarmTemplate(ctx context.Context, params *param.QueryParam) ([]view.ActiveAlarmTemplateInventoryView, int, error) {
	var activeAlarmTemplates []view.ActiveAlarmTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/activealarms/templates", params, &activeAlarmTemplates)
	return activeAlarmTemplates, total, err
}
// UpdateActiveAlarmTemplate updates ActiveAlarmTemplate
func (cli *ZSClient) UpdateActiveAlarmTemplate(ctx context.Context, uuid string, params param.UpdateActiveAlarmTemplateParam) (*view.ActiveAlarmTemplateInventoryView, error) {
	resp := view.ActiveAlarmTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/activealarms/templates", uuid, "actions", "inventory", map[string]interface{}{
		"updateActiveAlarmTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
