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

// AddSimulatorHost adds SimulatorHost
func (cli *ZSClient) AddSimulatorHost(ctx context.Context, params param.AddSimulatorHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/hosts/simulators"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
