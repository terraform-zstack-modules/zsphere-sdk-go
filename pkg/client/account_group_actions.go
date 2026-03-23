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

// CreateAccountGroup creates AccountGroup
func (cli *ZSClient) CreateAccountGroup(ctx context.Context, params param.CreateAccountGroupParam) (*view.AccountGroupInventoryView, error) {
	resp := view.AccountGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/account-groups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateAccountGroup updates AccountGroup
func (cli *ZSClient) UpdateAccountGroup(ctx context.Context, uuid string, params param.UpdateAccountGroupParam) (*view.AccountGroupInventoryView, error) {
	resp := view.AccountGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/account-groups", uuid, "actions", "inventory", map[string]interface{}{
		"updateAccountGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAccountGroup deletes AccountGroup
func (cli *ZSClient) DeleteAccountGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/account-groups", uuid, string(deleteMode))
}
// QueryAccountGroup queries AccountGroup list
func (cli *ZSClient) QueryAccountGroup(ctx context.Context, params *param.QueryParam) ([]view.AccountGroupInventoryView, error) {
	var resp []view.AccountGroupInventoryView
	return resp, cli.List(ctx, "v1/account-groups", params, &resp)
}

func (cli *ZSClient) GetAccountGroup(ctx context.Context, uuid string) (*view.AccountGroupInventoryView, error) {
	var resp view.AccountGroupInventoryView
	if err := cli.Get(ctx, "v1/account-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAccountGroup Pagination
func (cli *ZSClient) PageAccountGroup(ctx context.Context, params *param.QueryParam) ([]view.AccountGroupInventoryView, int, error) {
	var accountGroups []view.AccountGroupInventoryView
	total, err := cli.Page(ctx, "v1/account-groups", params, &accountGroups)
	return accountGroups, total, err
}
