// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryAccountResourceRef queries AccountResourceRef list
func (cli *ZSClient) QueryAccountResourceRef(ctx context.Context, params *param.QueryParam) ([]view.AccountResourceRefInventoryView, error) {
	var resp []view.AccountResourceRefInventoryView
	return resp, cli.List(ctx, "v1/accounts/resources/refs", params, &resp)
}

// PageAccountResourceRef Pagination
func (cli *ZSClient) PageAccountResourceRef(ctx context.Context, params *param.QueryParam) ([]view.AccountResourceRefInventoryView, int, error) {
	var accountResourceRefs []view.AccountResourceRefInventoryView
	total, err := cli.Page(ctx, "v1/accounts/resources/refs", params, &accountResourceRefs)
	return accountResourceRefs, total, err
}
