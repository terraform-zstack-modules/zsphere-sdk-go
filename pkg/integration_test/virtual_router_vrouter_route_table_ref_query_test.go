// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVirtualRouterVRouterRouteTableRef(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVirtualRouterVRouterRouteTableRef(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVRouterRouteTableRef error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVRouterRouteTableRef result count: %d", len(result))
}

