// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryPolicyRouteTableVRouterRef queries PolicyRouteTableVRouterRef list
func (cli *ZSClient) QueryPolicyRouteTableVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteTableVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/tables/vrouters/refs", params, &resp)
}

// PagePolicyRouteTableVRouterRef Pagination
func (cli *ZSClient) PagePolicyRouteTableVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteTableVRouterRefInventoryView, int, error) {
	var policyRouteTableVRouterRefs []view.PolicyRouteTableVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/tables/vrouters/refs", params, &policyRouteTableVRouterRefs)
	return policyRouteTableVRouterRefs, total, err
}
