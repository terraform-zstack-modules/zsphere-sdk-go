// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVpcSnatState queries VpcSnatState list
func (cli *ZSClient) QueryVpcSnatState(ctx context.Context, params *param.QueryParam) ([]view.VpcSnatStateInventoryView, error) {
	var resp []view.VpcSnatStateInventoryView
	return resp, cli.List(ctx, "v1/vpc/virtual-routers/networkservicestate/snat", params, &resp)
}

func (cli *ZSClient) GetVpcSnatState(ctx context.Context, uuid string) (*view.VpcSnatStateInventoryView, error) {
	var resp view.VpcSnatStateInventoryView
	if err := cli.Get(ctx, "v1/vpc/virtual-routers/networkservicestate/snat", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcSnatState Pagination
func (cli *ZSClient) PageVpcSnatState(ctx context.Context, params *param.QueryParam) ([]view.VpcSnatStateInventoryView, int, error) {
	var vpcSnatStates []view.VpcSnatStateInventoryView
	total, err := cli.Page(ctx, "v1/vpc/virtual-routers/networkservicestate/snat", params, &vpcSnatStates)
	return vpcSnatStates, total, err
}
