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

// UpdateThirdpartyPlatform updates ThirdpartyPlatform
func (cli *ZSClient) UpdateThirdpartyPlatform(ctx context.Context, uuid string, params param.UpdateThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	resp := view.ThirdpartyPlatformInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/third-party/platforms", uuid, "actions", "inventory", map[string]interface{}{
		"updateThirdpartyPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddThirdpartyPlatform adds ThirdpartyPlatform
func (cli *ZSClient) AddThirdpartyPlatform(ctx context.Context, params param.AddThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	resp := view.ThirdpartyPlatformInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/third-party/platforms"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryThirdpartyPlatform queries ThirdpartyPlatform list
func (cli *ZSClient) QueryThirdpartyPlatform(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, error) {
	var resp []view.ThirdpartyPlatformInventoryView
	return resp, cli.List(ctx, "v1/zwatch/third-party/platforms", params, &resp)
}

func (cli *ZSClient) GetThirdpartyPlatform(ctx context.Context, uuid string) (*view.ThirdpartyPlatformInventoryView, error) {
	var resp view.ThirdpartyPlatformInventoryView
	if err := cli.Get(ctx, "v1/zwatch/third-party/platforms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageThirdpartyPlatform Pagination
func (cli *ZSClient) PageThirdpartyPlatform(ctx context.Context, params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, int, error) {
	var thirdpartyPlatforms []view.ThirdpartyPlatformInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/third-party/platforms", params, &thirdpartyPlatforms)
	return thirdpartyPlatforms, total, err
}
// DeleteThirdpartyPlatform deletes ThirdpartyPlatform
func (cli *ZSClient) DeleteThirdpartyPlatform(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/third-party/platforms", uuid, string(deleteMode))
}
