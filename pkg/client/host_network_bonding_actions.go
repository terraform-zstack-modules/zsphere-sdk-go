// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryHostNetworkBonding queries HostNetworkBonding list
func (cli *ZSClient) QueryHostNetworkBonding(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, error) {
	var resp []view.HostNetworkBondingInventoryView
	return resp, cli.List(ctx, "v1/hosts/bondings", params, &resp)
}

func (cli *ZSClient) GetHostNetworkBonding(ctx context.Context, uuid string) (*view.HostNetworkBondingInventoryView, error) {
	var resp view.HostNetworkBondingInventoryView
	if err := cli.Get(ctx, "v1/hosts/bondings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostNetworkBonding Pagination
func (cli *ZSClient) PageHostNetworkBonding(ctx context.Context, params *param.QueryParam) ([]view.HostNetworkBondingInventoryView, int, error) {
	var hostNetworkBondings []view.HostNetworkBondingInventoryView
	total, err := cli.Page(ctx, "v1/hosts/bondings", params, &hostNetworkBondings)
	return hostNetworkBondings, total, err
}
