// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVmInstance(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstance(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryVmInstance result count: %d", len(result))
}

