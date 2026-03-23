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

// CreateSNSMicrosoftTeamsEndpoint creates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) CreateSNSMicrosoftTeamsEndpoint(ctx context.Context, params param.CreateSNSMicrosoftTeamsEndpointParam) (*view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	resp := view.SNSMicrosoftTeamsEndpointInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/microsoft-teams"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSNSMicrosoftTeamsEndpoint updates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) UpdateSNSMicrosoftTeamsEndpoint(ctx context.Context, uuid string, params param.UpdateSNSMicrosoftTeamsEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints/microsoft-teams", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSMicrosoftTeamsEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSMicrosoftTeamsEndpoint queries SNSMicrosoftTeamsEndpoint list
func (cli *ZSClient) QuerySNSMicrosoftTeamsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp []view.SNSMicrosoftTeamsEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/microsoft-teams", params, &resp)
}

func (cli *ZSClient) GetSNSMicrosoftTeamsEndpoint(ctx context.Context, uuid string) (*view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp view.SNSMicrosoftTeamsEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/microsoft-teams", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSMicrosoftTeamsEndpoint Pagination
func (cli *ZSClient) PageSNSMicrosoftTeamsEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSMicrosoftTeamsEndpointInventoryView, int, error) {
	var sNSMicrosoftTeamsEndpoints []view.SNSMicrosoftTeamsEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/microsoft-teams", params, &sNSMicrosoftTeamsEndpoints)
	return sNSMicrosoftTeamsEndpoints, total, err
}
