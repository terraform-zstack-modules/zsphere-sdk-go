// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateTemplateConfig updates TemplateConfig
func (cli *ZSClient) UpdateTemplateConfig(ctx context.Context, templateUuid string, params param.UpdateTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/template-configurations", templateUuid, "actions", "inventory", map[string]interface{}{
		"updateTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ApplyTemplateConfig operates on TemplateConfig
func (cli *ZSClient) ApplyTemplateConfig(ctx context.Context, templateUuid string, params param.ApplyTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/template-configurations", templateUuid, "actions", "", map[string]interface{}{
		"applyTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryTemplateConfig queries TemplateConfig list
func (cli *ZSClient) QueryTemplateConfig(ctx context.Context, params *param.QueryParam) ([]view.TemplateConfigInventoryView, error) {
	var resp []view.TemplateConfigInventoryView
	return resp, cli.List(ctx, "v1/template-configurations/configs", params, &resp)
}

func (cli *ZSClient) GetTemplateConfig(ctx context.Context, uuid string) (*view.TemplateConfigInventoryView, error) {
	var resp view.TemplateConfigInventoryView
	if err := cli.Get(ctx, "v1/template-configurations/configs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTemplateConfig Pagination
func (cli *ZSClient) PageTemplateConfig(ctx context.Context, params *param.QueryParam) ([]view.TemplateConfigInventoryView, int, error) {
	var templateConfigs []view.TemplateConfigInventoryView
	total, err := cli.Page(ctx, "v1/template-configurations/configs", params, &templateConfigs)
	return templateConfigs, total, err
}
// RevertTemplateConfig operates on TemplateConfig
func (cli *ZSClient) RevertTemplateConfig(ctx context.Context, templateUuid string, params param.RevertTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/template-configurations", templateUuid, "actions", "", map[string]interface{}{
		"revertTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResetTemplateConfig operates on TemplateConfig
func (cli *ZSClient) ResetTemplateConfig(ctx context.Context, templateUuid string, params param.ResetTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/template-configurations", templateUuid, "actions", "", map[string]interface{}{
		"resetTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
