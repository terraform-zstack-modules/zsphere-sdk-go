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

// RemoveAccessControlListEntry removes AccessControlListEntry
func (cli *ZSClient) RemoveAccessControlListEntry(ctx context.Context, aclUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/access-control-lists", aclUuid, fmt.Sprintf("ipentries/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddAccessControlListEntry adds AccessControlListEntry
func (cli *ZSClient) AddAccessControlListEntry(ctx context.Context, aclUuid string, params param.AddAccessControlListEntryParam) (*view.AccessControlListEntryInventoryView, error) {
	resp := view.AccessControlListEntryInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/access-control-lists/%s/ipentries", aclUuid), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
