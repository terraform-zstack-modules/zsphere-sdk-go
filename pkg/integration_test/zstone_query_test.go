// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryZStone(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryZStone(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryZStone error: %v", err)
		return
	}
	golog.Infof("QueryZStone result count: %d", len(result))
}

