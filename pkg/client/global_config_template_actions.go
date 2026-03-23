// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryGlobalConfigTemplate queries GlobalConfigTemplate list
func (cli *ZSClient) QueryGlobalConfigTemplate(ctx context.Context, params *param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, error) {
	var resp []view.GlobalConfigTemplateInventoryView
	return resp, cli.List(ctx, "v1/template-configurations/templates", params, &resp)
}

func (cli *ZSClient) GetGlobalConfigTemplate(ctx context.Context, uuid string) (*view.GlobalConfigTemplateInventoryView, error) {
	var resp view.GlobalConfigTemplateInventoryView
	if err := cli.Get(ctx, "v1/template-configurations/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGlobalConfigTemplate Pagination
func (cli *ZSClient) PageGlobalConfigTemplate(ctx context.Context, params *param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, int, error) {
	var globalConfigTemplates []view.GlobalConfigTemplateInventoryView
	total, err := cli.Page(ctx, "v1/template-configurations/templates", params, &globalConfigTemplates)
	return globalConfigTemplates, total, err
}
