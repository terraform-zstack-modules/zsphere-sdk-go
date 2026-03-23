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

// QueryPolicyRouteRule queries PolicyRouteRule list
func (cli *ZSClient) QueryPolicyRouteRule(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, error) {
	var resp []view.PolicyRouteRuleInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/rules", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRule(ctx context.Context, uuid string) (*view.PolicyRouteRuleInventoryView, error) {
	var resp view.PolicyRouteRuleInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteRule Pagination
func (cli *ZSClient) PagePolicyRouteRule(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleInventoryView, int, error) {
	var policyRouteRules []view.PolicyRouteRuleInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/rules", params, &policyRouteRules)
	return policyRouteRules, total, err
}
// CreatePolicyRouteRule creates PolicyRouteRule
func (cli *ZSClient) CreatePolicyRouteRule(ctx context.Context, params param.CreatePolicyRouteRuleParam) (*view.PolicyRouteRuleInventoryView, error) {
	resp := view.PolicyRouteRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/policy-routes/rules"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteRule deletes PolicyRouteRule
func (cli *ZSClient) DeletePolicyRouteRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/policy-routes/rules", uuid, string(deleteMode))
}
