// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVpcFirewallVRouterRef queries VpcFirewallVRouterRef list
func (cli *ZSClient) QueryVpcFirewallVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, error) {
	var resp []view.VpcFirewallVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/vpcfirewalls/vrouters/refs", params, &resp)
}

// PageVpcFirewallVRouterRef Pagination
func (cli *ZSClient) PageVpcFirewallVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.VpcFirewallVRouterRefInventoryView, int, error) {
	var vpcFirewallVRouterRefs []view.VpcFirewallVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/vpcfirewalls/vrouters/refs", params, &vpcFirewallVRouterRefs)
	return vpcFirewallVRouterRefs, total, err
}
