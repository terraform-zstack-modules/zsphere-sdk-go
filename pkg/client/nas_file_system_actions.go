// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryNasFileSystem queries NasFileSystem list
func (cli *ZSClient) QueryNasFileSystem(ctx context.Context, params *param.QueryParam) ([]view.NasFileSystemInventoryView, error) {
	var resp []view.NasFileSystemInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/nas", params, &resp)
}

func (cli *ZSClient) GetNasFileSystem(ctx context.Context, uuid string) (*view.NasFileSystemInventoryView, error) {
	var resp view.NasFileSystemInventoryView
	if err := cli.Get(ctx, "v1/primary-storage/nas", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNasFileSystem Pagination
func (cli *ZSClient) PageNasFileSystem(ctx context.Context, params *param.QueryParam) ([]view.NasFileSystemInventoryView, int, error) {
	var nasFileSystems []view.NasFileSystemInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/nas", params, &nasFileSystems)
	return nasFileSystems, total, err
}
// UpdateNasFileSystem updates NasFileSystem
func (cli *ZSClient) UpdateNasFileSystem(ctx context.Context, uuid string, params param.UpdateNasFileSystemParam) (*view.NasFileSystemInventoryView, error) {
	resp := view.NasFileSystemInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/primary-storage/nas", uuid, "actions", "inventory", map[string]interface{}{
		"updateNasFileSystem": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteNasFileSystem deletes NasFileSystem
func (cli *ZSClient) DeleteNasFileSystem(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/primary-storage/nas", uuid, string(deleteMode))
}
