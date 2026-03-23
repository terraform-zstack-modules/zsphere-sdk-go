// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryVolumeSnapshotTree queries VolumeSnapshotTree list
func (cli *ZSClient) QueryVolumeSnapshotTree(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotTreeInventoryView, error) {
	var resp []view.VolumeSnapshotTreeInventoryView
	return resp, cli.List(ctx, "v1/volume-snapshots/trees", params, &resp)
}

func (cli *ZSClient) GetVolumeSnapshotTree(ctx context.Context, uuid string) (*view.VolumeSnapshotTreeInventoryView, error) {
	var resp view.VolumeSnapshotTreeInventoryView
	if err := cli.Get(ctx, "v1/volume-snapshots/trees", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeSnapshotTree Pagination
func (cli *ZSClient) PageVolumeSnapshotTree(ctx context.Context, params *param.QueryParam) ([]view.VolumeSnapshotTreeInventoryView, int, error) {
	var volumeSnapshotTrees []view.VolumeSnapshotTreeInventoryView
	total, err := cli.Page(ctx, "v1/volume-snapshots/trees", params, &volumeSnapshotTrees)
	return volumeSnapshotTrees, total, err
}
