// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPreconfigurationTemplate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPreconfigurationTemplate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPreconfigurationTemplate error: %v", err)
		return
	}
	golog.Infof("QueryPreconfigurationTemplate result count: %d", len(result))
}

