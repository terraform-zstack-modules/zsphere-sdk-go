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

// ValidateVmSchedulingRule operates on VmSchedulingRule
func (cli *ZSClient) ValidateVmSchedulingRule(ctx context.Context, params param.ValidateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/validate/vmSchedulingRule", "", "", map[string]interface{}{
		"validateVmSchedulingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmSchedulingRule updates VmSchedulingRule
func (cli *ZSClient) UpdateVmSchedulingRule(ctx context.Context, uuid string, params param.UpdateVmSchedulingRuleParam) (*view.VmSchedulingRuleInventoryView, error) {
	resp := view.VmSchedulingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vmSchedulingRule", uuid, "update", "inventory", map[string]interface{}{
		"updateVmSchedulingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmSchedulingRule queries VmSchedulingRule list
func (cli *ZSClient) QueryVmSchedulingRule(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, error) {
	var resp []view.VmSchedulingRuleInventoryView
	return resp, cli.List(ctx, "v1/query/vm/schedulingRule", params, &resp)
}

// PageVmSchedulingRule Pagination
func (cli *ZSClient) PageVmSchedulingRule(ctx context.Context, params *param.QueryParam) ([]view.VmSchedulingRuleInventoryView, int, error) {
	var vmSchedulingRules []view.VmSchedulingRuleInventoryView
	total, err := cli.Page(ctx, "v1/query/vm/schedulingRule", params, &vmSchedulingRules)
	return vmSchedulingRules, total, err
}
// RemoveVmSchedulingRule removes VmSchedulingRule
func (cli *ZSClient) RemoveVmSchedulingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vmSchedulingRule", uuid, string(deleteMode))
}
// CreateVmSchedulingRule creates VmSchedulingRule
func (cli *ZSClient) CreateVmSchedulingRule(ctx context.Context, params param.CreateVmSchedulingRuleParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vmsSchedulingRule"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
