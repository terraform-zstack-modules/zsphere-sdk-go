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

// DeleteStackTemplate deletes StackTemplate
func (cli *ZSClient) DeleteStackTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cloudformation/template", uuid, string(deleteMode))
}
// QueryStackTemplate queries StackTemplate list
func (cli *ZSClient) QueryStackTemplate(ctx context.Context, params *param.QueryParam) ([]view.StackTemplateInventoryView, error) {
	var resp []view.StackTemplateInventoryView
	return resp, cli.List(ctx, "v1/cloudformation/template", params, &resp)
}

func (cli *ZSClient) GetStackTemplate(ctx context.Context, uuid string) (*view.StackTemplateInventoryView, error) {
	var resp view.StackTemplateInventoryView
	if err := cli.Get(ctx, "v1/cloudformation/template", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageStackTemplate Pagination
func (cli *ZSClient) PageStackTemplate(ctx context.Context, params *param.QueryParam) ([]view.StackTemplateInventoryView, int, error) {
	var stackTemplates []view.StackTemplateInventoryView
	total, err := cli.Page(ctx, "v1/cloudformation/template", params, &stackTemplates)
	return stackTemplates, total, err
}
// UpdateStackTemplate updates StackTemplate
func (cli *ZSClient) UpdateStackTemplate(ctx context.Context, uuid string, params param.UpdateStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/cloudformation/template", uuid, "actions", "inventory", map[string]interface{}{
		"updateStackTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddStackTemplate adds StackTemplate
func (cli *ZSClient) AddStackTemplate(ctx context.Context, params param.AddStackTemplateParam) (*view.StackTemplateInventoryView, error) {
	resp := view.StackTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/template"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
