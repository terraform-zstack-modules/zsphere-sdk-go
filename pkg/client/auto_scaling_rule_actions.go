// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteAutoScalingRule deletes AutoScalingRule
func (cli *ZSClient) DeleteAutoScalingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/autoscaling/rules", uuid, string(deleteMode))
}
// QueryAutoScalingRule queries AutoScalingRule list
func (cli *ZSClient) QueryAutoScalingRule(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, error) {
	var resp []view.AutoScalingRuleInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/groups/rules", params, &resp)
}

func (cli *ZSClient) GetAutoScalingRule(ctx context.Context, uuid string) (*view.AutoScalingRuleInventoryView, error) {
	var resp view.AutoScalingRuleInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/groups/rules", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingRule Pagination
func (cli *ZSClient) PageAutoScalingRule(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingRuleInventoryView, int, error) {
	var autoScalingRules []view.AutoScalingRuleInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/groups/rules", params, &autoScalingRules)
	return autoScalingRules, total, err
}
// UpdateAutoScalingRule updates AutoScalingRule
func (cli *ZSClient) UpdateAutoScalingRule(ctx context.Context, uuid string, params param.UpdateAutoScalingRuleParam) (*view.AutoScalingRuleInventoryView, error) {
	resp := view.AutoScalingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/rules", uuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
