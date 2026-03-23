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

// CreatePriceTable creates PriceTable
func (cli *ZSClient) CreatePriceTable(ctx context.Context, params param.CreatePriceTableParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/billings/price-tables"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPriceTable queries PriceTable list
func (cli *ZSClient) QueryPriceTable(ctx context.Context, params *param.QueryParam) ([]view.PriceTableInventoryView, error) {
	var resp []view.PriceTableInventoryView
	return resp, cli.List(ctx, "v1/billings/price-tables", params, &resp)
}

// PagePriceTable Pagination
func (cli *ZSClient) PagePriceTable(ctx context.Context, params *param.QueryParam) ([]view.PriceTableInventoryView, int, error) {
	var priceTables []view.PriceTableInventoryView
	total, err := cli.Page(ctx, "v1/billings/price-tables", params, &priceTables)
	return priceTables, total, err
}
// UpdatePriceTable updates PriceTable
func (cli *ZSClient) UpdatePriceTable(ctx context.Context, uuid string, params param.UpdatePriceTableParam) (*view.PriceTableInventoryView, error) {
	resp := view.PriceTableInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/billings/price-tables", uuid, "actions", "inventory", map[string]interface{}{
		"updatePriceTable": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePriceTable deletes PriceTable
func (cli *ZSClient) DeletePriceTable(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/billings/price-tables", uuid, string(deleteMode))
}
