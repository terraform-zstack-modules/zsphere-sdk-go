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

// QueryCephBackupStorage queries CephBackupStorage list
func (cli *ZSClient) QueryCephBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, error) {
	var resp []view.BackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/ceph", params, &resp)
}

func (cli *ZSClient) GetCephBackupStorage(ctx context.Context, uuid string) (*view.BackupStorageInventoryView, error) {
	var resp view.BackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/ceph", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephBackupStorage Pagination
func (cli *ZSClient) PageCephBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.BackupStorageInventoryView, int, error) {
	var cephBackupStorages []view.BackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/ceph", params, &cephBackupStorages)
	return cephBackupStorages, total, err
}
// AddCephBackupStorage adds CephBackupStorage
func (cli *ZSClient) AddCephBackupStorage(ctx context.Context, params param.AddCephBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/ceph"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
