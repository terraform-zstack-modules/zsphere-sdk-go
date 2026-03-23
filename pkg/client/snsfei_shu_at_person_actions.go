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

// QuerySNSFeiShuAtPerson queries SNSFeiShuAtPerson list
func (cli *ZSClient) QuerySNSFeiShuAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, error) {
	var resp []view.SNSFeiShuAtPersonInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/feishu/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSFeiShuAtPerson(ctx context.Context, uuid string) (*view.SNSFeiShuAtPersonInventoryView, error) {
	var resp view.SNSFeiShuAtPersonInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/feishu/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSFeiShuAtPerson Pagination
func (cli *ZSClient) PageSNSFeiShuAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSFeiShuAtPersonInventoryView, int, error) {
	var sNSFeiShuAtPersons []view.SNSFeiShuAtPersonInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/feishu/at-persons", params, &sNSFeiShuAtPersons)
	return sNSFeiShuAtPersons, total, err
}
// RemoveSNSFeiShuAtPerson removes SNSFeiShuAtPerson
func (cli *ZSClient) RemoveSNSFeiShuAtPerson(ctx context.Context, endpointUuid string, userId string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/application-endpoints/feishu", endpointUuid, fmt.Sprintf("at-persons/%s", userId), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddSNSFeiShuAtPerson adds SNSFeiShuAtPerson
func (cli *ZSClient) AddSNSFeiShuAtPerson(ctx context.Context, params param.AddSNSFeiShuAtPersonParam) (*view.SNSFeiShuAtPersonInventoryView, error) {
	resp := view.SNSFeiShuAtPersonInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/feishu/at-persons"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
