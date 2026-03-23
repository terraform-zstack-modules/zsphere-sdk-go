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

// QueryResourceAttributeKey queries ResourceAttributeKey list
func (cli *ZSClient) QueryResourceAttributeKey(ctx context.Context, params *param.QueryParam) ([]view.ResourceAttributeKeyInventoryView, error) {
	var resp []view.ResourceAttributeKeyInventoryView
	return resp, cli.List(ctx, "v1/resource-attributes/keys", params, &resp)
}

func (cli *ZSClient) GetResourceAttributeKey(ctx context.Context, uuid string) (*view.ResourceAttributeKeyInventoryView, error) {
	var resp view.ResourceAttributeKeyInventoryView
	if err := cli.Get(ctx, "v1/resource-attributes/keys", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageResourceAttributeKey Pagination
func (cli *ZSClient) PageResourceAttributeKey(ctx context.Context, params *param.QueryParam) ([]view.ResourceAttributeKeyInventoryView, int, error) {
	var resourceAttributeKeys []view.ResourceAttributeKeyInventoryView
	total, err := cli.Page(ctx, "v1/resource-attributes/keys", params, &resourceAttributeKeys)
	return resourceAttributeKeys, total, err
}
// DeleteResourceAttributeKey deletes ResourceAttributeKey
func (cli *ZSClient) DeleteResourceAttributeKey(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/resource-attributes", uuid, string(deleteMode))
}
// UpdateResourceAttributeKey updates ResourceAttributeKey
func (cli *ZSClient) UpdateResourceAttributeKey(ctx context.Context, uuid string, params param.UpdateResourceAttributeKeyParam) (*view.ResourceAttributeKeyInventoryView, error) {
	resp := view.ResourceAttributeKeyInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/resource-attributes", uuid, "actions", "inventory", map[string]interface{}{
		"updateResourceAttributeKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateResourceAttributeKey creates ResourceAttributeKey
func (cli *ZSClient) CreateResourceAttributeKey(ctx context.Context, params param.CreateResourceAttributeKeyParam) (*view.ResourceAttributeKeyInventoryView, error) {
	resp := view.ResourceAttributeKeyInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/resource-attributes"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
