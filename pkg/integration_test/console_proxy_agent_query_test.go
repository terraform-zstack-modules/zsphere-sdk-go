// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryConsoleProxyAgent(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryConsoleProxyAgent(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryConsoleProxyAgent error: %v", err)
		return
	}
	golog.Infof("QueryConsoleProxyAgent result count: %d", len(result))
}

