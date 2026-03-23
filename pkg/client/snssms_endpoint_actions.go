// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySNSSmsEndpoint queries SNSSmsEndpoint list
func (cli *ZSClient) QuerySNSSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSSmsEndpointInventoryView, error) {
	var resp []view.SNSSmsEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/sms-endpoints", params, &resp)
}

func (cli *ZSClient) GetSNSSmsEndpoint(ctx context.Context, uuid string) (*view.SNSSmsEndpointInventoryView, error) {
	var resp view.SNSSmsEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/sms-endpoints", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSSmsEndpoint Pagination
func (cli *ZSClient) PageSNSSmsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSSmsEndpointInventoryView, int, error) {
	var sNSSmsEndpoints []view.SNSSmsEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/sms-endpoints", params, &sNSSmsEndpoints)
	return sNSSmsEndpoints, total, err
}
