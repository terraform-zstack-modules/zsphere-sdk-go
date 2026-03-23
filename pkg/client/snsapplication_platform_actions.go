// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteSNSApplicationPlatform deletes SNSApplicationPlatform
func (cli *ZSClient) DeleteSNSApplicationPlatform(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/sns/application-platforms", uuid, string(deleteMode))
}
// UpdateSNSApplicationPlatform updates SNSApplicationPlatform
func (cli *ZSClient) UpdateSNSApplicationPlatform(ctx context.Context, uuid string, params param.UpdateSNSApplicationPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-platforms", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSApplicationPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSApplicationPlatform queries SNSApplicationPlatform list
func (cli *ZSClient) QuerySNSApplicationPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, error) {
	var resp []view.SNSApplicationPlatformInventoryView
	return resp, cli.List(ctx, "v1/sns/application-platforms", params, &resp)
}

func (cli *ZSClient) GetSNSApplicationPlatform(ctx context.Context, uuid string) (*view.SNSApplicationPlatformInventoryView, error) {
	var resp view.SNSApplicationPlatformInventoryView
	if err := cli.Get(ctx, "v1/sns/application-platforms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSApplicationPlatform Pagination
func (cli *ZSClient) PageSNSApplicationPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSApplicationPlatformInventoryView, int, error) {
	var sNSApplicationPlatforms []view.SNSApplicationPlatformInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-platforms", params, &sNSApplicationPlatforms)
	return sNSApplicationPlatforms, total, err
}
