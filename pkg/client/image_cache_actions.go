// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryImageCache queries ImageCache list
func (cli *ZSClient) QueryImageCache(ctx context.Context, params *param.QueryParam) ([]view.ImageCacheInventoryView, error) {
	var resp []view.ImageCacheInventoryView
	return resp, cli.List(ctx, "v1/primary-storage/imagecache", params, &resp)
}

// PageImageCache Pagination
func (cli *ZSClient) PageImageCache(ctx context.Context, params *param.QueryParam) ([]view.ImageCacheInventoryView, int, error) {
	var imageCaches []view.ImageCacheInventoryView
	total, err := cli.Page(ctx, "v1/primary-storage/imagecache", params, &imageCaches)
	return imageCaches, total, err
}
