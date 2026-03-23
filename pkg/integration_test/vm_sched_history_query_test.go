// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVmSchedHistory(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmSchedHistory(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVmSchedHistory error: %v", err)
		return
	}
	golog.Infof("QueryVmSchedHistory result count: %d", len(result))
}

