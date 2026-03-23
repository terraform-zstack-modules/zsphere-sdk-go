// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVirtualRouterVm(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVirtualRouterVm(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterVm error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterVm result count: %d", len(result))
}

