// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryResourceConfig(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryResourceConfig(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryResourceConfig error: %v", err)
		return
	}
	golog.Infof("QueryResourceConfig result count: %d", len(result))
}

