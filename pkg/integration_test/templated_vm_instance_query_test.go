// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryTemplatedVmInstance(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryTemplatedVmInstance(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryTemplatedVmInstance error: %v", err)
		return
	}
	golog.Infof("QueryTemplatedVmInstance result count: %d", len(result))
}

