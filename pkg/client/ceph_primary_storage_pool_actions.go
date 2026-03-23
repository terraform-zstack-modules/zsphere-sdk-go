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

// AddCephPrimaryStoragePool adds CephPrimaryStoragePool
func (cli *ZSClient) AddCephPrimaryStoragePool(ctx context.Context, primaryStorageUuid string, params param.AddCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	resp := view.CephPrimaryStoragePoolInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/primary-storage/ceph/%s/pools", primaryStorageUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCephPrimaryStoragePool queries CephPrimaryStoragePool list
func (cli *ZSClient) QueryCephPrimaryStoragePool(ctx context.Context, params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, error) {
	var resp []view.CephPrimaryStoragePoolInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/ceph/pools", params, &resp)
}

func (cli *ZSClient) GetCephPrimaryStoragePool(ctx context.Context, uuid string) (*view.CephPrimaryStoragePoolInventoryView, error) {
	var resp view.CephPrimaryStoragePoolInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/ceph/pools", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCephPrimaryStoragePool Pagination
func (cli *ZSClient) PageCephPrimaryStoragePool(ctx context.Context, params *param.QueryParam) ([]view.CephPrimaryStoragePoolInventoryView, int, error) {
	var cephPrimaryStoragePools []view.CephPrimaryStoragePoolInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/ceph/pools", params, &cephPrimaryStoragePools)
	return cephPrimaryStoragePools, total, err
}
// DeleteCephPrimaryStoragePool deletes CephPrimaryStoragePool
func (cli *ZSClient) DeleteCephPrimaryStoragePool(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage/ceph/pools", uuid, string(deleteMode))
}
// UpdateCephPrimaryStoragePool updates CephPrimaryStoragePool
func (cli *ZSClient) UpdateCephPrimaryStoragePool(ctx context.Context, uuid string, params param.UpdateCephPrimaryStoragePoolParam) (*view.CephPrimaryStoragePoolInventoryView, error) {
	resp := view.CephPrimaryStoragePoolInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/ceph/pools", uuid, "actions", "inventory", map[string]interface{}{
		"updateCephPrimaryStoragePool": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
