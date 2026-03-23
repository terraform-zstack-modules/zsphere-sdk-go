// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryHostPhysicalCpu queries HostPhysicalCpu list
func (cli *ZSClient) QueryHostPhysicalCpu(ctx context.Context, params *param.QueryParam) ([]view.HostPhysicalCpuInventoryView, error) {
	var resp []view.HostPhysicalCpuInventoryView
	return resp, cli.List(ctx, "v1/hosts/physical-cpu", params, &resp)
}

func (cli *ZSClient) GetHostPhysicalCpu(ctx context.Context, uuid string) (*view.HostPhysicalCpuInventoryView, error) {
	var resp view.HostPhysicalCpuInventoryView
	if err := cli.Get(ctx, "v1/hosts/physical-cpu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHostPhysicalCpu Pagination
func (cli *ZSClient) PageHostPhysicalCpu(ctx context.Context, params *param.QueryParam) ([]view.HostPhysicalCpuInventoryView, int, error) {
	var hostPhysicalCpus []view.HostPhysicalCpuInventoryView
	total, err := cli.Page(ctx, "v1/hosts/physical-cpu", params, &hostPhysicalCpus)
	return hostPhysicalCpus, total, err
}
