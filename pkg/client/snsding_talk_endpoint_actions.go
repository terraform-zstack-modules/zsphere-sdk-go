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

// UpdateSNSDingTalkEndpoint updates SNSDingTalkEndpoint
func (cli *ZSClient) UpdateSNSDingTalkEndpoint(ctx context.Context, uuid string, params param.UpdateSNSDingTalkEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/ding-talk", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSDingTalkEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSDingTalkEndpoint creates SNSDingTalkEndpoint
func (cli *ZSClient) CreateSNSDingTalkEndpoint(ctx context.Context, params param.CreateSNSDingTalkEndpointParam) (*view.SNSDingTalkEndpointInventoryView, error) {
	resp := view.SNSDingTalkEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/ding-talk"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSDingTalkEndpoint queries SNSDingTalkEndpoint list
func (cli *ZSClient) QuerySNSDingTalkEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, error) {
	var resp []view.SNSDingTalkEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/ding-talk", params, &resp)
}

func (cli *ZSClient) GetSNSDingTalkEndpoint(ctx context.Context, uuid string) (*view.SNSDingTalkEndpointInventoryView, error) {
	var resp view.SNSDingTalkEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/ding-talk", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSDingTalkEndpoint Pagination
func (cli *ZSClient) PageSNSDingTalkEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, int, error) {
	var sNSDingTalkEndpoints []view.SNSDingTalkEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/ding-talk", params, &sNSDingTalkEndpoints)
	return sNSDingTalkEndpoints, total, err
}
