// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryIpRange(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryIpRange(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryIpRange error: %v", err)
		return
	}
	golog.Infof("QueryIpRange result count: %d", len(result))
}

