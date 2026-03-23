// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteSNSApplicationEndpoint deletes SNSApplicationEndpoint
func (cli *ZSClient) DeleteSNSApplicationEndpoint(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sns/application-endpoints", uuid, string(deleteMode))
}
// QuerySNSApplicationEndpoint queries SNSApplicationEndpoint list
func (cli *ZSClient) QuerySNSApplicationEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, error) {
	var resp []view.SNSApplicationEndpointInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints", params, &resp)
}

func (cli *ZSClient) GetSNSApplicationEndpoint(ctx context.Context, uuid string) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.SNSApplicationEndpointInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSApplicationEndpoint Pagination
func (cli *ZSClient) PageSNSApplicationEndpoint(ctx context.Context, params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, int, error) {
	var sNSApplicationEndpoints []view.SNSApplicationEndpointInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints", params, &sNSApplicationEndpoints)
	return sNSApplicationEndpoints, total, err
}
// UpdateSNSApplicationEndpoint updates SNSApplicationEndpoint
func (cli *ZSClient) UpdateSNSApplicationEndpoint(ctx context.Context, uuid string, params param.UpdateSNSApplicationEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-endpoints", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSApplicationEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
