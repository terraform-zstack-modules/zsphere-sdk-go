// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(ctx context.Context, uuid string, params param.UpdateImagePackageParam) (*view.ImagePackageInventoryView, error) {
	resp := view.ImagePackageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/image-packages", uuid, "", "inventory", map[string]interface{}{
		"updateImagePackage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImagePackage queries ImagePackage list
func (cli *ZSClient) QueryImagePackage(ctx context.Context, params *param.QueryParam) ([]view.ImagePackageInventoryView, error) {
	var resp []view.ImagePackageInventoryView
	return resp, cli.List(ctx, "v1/image-packages", params, &resp)
}

func (cli *ZSClient) GetImagePackage(ctx context.Context, uuid string) (*view.ImagePackageInventoryView, error) {
	var resp view.ImagePackageInventoryView
	if err := cli.Get(ctx, "v1/image-packages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImagePackage Pagination
func (cli *ZSClient) PageImagePackage(ctx context.Context, params *param.QueryParam) ([]view.ImagePackageInventoryView, int, error) {
	var imagePackages []view.ImagePackageInventoryView
	total, err := cli.Page(ctx, "v1/image-packages", params, &imagePackages)
	return imagePackages, total, err
}
// DeleteImagePackage deletes ImagePackage
func (cli *ZSClient) DeleteImagePackage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/image-packages", uuid, string(deleteMode))
}
