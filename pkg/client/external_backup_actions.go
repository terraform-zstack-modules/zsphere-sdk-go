// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryExternalBackup queries ExternalBackup list
func (cli *ZSClient) QueryExternalBackup(ctx context.Context, params *param.QueryParam) ([]view.ExternalBackupInventoryView, error) {
	var resp []view.ExternalBackupInventoryView
	return resp, cli.List(ctx, "v1/externalbackup", params, &resp)
}

func (cli *ZSClient) GetExternalBackup(ctx context.Context, uuid string) (*view.ExternalBackupInventoryView, error) {
	var resp view.ExternalBackupInventoryView
	if err := cli.Get(ctx, "v1/externalbackup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageExternalBackup Pagination
func (cli *ZSClient) PageExternalBackup(ctx context.Context, params *param.QueryParam) ([]view.ExternalBackupInventoryView, int, error) {
	var externalBackups []view.ExternalBackupInventoryView
	total, err := cli.Page(ctx, "v1/externalbackup", params, &externalBackups)
	return externalBackups, total, err
}
// DeleteExternalBackup deletes ExternalBackup
func (cli *ZSClient) DeleteExternalBackup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/externalbackup", uuid, string(deleteMode))
}
