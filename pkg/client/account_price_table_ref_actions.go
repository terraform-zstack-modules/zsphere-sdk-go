// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryAccountPriceTableRef queries AccountPriceTableRef list
func (cli *ZSClient) QueryAccountPriceTableRef(ctx context.Context, params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, error) {
	var resp []view.AccountPriceTableRefInventoryView
	return resp, cli.List(ctx, "v1/accounts/price-tables/refs", params, &resp)
}

// PageAccountPriceTableRef Pagination
func (cli *ZSClient) PageAccountPriceTableRef(ctx context.Context, params *param.QueryParam) ([]view.AccountPriceTableRefInventoryView, int, error) {
	var accountPriceTableRefs []view.AccountPriceTableRefInventoryView
	total, err := cli.Page(ctx, "v1/accounts/price-tables/refs", params, &accountPriceTableRefs)
	return accountPriceTableRefs, total, err
}
