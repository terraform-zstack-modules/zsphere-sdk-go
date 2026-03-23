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

// ChangeInstanceOffering changes InstanceOffering
func (cli *ZSClient) ChangeInstanceOffering(ctx context.Context, vmInstanceUuid string, params param.ChangeInstanceOfferingParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-instances", vmInstanceUuid, "actions", "inventory", map[string]interface{}{
		"changeInstanceOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateInstanceOffering updates InstanceOffering
func (cli *ZSClient) UpdateInstanceOffering(ctx context.Context, uuid string, params param.UpdateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/instance-offerings", uuid, "actions", "inventory", map[string]interface{}{
		"updateInstanceOffering": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateInstanceOffering creates InstanceOffering
func (cli *ZSClient) CreateInstanceOffering(ctx context.Context, params param.CreateInstanceOfferingParam) (*view.InstanceOfferingInventoryView, error) {
	resp := view.InstanceOfferingInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/instance-offerings"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryInstanceOffering queries InstanceOffering list
func (cli *ZSClient) QueryInstanceOffering(ctx context.Context, params *param.QueryParam) ([]view.InstanceOfferingInventoryView, error) {
	var resp []view.InstanceOfferingInventoryView
	return resp, cli.List(ctx, "v1/instance-offerings", params, &resp)
}

func (cli *ZSClient) GetInstanceOffering(ctx context.Context, uuid string) (*view.InstanceOfferingInventoryView, error) {
	var resp view.InstanceOfferingInventoryView
	if err := cli.Get(ctx, "v1/instance-offerings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageInstanceOffering Pagination
func (cli *ZSClient) PageInstanceOffering(ctx context.Context, params *param.QueryParam) ([]view.InstanceOfferingInventoryView, int, error) {
	var instanceOfferings []view.InstanceOfferingInventoryView
	total, err := cli.Page(ctx, "v1/instance-offerings", params, &instanceOfferings)
	return instanceOfferings, total, err
}
// DeleteInstanceOffering deletes InstanceOffering
func (cli *ZSClient) DeleteInstanceOffering(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/instance-offerings", uuid, string(deleteMode))
}
