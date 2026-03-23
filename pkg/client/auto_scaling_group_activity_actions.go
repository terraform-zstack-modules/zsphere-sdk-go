// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryAutoScalingGroupActivity queries AutoScalingGroupActivity list
func (cli *ZSClient) QueryAutoScalingGroupActivity(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupActivityInventoryView, error) {
	var resp []view.AutoScalingGroupActivityInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/groups/activities", params, &resp)
}

func (cli *ZSClient) GetAutoScalingGroupActivity(ctx context.Context, uuid string) (*view.AutoScalingGroupActivityInventoryView, error) {
	var resp view.AutoScalingGroupActivityInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/groups/activities", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingGroupActivity Pagination
func (cli *ZSClient) PageAutoScalingGroupActivity(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupActivityInventoryView, int, error) {
	var autoScalingGroupActivities []view.AutoScalingGroupActivityInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/groups/activities", params, &autoScalingGroupActivities)
	return autoScalingGroupActivities, total, err
}
