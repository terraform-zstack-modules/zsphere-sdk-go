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

// DeletePreconfigurationTemplate deletes PreconfigurationTemplate
func (cli *ZSClient) DeletePreconfigurationTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal/preconfigurations", uuid, string(deleteMode))
}
// UpdatePreconfigurationTemplate updates PreconfigurationTemplate
func (cli *ZSClient) UpdatePreconfigurationTemplate(ctx context.Context, uuid string, params param.UpdatePreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/preconfigurations", uuid, "actions", "inventory", map[string]interface{}{
		"updatePreconfigurationTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryPreconfigurationTemplate queries PreconfigurationTemplate list
func (cli *ZSClient) QueryPreconfigurationTemplate(ctx context.Context, params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, error) {
	var resp []view.PreconfigurationTemplateInventoryView
	return resp, cli.List(ctx, "v1/baremetal/preconfigurations", params, &resp)
}

func (cli *ZSClient) GetPreconfigurationTemplate(ctx context.Context, uuid string) (*view.PreconfigurationTemplateInventoryView, error) {
	var resp view.PreconfigurationTemplateInventoryView
	if err := cli.Get(ctx, "v1/baremetal/preconfigurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePreconfigurationTemplate Pagination
func (cli *ZSClient) PagePreconfigurationTemplate(ctx context.Context, params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, int, error) {
	var preconfigurationTemplates []view.PreconfigurationTemplateInventoryView
	total, err := cli.Page(ctx, "v1/baremetal/preconfigurations", params, &preconfigurationTemplates)
	return preconfigurationTemplates, total, err
}
// AddPreconfigurationTemplate adds PreconfigurationTemplate
func (cli *ZSClient) AddPreconfigurationTemplate(ctx context.Context, params param.AddPreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/preconfigurations"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
