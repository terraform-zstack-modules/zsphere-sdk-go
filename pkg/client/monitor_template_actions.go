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

// CloneMonitorTemplate operates on MonitorTemplate
func (cli *ZSClient) CloneMonitorTemplate(ctx context.Context, uuid string, params param.CloneMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitortemplates/%s/actions", uuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateMonitorTemplate updates MonitorTemplate
func (cli *ZSClient) UpdateMonitorTemplate(ctx context.Context, uuid string, params param.UpdateMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/monitortemplates", uuid, "actions", "inventory", map[string]interface{}{
		"updateMonitorTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMonitorTemplate queries MonitorTemplate list
func (cli *ZSClient) QueryMonitorTemplate(ctx context.Context, params *param.QueryParam) ([]view.MonitorTemplateInventoryView, error) {
	var resp []view.MonitorTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitortemplates", params, &resp)
}

func (cli *ZSClient) GetMonitorTemplate(ctx context.Context, uuid string) (*view.MonitorTemplateInventoryView, error) {
	var resp view.MonitorTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitortemplates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorTemplate Pagination
func (cli *ZSClient) PageMonitorTemplate(ctx context.Context, params *param.QueryParam) ([]view.MonitorTemplateInventoryView, int, error) {
	var monitorTemplates []view.MonitorTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitortemplates", params, &monitorTemplates)
	return monitorTemplates, total, err
}
// CreateMonitorTemplate creates MonitorTemplate
func (cli *ZSClient) CreateMonitorTemplate(ctx context.Context, params param.CreateMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitortemplates"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMonitorTemplate deletes MonitorTemplate
func (cli *ZSClient) DeleteMonitorTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/monitortemplates", uuid, string(deleteMode))
}
