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

// RemoveSNSDingTalkAtPerson removes SNSDingTalkAtPerson
func (cli *ZSClient) RemoveSNSDingTalkAtPerson(ctx context.Context, endpointUuid string, phoneNumber string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/application-endpoints/ding-talk", endpointUuid, fmt.Sprintf("at-persons/%s", phoneNumber), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddSNSDingTalkAtPerson adds SNSDingTalkAtPerson
func (cli *ZSClient) AddSNSDingTalkAtPerson(ctx context.Context, params param.AddSNSDingTalkAtPersonParam) (*view.SNSDingTalkAtPersonInventoryView, error) {
	resp := view.SNSDingTalkAtPersonInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/ding-talk/at-persons"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSDingTalkAtPerson queries SNSDingTalkAtPerson list
func (cli *ZSClient) QuerySNSDingTalkAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, error) {
	var resp []view.SNSDingTalkAtPersonInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/ding-talk/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSDingTalkAtPerson(ctx context.Context, uuid string) (*view.SNSDingTalkAtPersonInventoryView, error) {
	var resp view.SNSDingTalkAtPersonInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/ding-talk/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSDingTalkAtPerson Pagination
func (cli *ZSClient) PageSNSDingTalkAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSDingTalkAtPersonInventoryView, int, error) {
	var sNSDingTalkAtPersons []view.SNSDingTalkAtPersonInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/ding-talk/at-persons", params, &sNSDingTalkAtPersons)
	return sNSDingTalkAtPersons, total, err
}
