// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryAddressPool queries AddressPool list
func (cli *ZSClient) QueryAddressPool(ctx context.Context, params *param.QueryParam) ([]view.AddressPoolInventoryView, error) {
	var resp []view.AddressPoolInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/address-pools", params, &resp)
}

func (cli *ZSClient) GetAddressPool(ctx context.Context, uuid string) (*view.AddressPoolInventoryView, error) {
	var resp view.AddressPoolInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/address-pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAddressPool Pagination
func (cli *ZSClient) PageAddressPool(ctx context.Context, params *param.QueryParam) ([]view.AddressPoolInventoryView, int, error) {
	var addressPools []view.AddressPoolInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/address-pools", params, &addressPools)
	return addressPools, total, err
}
