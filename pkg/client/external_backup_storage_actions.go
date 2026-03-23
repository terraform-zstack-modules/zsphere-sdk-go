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

// AddExternalBackupStorage adds ExternalBackupStorage
func (cli *ZSClient) AddExternalBackupStorage(ctx context.Context, params param.AddExternalBackupStorageParam) (*view.ExternalBackupStorageInventoryView, error) {
	resp := view.ExternalBackupStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/backup-storage/addon"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
