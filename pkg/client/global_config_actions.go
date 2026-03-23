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

// UpdateGlobalConfig updates GlobalConfig
func (cli *ZSClient) UpdateGlobalConfig(ctx context.Context, category string, name string, params param.UpdateGlobalConfigParam) (*view.GlobalConfigInventoryView, error) {
	resp := view.GlobalConfigInventoryView{}
	err := cli.PutWithSpec(ctx, "v1/global-configurations", category, fmt.Sprintf("%s/actions", name), "inventory", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResetGlobalConfig operates on GlobalConfig
func (cli *ZSClient) ResetGlobalConfig(ctx context.Context) (*view.GlobalConfigInventoryView, error) {
	resp := view.GlobalConfigInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/global-configurations/actions", "", "", map[string]interface{}{
		"resetGlobalConfig": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryGlobalConfig queries GlobalConfig list
func (cli *ZSClient) QueryGlobalConfig(ctx context.Context, params *param.QueryParam) ([]view.GlobalConfigInventoryView, error) {
	var resp []view.GlobalConfigInventoryView
	return resp, cli.List(ctx, "v1/global-configurations", params, &resp)
}

// PageGlobalConfig Pagination
func (cli *ZSClient) PageGlobalConfig(ctx context.Context, params *param.QueryParam) ([]view.GlobalConfigInventoryView, int, error) {
	var globalConfigs []view.GlobalConfigInventoryView
	total, err := cli.Page(ctx, "v1/global-configurations", params, &globalConfigs)
	return globalConfigs, total, err
}
