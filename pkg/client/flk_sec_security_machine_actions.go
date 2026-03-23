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

// UpdateFlkSecSecurityMachine updates FlkSecSecurityMachine
func (cli *ZSClient) UpdateFlkSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-machines/flkSec", uuid, "actions", "", map[string]interface{}{
		"updateFlkSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddFlkSecSecurityMachine adds FlkSecSecurityMachine
func (cli *ZSClient) AddFlkSecSecurityMachine(ctx context.Context, params param.AddFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-machine/flkSec"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
