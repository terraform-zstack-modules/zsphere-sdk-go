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

// CreatePolicyRouteRuleSet creates PolicyRouteRuleSet
func (cli *ZSClient) CreatePolicyRouteRuleSet(ctx context.Context, params param.CreatePolicyRouteRuleSetParam) (*view.PolicyRouteRuleSetInventoryView, error) {
	resp := view.PolicyRouteRuleSetInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/policy-routes/rulesets"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePolicyRouteRuleSet updates PolicyRouteRuleSet
func (cli *ZSClient) UpdatePolicyRouteRuleSet(ctx context.Context, uuid string, params param.UpdatePolicyRouteRuleSetParam) (*view.PolicyRouteRuleSetInventoryView, error) {
	resp := view.PolicyRouteRuleSetInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/policy-routes/ruleSets", uuid, "actions", "inventory", map[string]interface{}{
		"updatePolicyRouteRuleSet": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteRuleSet deletes PolicyRouteRuleSet
func (cli *ZSClient) DeletePolicyRouteRuleSet(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/policy-routes/ruleSets", uuid, string(deleteMode))
}
// QueryPolicyRouteRuleSet queries PolicyRouteRuleSet list
func (cli *ZSClient) QueryPolicyRouteRuleSet(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, error) {
	var resp []view.PolicyRouteRuleSetInventoryView
	return resp, cli.List(ctx, "v1/policy-routes/rulesets", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteRuleSet(ctx context.Context, uuid string) (*view.PolicyRouteRuleSetInventoryView, error) {
	var resp view.PolicyRouteRuleSetInventoryView
	if err := cli.Get(ctx, "v1/policy-routes/rulesets", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteRuleSet Pagination
func (cli *ZSClient) PagePolicyRouteRuleSet(ctx context.Context, params *param.QueryParam) ([]view.PolicyRouteRuleSetInventoryView, int, error) {
	var policyRouteRuleSets []view.PolicyRouteRuleSetInventoryView
	total, err := cli.Page(ctx, "v1/policy-routes/rulesets", params, &policyRouteRuleSets)
	return policyRouteRuleSets, total, err
}
