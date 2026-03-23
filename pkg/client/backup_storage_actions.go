// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateBackupStorage updates BackupStorage
func (cli *ZSClient) UpdateBackupStorage(ctx context.Context, uuid string, params param.UpdateBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "actions", "inventory", map[string]interface{}{
		"updateBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteBackupStorage deletes BackupStorage
func (cli *ZSClient) DeleteBackupStorage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/backup-storage", uuid, string(deleteMode))
}
// QueryBackupStorage queries BackupStorage list
func (cli *ZSClient) QueryBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage", params, &resp)
}

func (cli *ZSClient) GetBackupStorage(ctx context.Context, uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBackupStorage Pagination
func (cli *ZSClient) PageBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var backupStorages []view.BackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage", params, &backupStorages)
	return backupStorages, total, err
}
// ReconnectBackupStorage operates on BackupStorage
func (cli *ZSClient) ReconnectBackupStorage(ctx context.Context, uuid string, params param.ReconnectBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
