// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// DeleteMedia deletes Media
func (cli *ZSClient) DeleteMedia(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/media", uuid, string(deleteMode))
}
// QueryMedia queries Media list
func (cli *ZSClient) QueryMedia(ctx context.Context, params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List(ctx, "v1/media", params, &resp)
}

func (cli *ZSClient) GetMedia(ctx context.Context, uuid string) (*view.MediaInventoryView, error) {
	var resp view.MediaInventoryView
	if err := cli.Get(ctx, "v1/media", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMedia Pagination
func (cli *ZSClient) PageMedia(ctx context.Context, params *param.QueryParam) ([]view.MediaInventoryView, int, error) {
	var medias []view.MediaInventoryView
	total, err := cli.Page(ctx, "v1/media", params, &medias)
	return medias, total, err
}
