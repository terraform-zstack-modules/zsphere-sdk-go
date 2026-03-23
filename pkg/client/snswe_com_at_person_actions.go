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

// RemoveSNSWeComAtPerson removes SNSWeComAtPerson
func (cli *ZSClient) RemoveSNSWeComAtPerson(ctx context.Context, endpointUuid string, userId string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec(ctx, "v1/sns/application-endpoints/we-com", endpointUuid, fmt.Sprintf("at-persons/%s", userId), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
// AddSNSWeComAtPerson adds SNSWeComAtPerson
func (cli *ZSClient) AddSNSWeComAtPerson(ctx context.Context, params param.AddSNSWeComAtPersonParam) (*view.SNSWeComAtPersonInventoryView, error) {
	resp := view.SNSWeComAtPersonInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/sns/application-endpoints/we-com/at-persons"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSWeComAtPerson queries SNSWeComAtPerson list
func (cli *ZSClient) QuerySNSWeComAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, error) {
	var resp []view.SNSWeComAtPersonInventoryView
	return resp, cli.List(ctx, "v1/sns/application-endpoints/we-com/at-persons", params, &resp)
}

func (cli *ZSClient) GetSNSWeComAtPerson(ctx context.Context, uuid string) (*view.SNSWeComAtPersonInventoryView, error) {
	var resp view.SNSWeComAtPersonInventoryView
	if err := cli.Get(ctx, "v1/sns/application-endpoints/we-com/at-persons", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSWeComAtPerson Pagination
func (cli *ZSClient) PageSNSWeComAtPerson(ctx context.Context, params *param.QueryParam) ([]view.SNSWeComAtPersonInventoryView, int, error) {
	var sNSWeComAtPersons []view.SNSWeComAtPersonInventoryView
	total, err := cli.Page(ctx, "v1/sns/application-endpoints/we-com/at-persons", params, &sNSWeComAtPersons)
	return sNSWeComAtPersons, total, err
}
