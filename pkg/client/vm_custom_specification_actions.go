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

// QueryVmCustomSpecification queries VmCustomSpecification list
func (cli *ZSClient) QueryVmCustomSpecification(ctx context.Context, params *param.QueryParam) ([]view.VmCustomSpecificationInventoryView, error) {
	var resp []view.VmCustomSpecificationInventoryView
	return resp, cli.List(ctx, "v1/vm-custom-specifications", params, &resp)
}

func (cli *ZSClient) GetVmCustomSpecification(ctx context.Context, uuid string) (*view.VmCustomSpecificationInventoryView, error) {
	var resp view.VmCustomSpecificationInventoryView
	if err := cli.Get(ctx, "v1/vm-custom-specifications", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmCustomSpecification Pagination
func (cli *ZSClient) PageVmCustomSpecification(ctx context.Context, params *param.QueryParam) ([]view.VmCustomSpecificationInventoryView, int, error) {
	var vmCustomSpecifications []view.VmCustomSpecificationInventoryView
	total, err := cli.Page(ctx, "v1/vm-custom-specifications", params, &vmCustomSpecifications)
	return vmCustomSpecifications, total, err
}
// DeleteVmCustomSpecification deletes VmCustomSpecification
func (cli *ZSClient) DeleteVmCustomSpecification(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vm-custom-specifications", uuid, string(deleteMode))
}
// UpdateVmCustomSpecification updates VmCustomSpecification
func (cli *ZSClient) UpdateVmCustomSpecification(ctx context.Context, uuid string, params param.UpdateVmCustomSpecificationParam) (*view.VmCustomSpecificationInventoryView, error) {
	resp := view.VmCustomSpecificationInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/vm-custom-specifications", uuid, "actions", "inventory", map[string]interface{}{
		"updateVmCustomSpecification": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVmCustomSpecification creates VmCustomSpecification
func (cli *ZSClient) CreateVmCustomSpecification(ctx context.Context, params param.CreateVmCustomSpecificationParam) (*view.VmCustomSpecificationInventoryView, error) {
	resp := view.VmCustomSpecificationInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vm-custom-specifications"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
