// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateAutoScalingGroupInstance updates AutoScalingGroupInstance
func (cli *ZSClient) UpdateAutoScalingGroupInstance(ctx context.Context, instanceUuid string, params param.UpdateAutoScalingGroupInstanceParam) (*view.AutoScalingGroupInstanceInventoryView, error) {
	resp := view.AutoScalingGroupInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/groups/instances", instanceUuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingGroupInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAutoScalingGroupInstance queries AutoScalingGroupInstance list
func (cli *ZSClient) QueryAutoScalingGroupInstance(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupInstanceInventoryView, error) {
	var resp []view.AutoScalingGroupInstanceInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/groups/instances", params, &resp)
}

func (cli *ZSClient) GetAutoScalingGroupInstance(ctx context.Context, uuid string) (*view.AutoScalingGroupInstanceInventoryView, error) {
	var resp view.AutoScalingGroupInstanceInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/groups/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingGroupInstance Pagination
func (cli *ZSClient) PageAutoScalingGroupInstance(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingGroupInstanceInventoryView, int, error) {
	var autoScalingGroupInstances []view.AutoScalingGroupInstanceInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/groups/instances", params, &autoScalingGroupInstances)
	return autoScalingGroupInstances, total, err
}
// DeleteAutoScalingGroupInstance deletes AutoScalingGroupInstance
func (cli *ZSClient) DeleteAutoScalingGroupInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/autoscaling/groups/instances", uuid, string(deleteMode))
}
