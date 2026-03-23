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

// ValidateSNSEmailPlatform operates on SNSEmailPlatform
func (cli *ZSClient) ValidateSNSEmailPlatform(ctx context.Context, uuid string, params param.ValidateSNSEmailPlatformParam) (*view.SNSEmailPlatformInventoryView, error) {
	resp := view.SNSEmailPlatformInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/sns/application-platforms/email", uuid, "actions", "", map[string]interface{}{
		"validateSNSEmailPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSEmailPlatform creates SNSEmailPlatform
func (cli *ZSClient) CreateSNSEmailPlatform(ctx context.Context, params param.CreateSNSEmailPlatformParam) (*view.SNSApplicationPlatformInventoryView, error) {
	resp := view.SNSApplicationPlatformInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-platforms/email"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSEmailPlatform queries SNSEmailPlatform list
func (cli *ZSClient) QuerySNSEmailPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, error) {
	var resp []view.SNSEmailPlatformInventoryView
	return resp, cli.List(ctx, "v1/sns/application-platforms/email", params, &resp)
}

func (cli *ZSClient) GetSNSEmailPlatform(ctx context.Context, uuid string) (*view.SNSEmailPlatformInventoryView, error) {
	var resp view.SNSEmailPlatformInventoryView
	if err := cli.Get(ctx, "v1/sns/application-platforms/email", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSEmailPlatform Pagination
func (cli *ZSClient) PageSNSEmailPlatform(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailPlatformInventoryView, int, error) {
	var sNSEmailPlatforms []view.SNSEmailPlatformInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-platforms/email", params, &sNSEmailPlatforms)
	return sNSEmailPlatforms, total, err
}
