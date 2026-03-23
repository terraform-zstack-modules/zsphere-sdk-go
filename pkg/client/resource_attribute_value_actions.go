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

// DeleteResourceAttributeValue deletes ResourceAttributeValue
func (cli *ZSClient) DeleteResourceAttributeValue(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/resource-attributes", uuid, string(deleteMode))
}
// CreateResourceAttributeValue creates ResourceAttributeValue
func (cli *ZSClient) CreateResourceAttributeValue(ctx context.Context, keyUuid string, params param.CreateResourceAttributeValueParam) (*view.CreateResourceAttributeResultView, error) {
	resp := view.CreateResourceAttributeResultView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/resource-attributes/%s/resources", keyUuid), "inventories", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryResourceAttributeValue queries ResourceAttributeValue list
func (cli *ZSClient) QueryResourceAttributeValue(ctx context.Context, params *param.QueryParam) ([]view.ResourceAttributeValueInventoryView, error) {
	var resp []view.ResourceAttributeValueInventoryView
	return resp, cli.List(ctx, "v1/resource-attributes", params, &resp)
}

// PageResourceAttributeValue Pagination
func (cli *ZSClient) PageResourceAttributeValue(ctx context.Context, params *param.QueryParam) ([]view.ResourceAttributeValueInventoryView, int, error) {
	var resourceAttributeValues []view.ResourceAttributeValueInventoryView
	total, err := cli.Page(ctx, "v1/resource-attributes", params, &resourceAttributeValues)
	return resourceAttributeValues, total, err
}
