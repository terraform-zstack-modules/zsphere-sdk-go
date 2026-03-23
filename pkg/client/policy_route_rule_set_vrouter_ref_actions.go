// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryPolicyRouteRuleSetVRouterRef queries PolicyRouteRuleSetVRouterRef list
func (cli *ZSClient) QueryPolicyRouteRuleSetVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetVRouterRefInventoryView, error) {
	var resp []view.PolicyRouteRuleSetVRouterRefInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/rulesets/vrouters/refs", params, &resp)
}

// PagePolicyRouteRuleSetVRouterRef Pagination
func (cli *ZSClient) PagePolicyRouteRuleSetVRouterRef(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetVRouterRefInventoryView, int, error) {
	var policyRouteRuleSetVRouterRefs []view.PolicyRouteRuleSetVRouterRefInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/rulesets/vrouters/refs", params, &policyRouteRuleSetVRouterRefs)
	return policyRouteRuleSetVRouterRefs, total, err
}
