// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteL2Network deletes L2Network
func (cli *ZSClient) DeleteL2Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l2-networks", uuid, string(deleteMode))
}
// UpdateL2Network updates L2Network
func (cli *ZSClient) UpdateL2Network(ctx context.Context, uuid string, params param.UpdateL2NetworkParam) (*view.L2NetworkInventoryView, error) {
	resp := view.L2NetworkInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l2-networks", uuid, "actions", "inventory", map[string]interface{}{
		"updateL2Network": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryL2Network queries L2Network list
func (cli *ZSClient) QueryL2Network(ctx context.Context, params *param.QueryParam) ([]view.L2NetworkInventoryView, error) {
	var resp []view.L2NetworkInventoryView
	return resp, cli.List(ctx, "v1/l2-networks", params, &resp)
}

func (cli *ZSClient) GetL2Network(ctx context.Context, uuid string) (*view.L2NetworkInventoryView, error) {
	var resp view.L2NetworkInventoryView
	if err := cli.Get(ctx, "v1/l2-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL2Network Pagination
func (cli *ZSClient) PageL2Network(ctx context.Context, params *param.QueryParam) ([]view.L2NetworkInventoryView, int, error) {
	var l2Networks []view.L2NetworkInventoryView
	total, err := cli.Page(ctx, "v1/l2-networks", params, &l2Networks)
	return l2Networks, total, err
}
