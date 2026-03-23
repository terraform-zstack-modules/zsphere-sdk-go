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

// CreateL3Network creates L3Network
func (cli *ZSClient) CreateL3Network(ctx context.Context, params param.CreateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteL3Network deletes L3Network
func (cli *ZSClient) DeleteL3Network(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks", uuid, string(deleteMode))
}
// QueryL3Network queries L3Network list
func (cli *ZSClient) QueryL3Network(ctx context.Context, params *param.QueryParam) ([]view.L3NetworkInventoryView, error) {
	var resp []view.L3NetworkInventoryView
	return resp, cli.List(ctx, "v1/l3-networks", params, &resp)
}

func (cli *ZSClient) GetL3Network(ctx context.Context, uuid string) (*view.L3NetworkInventoryView, error) {
	var resp view.L3NetworkInventoryView
	if err := cli.Get(ctx, "v1/l3-networks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageL3Network Pagination
func (cli *ZSClient) PageL3Network(ctx context.Context, params *param.QueryParam) ([]view.L3NetworkInventoryView, int, error) {
	var l3Networks []view.L3NetworkInventoryView
	total, err := cli.Page(ctx, "v1/l3-networks", params, &l3Networks)
	return l3Networks, total, err
}
// UpdateL3Network updates L3Network
func (cli *ZSClient) UpdateL3Network(ctx context.Context, uuid string, params param.UpdateL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/l3-networks", uuid, "actions", "inventory", map[string]interface{}{
		"updateL3Network": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
