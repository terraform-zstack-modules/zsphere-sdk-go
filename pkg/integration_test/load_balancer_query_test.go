// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryLoadBalancer(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryLoadBalancer(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryLoadBalancer error: %v", err)
		return
	}
	golog.Infof("QueryLoadBalancer result count: %d", len(result))
}

