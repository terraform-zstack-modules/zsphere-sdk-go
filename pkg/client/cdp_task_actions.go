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

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(ctx context.Context, params param.CreateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cdp-backup-storage/task"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCdpTask queries CdpTask list
func (cli *ZSClient) QueryCdpTask(ctx context.Context, params *param.QueryParam) ([]view.CdpTaskInventoryView, error) {
	var resp []view.CdpTaskInventoryView
	return resp, cli.List(ctx, "v1/cdp-task", params, &resp)
}

func (cli *ZSClient) GetCdpTask(ctx context.Context, uuid string) (*view.CdpTaskInventoryView, error) {
	var resp view.CdpTaskInventoryView
	if err := cli.Get(ctx, "v1/cdp-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCdpTask Pagination
func (cli *ZSClient) PageCdpTask(ctx context.Context, params *param.QueryParam) ([]view.CdpTaskInventoryView, int, error) {
	var cdpTasks []view.CdpTaskInventoryView
	total, err := cli.Page(ctx, "v1/cdp-task", params, &cdpTasks)
	return cdpTasks, total, err
}
// DeleteCdpTask deletes CdpTask
func (cli *ZSClient) DeleteCdpTask(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cdp-task", uuid, string(deleteMode))
}
// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(ctx context.Context, uuid string, params param.UpdateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/cdp-backup-storage/task", uuid, "actions", "inventory", map[string]interface{}{
		"updateCdpTask": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
