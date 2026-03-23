// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySlbOffering(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySlbOffering(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySlbOffering error: %v", err)
		return
	}
	golog.Infof("QuerySlbOffering result count: %d", len(result))
}

