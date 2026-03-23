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

// QuerySNSWeComEndpoint queries SNSWeComEndpoint list
func (cli *ZSClient) QuerySNSWeComEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSWeComEndpointInventoryView, error) {
	var resp []view.SNSWeComEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/we-com", params, &resp)
}

func (cli *ZSClient) GetSNSWeComEndpoint(ctx context.Context, uuid string) (*view.SNSWeComEndpointInventoryView, error) {
	var resp view.SNSWeComEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/we-com", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSWeComEndpoint Pagination
func (cli *ZSClient) PageSNSWeComEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSWeComEndpointInventoryView, int, error) {
	var sNSWeComEndpoints []view.SNSWeComEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/we-com", params, &sNSWeComEndpoints)
	return sNSWeComEndpoints, total, err
}
// UpdateSNSWeComEndpoint updates SNSWeComEndpoint
func (cli *ZSClient) UpdateSNSWeComEndpoint(ctx context.Context, uuid string, params param.UpdateSNSWeComEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/we-com", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSWeComEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSWeComEndpoint creates SNSWeComEndpoint
func (cli *ZSClient) CreateSNSWeComEndpoint(ctx context.Context, params param.CreateSNSWeComEndpointParam) (*view.SNSWeComEndpointInventoryView, error) {
	resp := view.SNSWeComEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/we-com"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
