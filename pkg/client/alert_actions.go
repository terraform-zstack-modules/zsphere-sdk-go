// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/monitoring/alerts", uuid, string(deleteMode))
}
// QueryAlert queries Alert list
func (cli *ZSClient) QueryAlert(ctx context.Context, params *param.QueryParam) ([]view.AlertInventoryView, error) {
	var resp []view.AlertInventoryView
	return resp, cli.List(ctx, "v1/monitoring/alerts", params, &resp)
}

func (cli *ZSClient) GetAlert(ctx context.Context, uuid string) (*view.AlertInventoryView, error) {
	var resp view.AlertInventoryView
	if err := cli.Get(ctx, "v1/monitoring/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAlert Pagination
func (cli *ZSClient) PageAlert(ctx context.Context, params *param.QueryParam) ([]view.AlertInventoryView, int, error) {
	var alerts []view.AlertInventoryView
	total, err := cli.Page(ctx, "v1/monitoring/alerts", params, &alerts)
	return alerts, total, err
}
