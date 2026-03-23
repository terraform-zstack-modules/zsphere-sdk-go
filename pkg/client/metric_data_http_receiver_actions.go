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

// DeleteMetricDataHttpReceiver deletes MetricDataHttpReceiver
func (cli *ZSClient) DeleteMetricDataHttpReceiver(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/metrics/httpreceivers", uuid, string(deleteMode))
}
// CreateMetricDataHttpReceiver creates MetricDataHttpReceiver
func (cli *ZSClient) CreateMetricDataHttpReceiver(ctx context.Context, params param.CreateMetricDataHttpReceiverParam) (*view.MetricDataHttpReceiverInventoryView, error) {
	resp := view.MetricDataHttpReceiverInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/metrics/httpreceivers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMetricDataHttpReceiver queries MetricDataHttpReceiver list
func (cli *ZSClient) QueryMetricDataHttpReceiver(ctx context.Context, params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, error) {
	var resp []view.MetricDataHttpReceiverInventoryView
	return resp, cli.List(ctx, "v1/zwatch/metrics/httpreceivers", params, &resp)
}

func (cli *ZSClient) GetMetricDataHttpReceiver(ctx context.Context, uuid string) (*view.MetricDataHttpReceiverInventoryView, error) {
	var resp view.MetricDataHttpReceiverInventoryView
	if err := cli.Get(ctx, "v1/zwatch/metrics/httpreceivers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMetricDataHttpReceiver Pagination
func (cli *ZSClient) PageMetricDataHttpReceiver(ctx context.Context, params *param.QueryParam) ([]view.MetricDataHttpReceiverInventoryView, int, error) {
	var metricDataHttpReceivers []view.MetricDataHttpReceiverInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/metrics/httpreceivers", params, &metricDataHttpReceivers)
	return metricDataHttpReceivers, total, err
}
