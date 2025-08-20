// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryZone queries the list of zones
func (cli *ZSClient) QueryZone(params param.QueryParam) ([]view.ZoneView, error) {
	resp := make([]view.ZoneView, 0)
	return resp, cli.List("v1/zones", &params, &resp)
}
