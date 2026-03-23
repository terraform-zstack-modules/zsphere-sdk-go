// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVirtualRouterOffering(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVirtualRouterOffering(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVirtualRouterOffering error: %v", err)
		return
	}
	golog.Infof("QueryVirtualRouterOffering result count: %d", len(result))
}

