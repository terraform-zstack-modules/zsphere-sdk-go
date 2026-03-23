// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateNasMountTarget updates NasMountTarget
func (cli *ZSClient) UpdateNasMountTarget(ctx context.Context, uuid string, params param.UpdateNasMountTargetParam) (*view.NasMountTargetInventoryView, error) {
	resp := view.NasMountTargetInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/nas/mount", uuid, "actions", "inventory", map[string]interface{}{
		"updateNasMountTarget": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNasMountTarget deletes NasMountTarget
func (cli *ZSClient) DeleteNasMountTarget(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage/nas/mount", uuid, string(deleteMode))
}
// QueryNasMountTarget queries NasMountTarget list
func (cli *ZSClient) QueryNasMountTarget(ctx context.Context, params *param.QueryParam) ([]view.NasMountTargetInventoryView, error) {
	var resp []view.NasMountTargetInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/nas/mount", params, &resp)
}

func (cli *ZSClient) GetNasMountTarget(ctx context.Context, uuid string) (*view.NasMountTargetInventoryView, error) {
	var resp view.NasMountTargetInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/nas/mount", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNasMountTarget Pagination
func (cli *ZSClient) PageNasMountTarget(ctx context.Context, params *param.QueryParam) ([]view.NasMountTargetInventoryView, int, error) {
	var nasMountTargets []view.NasMountTargetInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/nas/mount", params, &nasMountTargets)
	return nasMountTargets, total, err
}
