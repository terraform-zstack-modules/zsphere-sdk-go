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

// UpdateSNSFeiShuEndpoint updates SNSFeiShuEndpoint
func (cli *ZSClient) UpdateSNSFeiShuEndpoint(ctx context.Context, uuid string, params param.UpdateSNSFeiShuEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/feishu", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSFeiShuEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSFeiShuEndpoint queries SNSFeiShuEndpoint list
func (cli *ZSClient) QuerySNSFeiShuEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSFeiShuEndpointInventoryView, error) {
	var resp []view.SNSFeiShuEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/feishu", params, &resp)
}

func (cli *ZSClient) GetSNSFeiShuEndpoint(ctx context.Context, uuid string) (*view.SNSFeiShuEndpointInventoryView, error) {
	var resp view.SNSFeiShuEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/feishu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSFeiShuEndpoint Pagination
func (cli *ZSClient) PageSNSFeiShuEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSFeiShuEndpointInventoryView, int, error) {
	var sNSFeiShuEndpoints []view.SNSFeiShuEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/feishu", params, &sNSFeiShuEndpoints)
	return sNSFeiShuEndpoints, total, err
}
// CreateSNSFeiShuEndpoint creates SNSFeiShuEndpoint
func (cli *ZSClient) CreateSNSFeiShuEndpoint(ctx context.Context, params param.CreateSNSFeiShuEndpointParam) (*view.SNSFeiShuEndpointInventoryView, error) {
	resp := view.SNSFeiShuEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/feishu"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
