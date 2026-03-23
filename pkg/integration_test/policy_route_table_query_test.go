// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPolicyRouteTable(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPolicyRouteTable(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteTable error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteTable result count: %d", len(result))
}

