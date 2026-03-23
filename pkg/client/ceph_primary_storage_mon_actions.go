// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateCephPrimaryStorageMon updates CephPrimaryStorageMon
func (cli *ZSClient) UpdateCephPrimaryStorageMon(ctx context.Context, monUuid string, params param.UpdateCephPrimaryStorageMonParam) (*view.CephPrimaryStorageInventoryView, error) {
	resp := view.CephPrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/ceph/mons", monUuid, "actions", "inventory", map[string]interface{}{
		"updateCephPrimaryStorageMon": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
