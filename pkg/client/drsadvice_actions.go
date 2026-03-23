// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ view.MapView // avoid unused import

// ApplyDRSAdvice operates on DRSAdvice
func (cli *ZSClient) ApplyDRSAdvice(ctx context.Context, adviceUuid string, params param.ApplyDRSAdviceParam) (*view.DRSAdviceInventoryView, error) {
	resp := view.DRSAdviceInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters/drs/advice", adviceUuid, "actions", "", map[string]interface{}{
		"applyDRSAdvice": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDRSAdvice queries DRSAdvice list
func (cli *ZSClient) QueryDRSAdvice(ctx context.Context, params *param.QueryParam) ([]view.DRSAdviceInventoryView, error) {
	var resp []view.DRSAdviceInventoryView
	return resp, cli.List(ctx, "v1/clusters/drs/advice", params, &resp)
}

func (cli *ZSClient) GetDRSAdvice(ctx context.Context, uuid string) (*view.DRSAdviceInventoryView, error) {
	var resp view.DRSAdviceInventoryView
	if err := cli.Get(ctx, "v1/clusters/drs/advice", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDRSAdvice Pagination
func (cli *ZSClient) PageDRSAdvice(ctx context.Context, params *param.QueryParam) ([]view.DRSAdviceInventoryView, int, error) {
	var dRSAdvices []view.DRSAdviceInventoryView
	total, err := cli.Page(ctx, "v1/clusters/drs/advice", params, &dRSAdvices)
	return dRSAdvices, total, err
}
