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

// CreateVolumeSnapshotGroup creates VolumeSnapshotGroup
func (cli *ZSClient) CreateVolumeSnapshotGroup(ctx context.Context, params param.CreateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	resp := view.VolumeSnapshotGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volume-snapshots/group"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVolumeSnapshotGroup updates VolumeSnapshotGroup
func (cli *ZSClient) UpdateVolumeSnapshotGroup(ctx context.Context, uuid string, params param.UpdateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	resp := view.VolumeSnapshotGroupInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots/group", uuid, "actions", "inventory", map[string]interface{}{
		"updateVolumeSnapshotGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volume-snapshots/group", uuid, string(deleteMode))
}
// QueryVolumeSnapshotGroup queries VolumeSnapshotGroup list
func (cli *ZSClient) QueryVolumeSnapshotGroup(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotGroupInventoryView, error) {
	var resp []view.VolumeSnapshotGroupInventoryView
	return resp, cli.List(ctx, "v1/volume-snapshots/group", params, &resp)
}

func (cli *ZSClient) GetVolumeSnapshotGroup(ctx context.Context, uuid string) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.VolumeSnapshotGroupInventoryView
	if err := cli.Get(ctx, "v1/volume-snapshots/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeSnapshotGroup Pagination
func (cli *ZSClient) PageVolumeSnapshotGroup(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotGroupInventoryView, int, error) {
	var volumeSnapshotGroups []view.VolumeSnapshotGroupInventoryView
	total, err := cli.Page(ctx, "v1/volume-snapshots/group", params, &volumeSnapshotGroups)
	return volumeSnapshotGroups, total, err
}
