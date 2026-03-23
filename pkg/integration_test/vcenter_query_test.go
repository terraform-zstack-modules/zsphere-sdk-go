// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVCenter(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVCenter(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenter error: %v", err)
		return
	}
	golog.Infof("QueryVCenter result count: %d", len(result))
}

