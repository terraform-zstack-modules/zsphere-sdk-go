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

// DeleteResourceStack deletes ResourceStack
func (cli *ZSClient) DeleteResourceStack(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/cloudformation/stack", uuid, string(deleteMode))
}
// CreateResourceStack creates ResourceStack
func (cli *ZSClient) CreateResourceStack(ctx context.Context, params param.CreateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/cloudformation/stack"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryResourceStack queries ResourceStack list
func (cli *ZSClient) QueryResourceStack(ctx context.Context, params *param.QueryParam) ([]view.ResourceStackInventoryView, error) {
	var resp []view.ResourceStackInventoryView
	return resp, cli.List(ctx, "v1/cloudformation/stack", params, &resp)
}

func (cli *ZSClient) GetResourceStack(ctx context.Context, uuid string) (*view.ResourceStackInventoryView, error) {
	var resp view.ResourceStackInventoryView
	if err := cli.Get(ctx, "v1/cloudformation/stack", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourceStack Pagination
func (cli *ZSClient) PageResourceStack(ctx context.Context, params *param.QueryParam) ([]view.ResourceStackInventoryView, int, error) {
	var resourceStacks []view.ResourceStackInventoryView
	total, err := cli.Page(ctx, "v1/cloudformation/stack", params, &resourceStacks)
	return resourceStacks, total, err
}
// UpdateResourceStack updates ResourceStack
func (cli *ZSClient) UpdateResourceStack(ctx context.Context, uuid string, params param.UpdateResourceStackParam) (*view.ResourceStackInventoryView, error) {
	resp := view.ResourceStackInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/cloudformation/stack", uuid, "actions", "inventory", map[string]interface{}{
		"updateResourceStack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
