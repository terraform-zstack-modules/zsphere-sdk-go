// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryFiberChannelLun(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryFiberChannelLun(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryFiberChannelLun error: %v", err)
		return
	}
	golog.Infof("QueryFiberChannelLun result count: %d", len(result))
}

