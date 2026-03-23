// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySlbGroup(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySlbGroup(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbGroup error: %v", err)
		return
	}
	golog.Infof("QuerySlbGroup result count: %d", len(result))
}

