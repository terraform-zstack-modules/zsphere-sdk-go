// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateCephBackupStorageMon updates CephBackupStorageMon
func (cli *ZSClient) UpdateCephBackupStorageMon(ctx context.Context, monUuid string, params param.UpdateCephBackupStorageMonParam) (*view.CephBackupStorageInventoryView, error) {
	resp := view.CephBackupStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/ceph/mons", monUuid, "actions", "inventory", map[string]interface{}{
		"updateCephBackupStorageMon": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
