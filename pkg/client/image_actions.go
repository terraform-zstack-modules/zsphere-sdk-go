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

// UpdateImage updates Image
func (cli *ZSClient) UpdateImage(ctx context.Context, uuid string, params param.UpdateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", uuid, "actions", "inventory", map[string]interface{}{
		"updateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddImage adds Image
func (cli *ZSClient) AddImage(ctx context.Context, params param.AddImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/images"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddImageAsync Async
func (cli *ZSClient) AddImageAsync(ctx context.Context, params param.AddImageParam) (string, error) {

	resource := "images"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(ctx, resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// QueryImage queries Image list
func (cli *ZSClient) QueryImage(ctx context.Context, params *param.QueryParam) ([]view.ImageInventoryView, error) {
	var resp []view.ImageInventoryView
	return resp, cli.List(ctx, "v1/images", params, &resp)
}

func (cli *ZSClient) GetImage(ctx context.Context, uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Get(ctx, "v1/images", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImage Pagination
func (cli *ZSClient) PageImage(ctx context.Context, params *param.QueryParam) ([]view.ImageInventoryView, int, error) {
	var images []view.ImageInventoryView
	total, err := cli.Page(ctx, "v1/images", params, &images)
	return images, total, err
}
// SyncImage operates on Image
func (cli *ZSClient) SyncImage(ctx context.Context, imageStoreUuid string, params param.SyncImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/backup-storage/image-store", imageStoreUuid, "actions", "", map[string]interface{}{
		"syncImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(ctx context.Context, uuid string, params param.ExpungeImageParam) error {
	return cli.Put(ctx, "v1/images", uuid, map[string]interface{}{
		"expungeImage": params.Params,
	}, nil)
}
// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(ctx context.Context, imageUuid string, params param.RecoverImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/images", imageUuid, "actions", "inventory", map[string]interface{}{
		"recoverImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteImage deletes Image
func (cli *ZSClient) DeleteImage(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/images", uuid, string(deleteMode))
}
