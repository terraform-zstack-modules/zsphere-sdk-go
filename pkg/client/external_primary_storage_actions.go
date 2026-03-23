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

// UpdateExternalPrimaryStorage updates ExternalPrimaryStorage
func (cli *ZSClient) UpdateExternalPrimaryStorage(ctx context.Context, uuid string, params param.UpdateExternalPrimaryStorageParam) (*view.ExternalPrimaryStorageInventoryView, error) {
	resp := view.ExternalPrimaryStorageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/addon", uuid, "actions", "inventory", map[string]interface{}{
		"updateExternalPrimaryStorage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddExternalPrimaryStorage adds ExternalPrimaryStorage
func (cli *ZSClient) AddExternalPrimaryStorage(ctx context.Context, params param.AddExternalPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/addon"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
