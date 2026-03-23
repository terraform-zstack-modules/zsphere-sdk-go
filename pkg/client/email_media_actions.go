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

// QueryEmailMedia queries EmailMedia list
func (cli *ZSClient) QueryEmailMedia(ctx context.Context, params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List(ctx, "v1/media/emails", params, &resp)
}

func (cli *ZSClient) GetEmailMedia(ctx context.Context, uuid string) (*view.MediaInventoryView, error) {
	var resp view.MediaInventoryView
	if err := cli.Get(ctx, "v1/media/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEmailMedia Pagination
func (cli *ZSClient) PageEmailMedia(ctx context.Context, params *param.QueryParam) ([]view.MediaInventoryView, int, error) {
	var emailMedias []view.MediaInventoryView
	total, err := cli.Page(ctx, "v1/media/emails", params, &emailMedias)
	return emailMedias, total, err
}
// CreateEmailMedia creates EmailMedia
func (cli *ZSClient) CreateEmailMedia(ctx context.Context, params param.CreateEmailMediaParam) (*view.MediaInventoryView, error) {
	resp := view.MediaInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/media/emails"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEmailMedia updates EmailMedia
func (cli *ZSClient) UpdateEmailMedia(ctx context.Context, uuid string, params param.UpdateEmailMediaParam) (*view.EmailMediaInventoryView, error) {
	resp := view.EmailMediaInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/media/emails", uuid, "", "inventory", map[string]interface{}{
		"updateEmailMedia": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
