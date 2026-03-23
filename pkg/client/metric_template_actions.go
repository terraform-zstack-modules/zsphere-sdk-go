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

// DeleteMetricTemplate deletes MetricTemplate
func (cli *ZSClient) DeleteMetricTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/metrics/httpreceivers/templates", uuid, string(deleteMode))
}
// QueryMetricTemplate queries MetricTemplate list
func (cli *ZSClient) QueryMetricTemplate(ctx context.Context, params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List(ctx, "v1/zwatch/metrics/httpreceivers/templates", params, &resp)
}

func (cli *ZSClient) GetMetricTemplate(ctx context.Context, uuid string) (*view.MetricDataHttpReceiverInventoryView, error) {
	var resp view.MetricDataHttpReceiverInventoryView
	if err := cli.Get(ctx, "v1/zwatch/metrics/httpreceivers/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMetricTemplate Pagination
func (cli *ZSClient) PageMetricTemplate(ctx context.Context, params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, int, error) {
	var metricTemplates []view.MetricDataHttpReceiverInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/metrics/httpreceivers/templates", params, &metricTemplates)
	return metricTemplates, total, err
}
// CreateMetricTemplate creates MetricTemplate
func (cli *ZSClient) CreateMetricTemplate(ctx context.Context, params param.CreateMetricTemplateParam) (*view.MetricTemplateInventoryView, error) {
	resp := view.MetricTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/metrics/httpreceivers/templates"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
