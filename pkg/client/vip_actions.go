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

// UpdateVip updates Vip
func (cli *ZSClient) UpdateVip(ctx context.Context, uuid string, params param.UpdateVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vips", uuid, "actions", "inventory", map[string]interface{}{
		"updateVip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVip creates Vip
func (cli *ZSClient) CreateVip(ctx context.Context, params param.CreateVipParam) (*view.VipInventoryView, error) {
	resp := view.VipInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vips"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVip queries Vip list
func (cli *ZSClient) QueryVip(ctx context.Context, params *param.QueryParam) ([]view.VipInventoryView, error) {
	var resp []view.VipInventoryView
	return resp, cli.List(ctx, "v1/vips", params, &resp)
}

func (cli *ZSClient) GetVip(ctx context.Context, uuid string) (*view.VipInventoryView, error) {
	var resp view.VipInventoryView
	if err := cli.Get(ctx, "v1/vips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVip Pagination
func (cli *ZSClient) PageVip(ctx context.Context, params *param.QueryParam) ([]view.VipInventoryView, int, error) {
	var vips []view.VipInventoryView
	total, err := cli.Page(ctx, "v1/vips", params, &vips)
	return vips, total, err
}
// DeleteVip deletes Vip
func (cli *ZSClient) DeleteVip(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vips", uuid, string(deleteMode))
}
