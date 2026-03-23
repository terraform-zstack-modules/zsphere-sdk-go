// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySharedBlock(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySharedBlock(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlock error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlock result count: %d", len(result))
}

