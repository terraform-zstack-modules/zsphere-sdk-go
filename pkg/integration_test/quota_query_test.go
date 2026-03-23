// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryQuota(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryQuota(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryQuota error: %v", err)
		return
	}
	golog.Infof("QueryQuota result count: %d", len(result))
}

