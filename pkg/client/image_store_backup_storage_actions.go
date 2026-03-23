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

// UpdateImageStoreBackupStorage updates ImageStoreBackupStorage
func (cli *ZSClient) UpdateImageStoreBackupStorage(ctx context.Context, uuid string, params param.UpdateImageStoreBackupStorageParam) (*view.BackupStorageInventoryView, error) {
	resp := view.BackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/image-store", uuid, "actions", "inventory", map[string]interface{}{
		"updateImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImageStoreBackupStorage queries ImageStoreBackupStorage list
func (cli *ZSClient) QueryImageStoreBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, error) {
	var resp []view.ImageStoreBackupStorageInventoryView
	return resp, cli.List(ctx, "v1/backup-storage/image-store", params, &resp)
}

func (cli *ZSClient) GetImageStoreBackupStorage(ctx context.Context, uuid string) (*view.ImageStoreBackupStorageInventoryView, error) {
	var resp view.ImageStoreBackupStorageInventoryView
	if err := cli.Get(ctx, "v1/backup-storage/image-store", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageStoreBackupStorage Pagination
func (cli *ZSClient) PageImageStoreBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.ImageStoreBackupStorageInventoryView, int, error) {
	var imageStoreBackupStorages []view.ImageStoreBackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/backup-storage/image-store", params, &imageStoreBackupStorages)
	return imageStoreBackupStorages, total, err
}
// ReconnectImageStoreBackupStorage operates on ImageStoreBackupStorage
func (cli *ZSClient) ReconnectImageStoreBackupStorage(ctx context.Context, uuid string, params param.ReconnectImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/image-store", uuid, "actions", "inventory", map[string]interface{}{
		"reconnectImageStoreBackupStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddImageStoreBackupStorage adds ImageStoreBackupStorage
func (cli *ZSClient) AddImageStoreBackupStorage(ctx context.Context, params param.AddImageStoreBackupStorageParam) (*view.ImageStoreBackupStorageInventoryView, error) {
	resp := view.ImageStoreBackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/image-store"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
