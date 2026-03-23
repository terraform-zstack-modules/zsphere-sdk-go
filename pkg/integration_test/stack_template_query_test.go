// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryStackTemplate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryStackTemplate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryStackTemplate error: %v", err)
		return
	}
	golog.Infof("QueryStackTemplate result count: %d", len(result))
}

