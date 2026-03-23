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

// QuerySftpBackupStorage queries SftpBackupStorage list
func (cli *ZSClient) QuerySftpBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, error) {
	var resp []view.SftpBackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/sftp", params, &resp)
}

func (cli *ZSClient) GetSftpBackupStorage(ctx context.Context, uuid string) (*view.SftpBackupStorageInventoryView, error) {
	var resp view.SftpBackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/sftp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSftpBackupStorage Pagination
func (cli *ZSClient) PageSftpBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.SftpBackupStorageInventoryView, int, error) {
	var sftpBackupStorages []view.SftpBackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/sftp", params, &sftpBackupStorages)
	return sftpBackupStorages, total, err
}
// UpdateSftpBackupStorage updates SftpBackupStorage
func (cli *ZSClient) UpdateSftpBackupStorage(ctx context.Context, uuid string, params param.UpdateSftpBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/sftp", uuid, "actions", "inventory", map[string]interface{}{
		"updateSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ReconnectSftpBackupStorage operates on SftpBackupStorage
func (cli *ZSClient) ReconnectSftpBackupStorage(ctx context.Context, uuid string, params param.ReconnectSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/sftp", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectSftpBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddSftpBackupStorage adds SftpBackupStorage
func (cli *ZSClient) AddSftpBackupStorage(ctx context.Context, params param.AddSftpBackupStorageParam) (*view.SftpBackupStorageInventoryView, error) {
	resp := view.SftpBackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/sftp"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
