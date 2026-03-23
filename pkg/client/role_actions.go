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

// CreateRole creates Role
func (cli *ZSClient) CreateRole(ctx context.Context, params param.CreateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/identities/roles"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteRole deletes Role
func (cli *ZSClient) DeleteRole(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/identities/roles", uuid, string(deleteMode))
}
// QueryRole queries Role list
func (cli *ZSClient) QueryRole(ctx context.Context, params *param.QueryParam) ([]view.RoleInventoryView, error) {
	var resp []view.RoleInventoryView
	return resp, cli.List(ctx, "v1/identities/roles", params, &resp)
}

func (cli *ZSClient) GetRole(ctx context.Context, uuid string) (*view.RoleInventoryView, error) {
	var resp view.RoleInventoryView
	if err := cli.Get(ctx, "v1/identities/roles", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageRole Pagination
func (cli *ZSClient) PageRole(ctx context.Context, params *param.QueryParam) ([]view.RoleInventoryView, int, error) {
	var roles []view.RoleInventoryView
	total, err := cli.Page(ctx, "v1/identities/roles", params, &roles)
	return roles, total, err
}
// UpdateRole updates Role
func (cli *ZSClient) UpdateRole(ctx context.Context, uuid string, params param.UpdateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/identities/roles", uuid, "actions", "inventory", map[string]interface{}{
		"updateRole": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
