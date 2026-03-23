// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryUserTag queries UserTag list
func (cli *ZSClient) QueryUserTag(ctx context.Context, params *param.QueryParam) ([]view.UserTagInventoryView, error) {
	var resp []view.UserTagInventoryView
	return resp, cli.List(ctx, "v1/user-tags", params, &resp)
}

func (cli *ZSClient) GetUserTag(ctx context.Context, uuid string) (*view.UserTagInventoryView, error) {
	var resp view.UserTagInventoryView
	if err := cli.Get(ctx, "v1/user-tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUserTag Pagination
func (cli *ZSClient) PageUserTag(ctx context.Context, params *param.QueryParam) ([]view.UserTagInventoryView, int, error) {
	var userTags []view.UserTagInventoryView
	total, err := cli.Page(ctx, "v1/user-tags", params, &userTags)
	return userTags, total, err
}
