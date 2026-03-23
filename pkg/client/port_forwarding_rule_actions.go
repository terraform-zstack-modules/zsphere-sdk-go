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

// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-forwarding", uuid, string(deleteMode))
}
// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-forwarding", uuid, string(deleteMode))
}
// AttachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) AttachPortForwardingRule(ctx context.Context, ruleUuid, vmNicUuid string, params param.AttachPortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/port-forwarding/%s/vm-instances/nics/%s", ruleUuid, vmNicUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(ctx context.Context, params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List(ctx, "v1/port-forwarding", params, &resp)
}

func (cli *ZSClient) GetPortForwardingRule(ctx context.Context, uuid string) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.PortForwardingRuleInventoryView
	if err := cli.Get(ctx, "v1/port-forwarding", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortForwardingRule Pagination
func (cli *ZSClient) PagePortForwardingRule(ctx context.Context, params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, int, error) {
	var portForwardingRules []view.PortForwardingRuleInventoryView
	total, err := cli.Page(ctx, "v1/port-forwarding", params, &portForwardingRules)
	return portForwardingRules, total, err
}
// UpdatePortForwardingRule updates PortForwardingRule
func (cli *ZSClient) UpdatePortForwardingRule(ctx context.Context, uuid string, params param.UpdatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/port-forwarding", uuid, "actions", "inventory", map[string]interface{}{
		"updatePortForwardingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(ctx context.Context, params param.CreatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/port-forwarding"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
