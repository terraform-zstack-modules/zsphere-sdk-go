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

// SyncVolumeBackup operates on VolumeBackup
func (cli *ZSClient) SyncVolumeBackup(ctx context.Context, imageStoreUuid string, params param.SyncVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-backups/imageStore", imageStoreUuid, "actions", "result", map[string]interface{}{
		"syncVolumeBackup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolumeBackup queries VolumeBackup list
func (cli *ZSClient) QueryVolumeBackup(ctx context.Context, params *param.QueryParam) ([]view.VolumeBackupInventoryView, error) {
	var resp []view.VolumeBackupInventoryView
	return resp, cli.List(ctx, "v1/volume-backups", params, &resp)
}

func (cli *ZSClient) GetVolumeBackup(ctx context.Context, uuid string) (*view.VolumeBackupInventoryView, error) {
	var resp view.VolumeBackupInventoryView
	if err := cli.Get(ctx, "v1/volume-backups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeBackup Pagination
func (cli *ZSClient) PageVolumeBackup(ctx context.Context, params *param.QueryParam) ([]view.VolumeBackupInventoryView, int, error) {
	var volumeBackups []view.VolumeBackupInventoryView
	total, err := cli.Page(ctx, "v1/volume-backups", params, &volumeBackups)
	return volumeBackups, total, err
}
// DeleteVolumeBackup deletes VolumeBackup
func (cli *ZSClient) DeleteVolumeBackup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volume-backups", uuid, string(deleteMode))
}
// CreateVolumeBackup creates VolumeBackup
func (cli *ZSClient) CreateVolumeBackup(ctx context.Context, volumeUuid string, params param.CreateVolumeBackupParam) (*view.VolumeBackupInventoryView, error) {
	resp := view.VolumeBackupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/%s/volume-backups", volumeUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumeBackupAsync Async
func (cli *ZSClient) CreateVolumeBackupAsync(ctx context.Context, params param.CreateVolumeBackupParam) (string, error) {

	resource := "volumes/{volumeUuid}/volume-backups"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
