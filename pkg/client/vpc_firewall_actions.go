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

// UpdateVpcFirewall updates VpcFirewall
func (cli *ZSClient) UpdateVpcFirewall(ctx context.Context, uuid string, params param.UpdateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vpcfirewalls", uuid, "actions", "inventory", map[string]interface{}{
		"updateVpcFirewall": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVpcFirewall creates VpcFirewall
func (cli *ZSClient) CreateVpcFirewall(ctx context.Context, params param.CreateVpcFirewallParam) (*view.VpcFirewallInventoryView, error) {
	resp := view.VpcFirewallInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpcfirewalls"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVpcFirewall queries VpcFirewall list
func (cli *ZSClient) QueryVpcFirewall(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallInventoryView, error) {
	var resp []view.VpcFirewallInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls", params, &resp)
}

func (cli *ZSClient) GetVpcFirewall(ctx context.Context, uuid string) (*view.VpcFirewallInventoryView, error) {
	var resp view.VpcFirewallInventoryView
	if err := cli.Get(ctx, "v1/vpcfirewalls", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVpcFirewall Pagination
func (cli *ZSClient) PageVpcFirewall(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallInventoryView, int, error) {
	var vpcFirewalls []view.VpcFirewallInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls", params, &vpcFirewalls)
	return vpcFirewalls, total, err
}
