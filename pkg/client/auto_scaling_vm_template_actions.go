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

// UpdateAutoScalingVmTemplate updates AutoScalingVmTemplate
func (cli *ZSClient) UpdateAutoScalingVmTemplate(ctx context.Context, uuid string, params param.UpdateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/autoscaling/vmtemplate", uuid, "actions", "inventory", map[string]interface{}{
		"updateAutoScalingVmTemplate": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateAutoScalingVmTemplate creates AutoScalingVmTemplate
func (cli *ZSClient) CreateAutoScalingVmTemplate(ctx context.Context, params param.CreateAutoScalingVmTemplateParam) (*view.AutoScalingTemplateInventoryView, error) {
	resp := view.AutoScalingTemplateInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/autoscaling/vmtemplate"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAutoScalingVmTemplate queries AutoScalingVmTemplate list
func (cli *ZSClient) QueryAutoScalingVmTemplate(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, error) {
	var resp []view.AutoScalingVmTemplateInventoryView
	return resp, cli.List(ctx, "v1/autoscaling/vmtemplate", params, &resp)
}

func (cli *ZSClient) GetAutoScalingVmTemplate(ctx context.Context, uuid string) (*view.AutoScalingVmTemplateInventoryView, error) {
	var resp view.AutoScalingVmTemplateInventoryView
	if err := cli.Get(ctx, "v1/autoscaling/vmtemplate", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAutoScalingVmTemplate Pagination
func (cli *ZSClient) PageAutoScalingVmTemplate(ctx context.Context, params *param.QueryParam) ([]view.AutoScalingVmTemplateInventoryView, int, error) {
	var autoScalingVmTemplates []view.AutoScalingVmTemplateInventoryView
	total, err := cli.Page(ctx, "v1/autoscaling/vmtemplate", params, &autoScalingVmTemplates)
	return autoScalingVmTemplates, total, err
}
