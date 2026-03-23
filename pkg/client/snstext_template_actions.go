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

// QuerySNSTextTemplate queries SNSTextTemplate list
func (cli *ZSClient) QuerySNSTextTemplate(ctx context.Context, params *param.QueryParam) ([]view.SNSTextTemplateInventoryView, error) {
	var resp []view.SNSTextTemplateInventoryView
	return resp, cli.List(ctx, "v1/zwatch/alarms/sns/text-templates", params, &resp)
}

func (cli *ZSClient) GetSNSTextTemplate(ctx context.Context, uuid string) (*view.SNSTextTemplateInventoryView, error) {
	var resp view.SNSTextTemplateInventoryView
	if err := cli.Get(ctx, "v1/zwatch/alarms/sns/text-templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSTextTemplate Pagination
func (cli *ZSClient) PageSNSTextTemplate(ctx context.Context, params *param.QueryParam) ([]view.SNSTextTemplateInventoryView, int, error) {
	var sNSTextTemplates []view.SNSTextTemplateInventoryView
	total, err := cli.Page(ctx, "v1/zwatch/alarms/sns/text-templates", params, &sNSTextTemplates)
	return sNSTextTemplates, total, err
}
// UpdateSNSTextTemplate updates SNSTextTemplate
func (cli *ZSClient) UpdateSNSTextTemplate(ctx context.Context, uuid string, params param.UpdateSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/zwatch/alarms/sns/text-templates", uuid, "actions", "inventory", map[string]interface{}{
		"updateSNSTextTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSTextTemplate creates SNSTextTemplate
func (cli *ZSClient) CreateSNSTextTemplate(ctx context.Context, params param.CreateSNSTextTemplateParam) (*view.SNSTextTemplateInventoryView, error) {
	resp := view.SNSTextTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/zwatch/alarms/sns/text-templates"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSNSTextTemplate deletes SNSTextTemplate
func (cli *ZSClient) DeleteSNSTextTemplate(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/zwatch/alarms/sns/text-templates", uuid, string(deleteMode))
}
