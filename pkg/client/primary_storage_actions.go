// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdatePrimaryStorage updates PrimaryStorage
func (cli *ZSClient) UpdatePrimaryStorage(ctx context.Context, uuid string, params param.UpdatePrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "actions", "inventory", map[string]interface{}{
		"updatePrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPrimaryStorage queries PrimaryStorage list
func (cli *ZSClient) QueryPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, error) {
	var resp []view.PrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage", params, &resp)
}

// PagePrimaryStorage Pagination
func (cli *ZSClient) PagePrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.PrimaryStorageInventoryView, int, error) {
	var primaryStorages []view.PrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage", params, &primaryStorages)
	return primaryStorages, total, err
}
// ReconnectPrimaryStorage operates on PrimaryStorage
func (cli *ZSClient) ReconnectPrimaryStorage(ctx context.Context, uuid string, params param.ReconnectPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePrimaryStorage deletes PrimaryStorage
func (cli *ZSClient) DeletePrimaryStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage", uuid, string(deleteMode))
}
