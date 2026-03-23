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

// DeleteAccessControlRule deletes AccessControlRule
func (cli *ZSClient) DeleteAccessControlRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/login-control/access-control/rules", uuid, string(deleteMode))
}
// AddAccessControlRule adds AccessControlRule
func (cli *ZSClient) AddAccessControlRule(ctx context.Context, params param.AddAccessControlRuleParam) (*view.AccessControlRuleInventoryView, error) {
	resp := view.AccessControlRuleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/login-control/access-control/rules"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAccessControlRule updates AccessControlRule
func (cli *ZSClient) UpdateAccessControlRule(ctx context.Context, uuid string, params param.UpdateAccessControlRuleParam) (*view.AccessControlRuleInventoryView, error) {
	resp := view.AccessControlRuleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/login-control/access-control/rules", uuid, "actions", "inventory", map[string]interface{}{
		"updateAccessControlRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAccessControlRule queries AccessControlRule list
func (cli *ZSClient) QueryAccessControlRule(ctx context.Context, params *param.QueryParam) ([]view.AccessControlRuleInventoryView, error) {
	var resp []view.AccessControlRuleInventoryView
	return resp, cli.List(ctx, "v1/login-control/access-control/rules", params, &resp)
}

// PageAccessControlRule Pagination
func (cli *ZSClient) PageAccessControlRule(ctx context.Context, params *param.QueryParam) ([]view.AccessControlRuleInventoryView, int, error) {
	var accessControlRules []view.AccessControlRuleInventoryView
	total, err := cli.Page(ctx, "v1/login-control/access-control/rules", params, &accessControlRules)
	return accessControlRules, total, err
}
