// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryManagementNode(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryManagementNode(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryManagementNode error: %v", err)
		return
	}
	golog.Infof("QueryManagementNode result count: %d", len(result))
}

