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

// DeleteVolumeSnapshot deletes VolumeSnapshot
func (cli *ZSClient) DeleteVolumeSnapshot(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/volume-snapshots", uuid, string(deleteMode))
}
// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(ctx context.Context, volumeUuid string, params param.CreateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/volumes/%s/volume-snapshots", volumeUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVolumeSnapshot updates VolumeSnapshot
func (cli *ZSClient) UpdateVolumeSnapshot(ctx context.Context, uuid string, params param.UpdateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/volume-snapshots", uuid, "actions", "inventory", map[string]interface{}{
		"updateVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolumeSnapshot queries VolumeSnapshot list
func (cli *ZSClient) QueryVolumeSnapshot(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, error) {
	var resp []view.VolumeSnapshotInventoryView
	return resp, cli.List(ctx, "v1/volume-snapshots", params, &resp)
}

func (cli *ZSClient) GetVolumeSnapshot(ctx context.Context, uuid string) (*view.VolumeSnapshotInventoryView, error) {
	var resp view.VolumeSnapshotInventoryView
	if err := cli.Get(ctx, "v1/volume-snapshots", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeSnapshot Pagination
func (cli *ZSClient) PageVolumeSnapshot(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, int, error) {
	var volumeSnapshots []view.VolumeSnapshotInventoryView
	total, err := cli.Page(ctx, "v1/volume-snapshots", params, &volumeSnapshots)
	return volumeSnapshots, total, err
}
