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

// AddSNSSmsReceiver adds SNSSmsReceiver
func (cli *ZSClient) AddSNSSmsReceiver(ctx context.Context, params param.AddSNSSmsReceiverParam) (*view.SNSSmsReceiverInventoryView, error) {
	resp := view.SNSSmsReceiverInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/sms-endpoints/receivers"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RemoveSNSSmsReceiver removes SNSSmsReceiver
func (cli *ZSClient) RemoveSNSSmsReceiver(ctx context.Context, endpointUuid string, phoneNumber string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/sms-endpoints", endpointUuid, fmt.Sprintf("receivers/%s", phoneNumber), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
