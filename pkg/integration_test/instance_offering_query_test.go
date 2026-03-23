// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryInstanceOffering(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryInstanceOffering(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryInstanceOffering error: %v", err)
		return
	}
	golog.Infof("QueryInstanceOffering result count: %d", len(result))
}

