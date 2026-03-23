// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVCenterBackupStorage queries VCenterBackupStorage list
func (cli *ZSClient) QueryVCenterBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.VCenterBackupStorageInventoryView, error) {
	var resp []view.VCenterBackupStorageInventoryView
	return resp, cli.List(ctx, "v1/vcenters/backup-storage", params, &resp)
}

func (cli *ZSClient) GetVCenterBackupStorage(ctx context.Context, uuid string) (*view.VCenterBackupStorageInventoryView, error) {
	var resp view.VCenterBackupStorageInventoryView
	if err := cli.Get(ctx, "v1/vcenters/backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterBackupStorage Pagination
func (cli *ZSClient) PageVCenterBackupStorage(ctx context.Context, params *param.QueryParam) ([]view.VCenterBackupStorageInventoryView, int, error) {
	var vCenterBackupStorages []view.VCenterBackupStorageInventoryView
	total, err := cli.Page(ctx, "v1/vcenters/backup-storage", params, &vCenterBackupStorages)
	return vCenterBackupStorages, total, err
}
