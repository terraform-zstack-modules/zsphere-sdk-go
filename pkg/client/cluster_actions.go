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

// DeleteCluster deletes Cluster
func (cli *ZSClient) DeleteCluster(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/clusters", uuid, string(deleteMode))
}
// UpdateCluster updates Cluster
func (cli *ZSClient) UpdateCluster(ctx context.Context, uuid string, params param.UpdateClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PutWithSpec(ctx, "v1/clusters", uuid, "actions", "inventory", map[string]interface{}{
		"updateCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCluster queries Cluster list
func (cli *ZSClient) QueryCluster(ctx context.Context, params *param.QueryParam) ([]view.ClusterInventoryView, error) {
	var resp []view.ClusterInventoryView
	return resp, cli.List(ctx, "v1/clusters", params, &resp)
}

func (cli *ZSClient) GetCluster(ctx context.Context, uuid string) (*view.ClusterInventoryView, error) {
	var resp view.ClusterInventoryView
	if err := cli.Get(ctx, "v1/clusters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCluster Pagination
func (cli *ZSClient) PageCluster(ctx context.Context, params *param.QueryParam) ([]view.ClusterInventoryView, int, error) {
	var clusters []view.ClusterInventoryView
	total, err := cli.Page(ctx, "v1/clusters", params, &clusters)
	return clusters, total, err
}
// CreateCluster creates Cluster
func (cli *ZSClient) CreateCluster(ctx context.Context, params param.CreateClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PostWithRespKey(ctx, fmt.Sprintf("v1/clusters"), "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
