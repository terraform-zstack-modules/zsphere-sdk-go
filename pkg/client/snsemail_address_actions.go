// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySNSEmailAddress queries SNSEmailAddress list
func (cli *ZSClient) QuerySNSEmailAddress(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailAddressInventoryView, error) {
	var resp []view.SNSEmailAddressInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/emails/email-addresses", params, &resp)
}

func (cli *ZSClient) GetSNSEmailAddress(ctx context.Context, uuid string) (*view.SNSEmailAddressInventoryView, error) {
	var resp view.SNSEmailAddressInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/emails/email-addresses", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSEmailAddress Pagination
func (cli *ZSClient) PageSNSEmailAddress(ctx context.Context, params *param.QueryParam) ([]view.SNSEmailAddressInventoryView, int, error) {
	var sNSEmailAddress []view.SNSEmailAddressInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/emails/email-addresses", params, &sNSEmailAddress)
	return sNSEmailAddress, total, err
}
