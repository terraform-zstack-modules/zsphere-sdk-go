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

// UpdateAutoScalingGroup updates AutoScalingGroup
func (cli *ZSClient) UpdateAutoScalingGroup(ctx context.Context, uuid string, params param.UpdateAutoScalingGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/groups", uuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAutoScalingGroup queries AutoScalingGroup list
func (cli *ZSClient) QueryAutoScalingGroup(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupInventoryView, error) {
	var resp []view.AutoScalingGroupInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/groups", params, &resp)
}

func (cli *ZSClient) GetAutoScalingGroup(ctx context.Context, uuid string) (*view.AutoScalingGroupInventoryView, error) {
	var resp view.AutoScalingGroupInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingGroup Pagination
func (cli *ZSClient) PageAutoScalingGroup(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupInventoryView, int, error) {
	var autoScalingGroups []view.AutoScalingGroupInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/groups", params, &autoScalingGroups)
	return autoScalingGroups, total, err
}
// DeleteAutoScalingGroup deletes AutoScalingGroup
func (cli *ZSClient) DeleteAutoScalingGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/autoscaling/groups", uuid, string(deleteMode))
}
// CreateAutoScalingGroup creates AutoScalingGroup
func (cli *ZSClient) CreateAutoScalingGroup(ctx context.Context, params param.CreateAutoScalingGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/autoscaling/groups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
