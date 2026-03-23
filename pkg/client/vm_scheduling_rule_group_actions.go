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

// UpdateVmSchedulingRuleGroup updates VmSchedulingRuleGroup
func (cli *ZSClient) UpdateVmSchedulingRuleGroup(ctx context.Context, uuid string, params param.UpdateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vmSchedulingRuleGroup", uuid, "update", "inventory", map[string]interface{}{
		"updateVmSchedulingRuleGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVmSchedulingRuleGroup deletes VmSchedulingRuleGroup
func (cli *ZSClient) DeleteVmSchedulingRuleGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRuleGroup", uuid, string(deleteMode))
}
// QueryVmSchedulingRuleGroup queries VmSchedulingRuleGroup list
func (cli *ZSClient) QueryVmSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, error) {
	var resp []view.VmSchedulingRuleGroupInventoryView
	return resp, cli.List(ctx, "v1/query/vm/schedulingRule/group", params, &resp)
}

// PageVmSchedulingRuleGroup Pagination
func (cli *ZSClient) PageVmSchedulingRuleGroup(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleGroupInventoryView, int, error) {
	var vmSchedulingRuleGroups []view.VmSchedulingRuleGroupInventoryView
	total, err := cli.Page(ctx, "v1/query/vm/schedulingRule/group", params, &vmSchedulingRuleGroups)
	return vmSchedulingRuleGroups, total, err
}
// CreateVmSchedulingRuleGroup creates VmSchedulingRuleGroup
func (cli *ZSClient) CreateVmSchedulingRuleGroup(ctx context.Context, params param.CreateVmSchedulingRuleGroupParam) (*view.VmSchedulingRuleGroupInventoryView, error) {
	resp := view.VmSchedulingRuleGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vmSchedulingRuleGroup"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
