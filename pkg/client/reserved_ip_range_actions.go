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

// DeleteReservedIpRange deletes ReservedIpRange
func (cli *ZSClient) DeleteReservedIpRange(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/l3-networks/reserved-ip-ranges", uuid, string(deleteMode))
}
// AddReservedIpRange adds ReservedIpRange
func (cli *ZSClient) AddReservedIpRange(ctx context.Context, l3NetworkUuid string, params param.AddReservedIpRangeParam) (*view.ReservedIpRangeInventoryView, error) {
	resp := view.ReservedIpRangeInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/l3-networks/%s/reserved-ip-ranges", l3NetworkUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
