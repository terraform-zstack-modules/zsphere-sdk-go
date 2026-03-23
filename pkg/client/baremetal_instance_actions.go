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

// UpdateBaremetalInstance updates BaremetalInstance
func (cli *ZSClient) UpdateBaremetalInstance(ctx context.Context, uuid string, params param.UpdateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/instances", uuid, "actions", "inventory", map[string]interface{}{
		"updateBaremetalInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RebootBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RebootBaremetalInstance(ctx context.Context, uuid string, params param.RebootBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/instances", uuid, "actions", "inventory", map[string]interface{}{
		"rebootBaremetalInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopBaremetalInstance stops BaremetalInstance
func (cli *ZSClient) StopBaremetalInstance(ctx context.Context, uuid string, params param.StopBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/instances", uuid, "actions", "inventory", map[string]interface{}{
		"stopBaremetalInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryBaremetalInstance queries BaremetalInstance list
func (cli *ZSClient) QueryBaremetalInstance(ctx context.Context, params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, error) {
	var resp []view.BaremetalInstanceInventoryView
	return resp, cli.List(ctx, "v1/baremetal/instances", params, &resp)
}

func (cli *ZSClient) GetBaremetalInstance(ctx context.Context, uuid string) (*view.BaremetalInstanceInventoryView, error) {
	var resp view.BaremetalInstanceInventoryView
	if err := cli.Get(ctx, "v1/baremetal/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBaremetalInstance Pagination
func (cli *ZSClient) PageBaremetalInstance(ctx context.Context, params *param.QueryParam) ([]view.BaremetalInstanceInventoryView, int, error) {
	var baremetalInstances []view.BaremetalInstanceInventoryView
	total, err := cli.Page(ctx, "v1/baremetal/instances", params, &baremetalInstances)
	return baremetalInstances, total, err
}
// StartBaremetalInstance starts BaremetalInstance
func (cli *ZSClient) StartBaremetalInstance(ctx context.Context, uuid string, params param.StartBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/instances", uuid, "actions", "inventory", map[string]interface{}{
		"startBaremetalInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateBaremetalInstance creates BaremetalInstance
func (cli *ZSClient) CreateBaremetalInstance(ctx context.Context, params param.CreateBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/baremetal/instances"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DestroyBaremetalInstance destroys BaremetalInstance
func (cli *ZSClient) DestroyBaremetalInstance(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/baremetal/instances", uuid, string(deleteMode))
}
// ExpungeBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) ExpungeBaremetalInstance(ctx context.Context, uuid string, params param.ExpungeBaremetalInstanceParam) error {
	return cli.Put(ctx, "v1/baremetal/instances", uuid, map[string]interface{}{
		"expungeBaremetalInstance": params.Params,
	}, nil)
}
// RecoverBaremetalInstance operates on BaremetalInstance
func (cli *ZSClient) RecoverBaremetalInstance(ctx context.Context, uuid string, params param.RecoverBaremetalInstanceParam) (*view.BaremetalInstanceInventoryView, error) {
	resp := view.BaremetalInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/baremetal/instances", uuid, "actions", "inventory", map[string]interface{}{
		"recoverBaremetalInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
