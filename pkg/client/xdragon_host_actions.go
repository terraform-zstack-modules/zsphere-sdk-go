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

// AddXDragonHost adds XDragonHost
func (cli *ZSClient) AddXDragonHost(ctx context.Context, params param.AddXDragonHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/xdragon"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
