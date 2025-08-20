// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/view"
)

// QueryCluster queries ZStack Clusters
func (cli *ZSClient) QueryCluster(params param.QueryParam) ([]view.ClusterInventoryView, error) {
	resp := make([]view.ClusterInventoryView, 0)
	return resp, cli.List("v1/clusters", &params, &resp)
}
