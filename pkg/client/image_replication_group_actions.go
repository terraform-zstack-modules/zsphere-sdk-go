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

// DeleteImageReplicationGroup deletes ImageReplicationGroup
func (cli *ZSClient) DeleteImageReplicationGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/image-replication-groups", uuid, string(deleteMode))
}
// QueryImageReplicationGroup queries ImageReplicationGroup list
func (cli *ZSClient) QueryImageReplicationGroup(ctx context.Context, params *param.QueryParam) ([]view.ImageReplicationGroupInventoryView, error) {
	var resp []view.ImageReplicationGroupInventoryView
	return resp, cli.List(ctx, "v1/image-replication-groups", params, &resp)
}

func (cli *ZSClient) GetImageReplicationGroup(ctx context.Context, uuid string) (*view.ImageReplicationGroupInventoryView, error) {
	var resp view.ImageReplicationGroupInventoryView
	if err := cli.Get(ctx, "v1/image-replication-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageReplicationGroup Pagination
func (cli *ZSClient) PageImageReplicationGroup(ctx context.Context, params *param.QueryParam) ([]view.ImageReplicationGroupInventoryView, int, error) {
	var imageReplicationGroups []view.ImageReplicationGroupInventoryView
	total, err := cli.Page(ctx, "v1/image-replication-groups", params, &imageReplicationGroups)
	return imageReplicationGroups, total, err
}
// CreateImageReplicationGroup creates ImageReplicationGroup
func (cli *ZSClient) CreateImageReplicationGroup(ctx context.Context, params param.CreateImageReplicationGroupParam) (*view.ImageReplicationGroupInventoryView, error) {
	resp := view.ImageReplicationGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/image-replication-groups"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
