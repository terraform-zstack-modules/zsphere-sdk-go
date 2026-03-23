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

// CreateEip creates Eip
func (cli *ZSClient) CreateEip(ctx context.Context, params param.CreateEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/eips"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AttachEip operates on Eip
func (cli *ZSClient) AttachEip(ctx context.Context, eipUuid, vmNicUuid string, params param.AttachEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/eips/%s/vm-instances/nics/%s", eipUuid, vmNicUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateEip updates Eip
func (cli *ZSClient) UpdateEip(ctx context.Context, uuid string, params param.UpdateEipParam) (*view.EipInventoryView, error) {
	resp := view.EipInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/eips", uuid, "actions", "inventory", map[string]interface{}{
		"updateEip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteEip deletes Eip
func (cli *ZSClient) DeleteEip(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/eips", uuid, string(deleteMode))
}
// QueryEip queries Eip list
func (cli *ZSClient) QueryEip(ctx context.Context, params *param.QueryParam) ([]view.EipInventoryView, error) {
	var resp []view.EipInventoryView
	return resp, cli.List(ctx, "v1/eips", params, &resp)
}

func (cli *ZSClient) GetEip(ctx context.Context, uuid string) (*view.EipInventoryView, error) {
	var resp view.EipInventoryView
	if err := cli.Get(ctx, "v1/eips", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEip Pagination
func (cli *ZSClient) PageEip(ctx context.Context, params *param.QueryParam) ([]view.EipInventoryView, int, error) {
	var eips []view.EipInventoryView
	total, err := cli.Page(ctx, "v1/eips", params, &eips)
	return eips, total, err
}
// DetachEip operates on Eip
func (cli *ZSClient) DetachEip(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/eips", uuid, string(deleteMode))
}
