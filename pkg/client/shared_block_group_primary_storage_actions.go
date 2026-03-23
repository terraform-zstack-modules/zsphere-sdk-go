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

// AddSharedBlockGroupPrimaryStorage adds SharedBlockGroupPrimaryStorage
func (cli *ZSClient) AddSharedBlockGroupPrimaryStorage(ctx context.Context, params param.AddSharedBlockGroupPrimaryStorageParam) (*view.PrimaryStorageInventoryView, error) {
	resp := view.PrimaryStorageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/sharedblockgroup"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySharedBlockGroupPrimaryStorage queries SharedBlockGroupPrimaryStorage list
func (cli *ZSClient) QuerySharedBlockGroupPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp []view.SharedBlockGroupPrimaryStorageInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/sharedblockgroup", params, &resp)
}

func (cli *ZSClient) GetSharedBlockGroupPrimaryStorage(ctx context.Context, uuid string) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	var resp view.SharedBlockGroupPrimaryStorageInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/sharedblockgroup", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSharedBlockGroupPrimaryStorage Pagination
func (cli *ZSClient) PageSharedBlockGroupPrimaryStorage(ctx context.Context, params *param.QueryParam) ([]view.SharedBlockGroupPrimaryStorageInventoryView, int, error) {
	var sharedBlockGroupPrimaryStorages []view.SharedBlockGroupPrimaryStorageInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/sharedblockgroup", params, &sharedBlockGroupPrimaryStorages)
	return sharedBlockGroupPrimaryStorages, total, err
}
