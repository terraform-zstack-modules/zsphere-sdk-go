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

// CreateBaremetalBonding creates BaremetalBonding
func (cli *ZSClient) CreateBaremetalBonding(ctx context.Context, params param.CreateBaremetalBondingParam) (*view.BaremetalBondingInventoryView, error) {
	resp := view.BaremetalBondingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/network/bondings"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBaremetalBonding queries BaremetalBonding list
func (cli *ZSClient) QueryBaremetalBonding(ctx context.Context, params *param.QueryParam) ([]view.BaremetalBondingInventoryView, error) {
	var resp []view.BaremetalBondingInventoryView
	return resp, cli.List(ctx, "v1/baremetal/network/bondings", params, &resp)
}

func (cli *ZSClient) GetBaremetalBonding(ctx context.Context, uuid string) (*view.BaremetalBondingInventoryView, error) {
	var resp view.BaremetalBondingInventoryView
	if err := cli.Get(ctx, "v1/baremetal/network/bondings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBaremetalBonding Pagination
func (cli *ZSClient) PageBaremetalBonding(ctx context.Context, params *param.QueryParam) ([]view.BaremetalBondingInventoryView, int, error) {
	var baremetalBondings []view.BaremetalBondingInventoryView
	total, err := cli.Page(ctx, "v1/baremetal/network/bondings", params, &baremetalBondings)
	return baremetalBondings, total, err
}
