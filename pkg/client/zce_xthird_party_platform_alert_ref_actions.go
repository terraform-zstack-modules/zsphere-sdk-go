// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// QueryZceXThirdPartyPlatformAlertRef queries ZceXThirdPartyPlatformAlertRef list
func (cli *ZSClient) QueryZceXThirdPartyPlatformAlertRef(ctx context.Context, params *param.QueryParam) ([]view.ZceXThirdPartyPlatformAlertRefInventoryView, error) {
	var resp []view.ZceXThirdPartyPlatformAlertRefInventoryView
	return resp, cli.List(ctx, "v1/zce-x-plugin/alert-platform", params, &resp)
}

// PageZceXThirdPartyPlatformAlertRef Pagination
func (cli *ZSClient) PageZceXThirdPartyPlatformAlertRef(ctx context.Context, params *param.QueryParam) ([]view.ZceXThirdPartyPlatformAlertRefInventoryView, int, error) {
	var zceXThirdPartyPlatformAlertRefs []view.ZceXThirdPartyPlatformAlertRefInventoryView
	total, err := cli.Page(ctx, "v1/zce-x-plugin/alert-platform", params, &zceXThirdPartyPlatformAlertRefs)
	return zceXThirdPartyPlatformAlertRefs, total, err
}
