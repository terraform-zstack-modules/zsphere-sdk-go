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

// UpdateInfoSecSecurityMachine updates InfoSecSecurityMachine
func (cli *ZSClient) UpdateInfoSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/security-machines/infoSec", uuid, "actions", "", map[string]interface{}{
		"updateInfoSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddInfoSecSecurityMachine adds InfoSecSecurityMachine
func (cli *ZSClient) AddInfoSecSecurityMachine(ctx context.Context, params param.AddInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/security-machine/infoSec"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
