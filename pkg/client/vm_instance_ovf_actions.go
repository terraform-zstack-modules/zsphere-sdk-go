// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// ParseOvf parses the OVF template information.
func (cli *ZSClient) ParseOvf(params param.ParseOvfParam) (*view.OvfInfo, error) {
	resp := view.OvfInfo{}
	return &resp, cli.PostWithRespKey("v1/ovf/parse", "ovfInfo", params, &resp)
}

// CreateVmInstanceFromOvf imports a virtual machine from an OVF template.
func (cli *ZSClient) CreateVmInstanceFromOvf(params param.CreateVmInstanceFromOvfParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	return &resp, cli.Post("v1/ovf/create-vm-instance", params, &resp)
}
