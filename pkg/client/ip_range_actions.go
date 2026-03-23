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

// AddIpRange adds IpRange
func (cli *ZSClient) AddIpRange(ctx context.Context, l3NetworkUuid string, params param.AddIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/ip-ranges", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIpRange queries IpRange list
func (cli *ZSClient) QueryIpRange(ctx context.Context, params *param.QueryParam) ([]view.IpRangeInventoryView, error) {
	var resp []view.IpRangeInventoryView
	return resp, cli.List(ctx, "v1/l3-networks/ip-ranges", params, &resp)
}

func (cli *ZSClient) GetIpRange(ctx context.Context, uuid string) (*view.IpRangeInventoryView, error) {
	var resp view.IpRangeInventoryView
	if err := cli.Get(ctx, "v1/l3-networks/ip-ranges", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIpRange Pagination
func (cli *ZSClient) PageIpRange(ctx context.Context, params *param.QueryParam) ([]view.IpRangeInventoryView, int, error) {
	var ipRanges []view.IpRangeInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks/ip-ranges", params, &ipRanges)
	return ipRanges, total, err
}
// DeleteIpRange deletes IpRange
func (cli *ZSClient) DeleteIpRange(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/ip-ranges", uuid, string(deleteMode))
}
// UpdateIpRange updates IpRange
func (cli *ZSClient) UpdateIpRange(ctx context.Context, uuid string, params param.UpdateIpRangeParam) (*view.IpRangeInventoryView, error) {
	resp := view.IpRangeInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks/ip-ranges", uuid, "actions", "inventory", map[string]interface{}{
		"updateIpRange": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
