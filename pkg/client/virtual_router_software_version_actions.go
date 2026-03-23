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

// GetVirtualRouterSoftwareVersion gets VirtualRouterSoftwareVersion by uuid
func (cli *ZSClient) GetVirtualRouterSoftwareVersion(ctx context.Context) (*view.VirtualRouterSoftwareVersionInventoryView, error) {
	var resp view.VirtualRouterSoftwareVersionInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/vpc/virtual-routers/softwareversion", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVirtualRouterSoftwareVersion updates VirtualRouterSoftwareVersion
func (cli *ZSClient) UpdateVirtualRouterSoftwareVersion(ctx context.Context, uuid string, params param.UpdateVirtualRouterSoftwareVersionParam) (*view.VirtualRouterSoftwareVersionInventoryView, error) {
	resp := view.VirtualRouterSoftwareVersionInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/vpc/virtual-routers/%s/softwareversion", uuid), "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
