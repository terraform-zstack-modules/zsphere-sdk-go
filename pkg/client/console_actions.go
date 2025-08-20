// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// RequestConsoleAccess Request console access URL
func (cli *ZSClient) RequestConsoleAccess(params param.RequestConsoleAccessParam) (view.ConsoleInventoryView, error) {
	var resp view.ConsoleInventoryView
	return resp, cli.Post("v1/consoles", &params, &resp)
}
