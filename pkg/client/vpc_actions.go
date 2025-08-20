// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryVpcRouter queries the list of VPC routers.
func (cli *ZSClient) QueryVpcRouter(params param.QueryParam) ([]view.VpcRouterVmInventoryView, error) {
	resp := make([]view.VpcRouterVmInventoryView, 0)
	return resp, cli.List("v1/vpc/virtual-routers", &params, &resp)
}

// GetVpcRouter retrieves details of a specific VPC router by its UUID.
func (cli *ZSClient) GetVpcRouter(uuid string) (view.VpcRouterVmInventoryView, error) {
	var resp view.VpcRouterVmInventoryView
	return resp, cli.Get("v1/vpc/virtual-routers", uuid, nil, &resp)
}
