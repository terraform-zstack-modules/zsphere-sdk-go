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

// CreateCbtTask creates CbtTask
func (cli *ZSClient) CreateCbtTask(ctx context.Context, params param.CreateCbtTaskParam) (*view.CbtTaskInventoryView, error) {
	resp := view.CbtTaskInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cbt-task/create"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(ctx context.Context, params *param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List(ctx, "v1/cbt-task", params, &resp)
}

func (cli *ZSClient) GetCbtTask(ctx context.Context, uuid string) (*view.CbtTaskInventoryView, error) {
	var resp view.CbtTaskInventoryView
	if err := cli.Get(ctx, "v1/cbt-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCbtTask Pagination
func (cli *ZSClient) PageCbtTask(ctx context.Context, params *param.QueryParam) ([]view.CbtTaskInventoryView, int, error) {
	var cbtTasks []view.CbtTaskInventoryView
	total, err := cli.Page(ctx, "v1/cbt-task", params, &cbtTasks)
	return cbtTasks, total, err
}
// DeleteCbtTask deletes CbtTask
func (cli *ZSClient) DeleteCbtTask(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cbt-task", uuid, string(deleteMode))
}
