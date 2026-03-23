// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// UpdateHostIpmi updates HostIpmi
func (cli *ZSClient) UpdateHostIpmi(ctx context.Context, uuid string, params param.UpdateHostIpmiParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/hosts/ipmi", uuid, "actions", "", map[string]interface{}{
		"updateHostIpmi": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
