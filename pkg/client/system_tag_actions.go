// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QuerySystemTag queries SystemTag list
func (cli *ZSClient) QuerySystemTag(ctx context.Context, params *param.QueryParam) ([]view.SystemTagInventoryView, error) {
	var resp []view.SystemTagInventoryView
	return resp, cli.List(ctx, "v1/system-tags", params, &resp)
}

func (cli *ZSClient) GetSystemTag(ctx context.Context, uuid string) (*view.SystemTagInventoryView, error) {
	var resp view.SystemTagInventoryView
	if err := cli.Get(ctx, "v1/system-tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSystemTag Pagination
func (cli *ZSClient) PageSystemTag(ctx context.Context, params *param.QueryParam) ([]view.SystemTagInventoryView, int, error) {
	var systemTags []view.SystemTagInventoryView
	total, err := cli.Page(ctx, "v1/system-tags", params, &systemTags)
	return systemTags, total, err
}
// UpdateSystemTag updates SystemTag
func (cli *ZSClient) UpdateSystemTag(ctx context.Context, uuid string, params param.UpdateSystemTagParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/system-tags", uuid, "actions", "inventory", map[string]interface{}{
		"updateSystemTag": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSystemTag creates SystemTag
func (cli *ZSClient) CreateSystemTag(ctx context.Context) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.Post(ctx, "v1/system-tags", map[string]interface{}{}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
