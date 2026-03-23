// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySNSEndpointThirdpartyAlertHistory queries SNSEndpointThirdpartyAlertHistory list
func (cli *ZSClient) QuerySNSEndpointThirdpartyAlertHistory(ctx context.Context, params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, error) {
	var resp []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	return resp, cli.List(ctx, "v1/zwatch/third-party/alert-publish-histories", params, &resp)
}

// PageSNSEndpointThirdpartyAlertHistory Pagination
func (cli *ZSClient) PageSNSEndpointThirdpartyAlertHistory(ctx context.Context, params *param.QueryParam) ([]view.SNSEndpointThirdpartyAlertHistoryInventoryView, int, error) {
	var sNSEndpointThirdpartyAlertHistories []view.SNSEndpointThirdpartyAlertHistoryInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/third-party/alert-publish-histories", params, &sNSEndpointThirdpartyAlertHistories)
	return sNSEndpointThirdpartyAlertHistories, total, err
}
