// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryAutoScalingRuleTrigger queries AutoScalingRuleTrigger list
func (cli *ZSClient) QueryAutoScalingRuleTrigger(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, error) {
	var resp []view.AutoScalingRuleTriggerInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/groups/rules/trigger", params, &resp)
}

func (cli *ZSClient) GetAutoScalingRuleTrigger(ctx context.Context, uuid string) (*view.AutoScalingRuleTriggerInventoryView, error) {
	var resp view.AutoScalingRuleTriggerInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/groups/rules/trigger", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingRuleTrigger Pagination
func (cli *ZSClient) PageAutoScalingRuleTrigger(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingRuleTriggerInventoryView, int, error) {
	var autoScalingRuleTriggers []view.AutoScalingRuleTriggerInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/groups/rules/trigger", params, &autoScalingRuleTriggers)
	return autoScalingRuleTriggers, total, err
}
// DeleteAutoScalingRuleTrigger deletes AutoScalingRuleTrigger
func (cli *ZSClient) DeleteAutoScalingRuleTrigger(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/autoscaling/groups/rules/triggers", uuid, string(deleteMode))
}
