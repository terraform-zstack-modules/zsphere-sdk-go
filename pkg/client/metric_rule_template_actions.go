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

// QueryMetricRuleTemplate queries MetricRuleTemplate list
func (cli *ZSClient) QueryMetricRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, error) {
	var resp []view.MetricRuleTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/monitortemplates/metricrules", params, &resp)
}

func (cli *ZSClient) GetMetricRuleTemplate(ctx context.Context, uuid string) (*view.MetricRuleTemplateInventoryView, error) {
	var resp view.MetricRuleTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/monitortemplates/metricrules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMetricRuleTemplate Pagination
func (cli *ZSClient) PageMetricRuleTemplate(ctx context.Context, params *param.QueryParam) ([]view.MetricRuleTemplateInventoryView, int, error) {
	var metricRuleTemplates []view.MetricRuleTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/monitortemplates/metricrules", params, &metricRuleTemplates)
	return metricRuleTemplates, total, err
}
// DeleteMetricRuleTemplate deletes MetricRuleTemplate
func (cli *ZSClient) DeleteMetricRuleTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/monitortemplates/metricrules", uuid, string(deleteMode))
}
// UpdateMetricRuleTemplate updates MetricRuleTemplate
func (cli *ZSClient) UpdateMetricRuleTemplate(ctx context.Context, uuid string, params param.UpdateMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	resp := view.MetricRuleTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/monitortemplates/metricrules", uuid, "actions", "inventory", map[string]interface{}{
		"updateMetricRuleTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddMetricRuleTemplate adds MetricRuleTemplate
func (cli *ZSClient) AddMetricRuleTemplate(ctx context.Context, monitorTemplateUuid string, params param.AddMetricRuleTemplateParam) (*view.MetricRuleTemplateInventoryView, error) {
	resp := view.MetricRuleTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/monitortemplates/%s/metricrules", monitorTemplateUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
