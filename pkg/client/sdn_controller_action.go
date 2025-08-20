// Copyright (c) ZStack.io, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// Query SDN Controller
func (cli *ZSClient) QuerySdnController(params param.QueryParam) ([]view.SdnControllerInventoryView, error) {
	var resp []view.SdnControllerInventoryView
	return resp, cli.List("v1/sdn-controllers", &params, &resp)
}

// Get SDN Controller by UUID
func (cli *ZSClient) GetSdnController(uuid string) (*view.SdnControllerInventoryView, error) {
	var resp view.SdnControllerInventoryView
	if err := cli.Get("v1/sdn-controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
