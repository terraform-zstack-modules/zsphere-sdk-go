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

// CreateSNSEmailEndpoint creates SNSEmailEndpoint
func (cli *ZSClient) CreateSNSEmailEndpoint(ctx context.Context, params param.CreateSNSEmailEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/emails"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSEmailEndpoint queries SNSEmailEndpoint list
func (cli *ZSClient) QuerySNSEmailEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, error) {
	var resp []view.SNSEmailEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/emails", params, &resp)
}

func (cli *ZSClient) GetSNSEmailEndpoint(ctx context.Context, uuid string) (*view.SNSEmailEndpointInventoryView, error) {
	var resp view.SNSEmailEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSEmailEndpoint Pagination
func (cli *ZSClient) PageSNSEmailEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, int, error) {
	var sNSEmailEndpoints []view.SNSEmailEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/emails", params, &sNSEmailEndpoints)
	return sNSEmailEndpoints, total, err
}
