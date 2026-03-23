// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryRoleAccountRef queries RoleAccountRef list
func (cli *ZSClient) QueryRoleAccountRef(ctx context.Context, params *param.QueryParam) ([]view.RoleAccountRefInventoryView, error) {
	var resp []view.RoleAccountRefInventoryView
	return resp, cli.List(ctx, "v1/identities/role-account-refs", params, &resp)
}

// PageRoleAccountRef Pagination
func (cli *ZSClient) PageRoleAccountRef(ctx context.Context, params *param.QueryParam) ([]view.RoleAccountRefInventoryView, int, error) {
	var roleAccountRefs []view.RoleAccountRefInventoryView
	total, err := cli.Page(ctx, "v1/identities/role-account-refs", params, &roleAccountRefs)
	return roleAccountRefs, total, err
}
