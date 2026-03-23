// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryImagePackage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryImagePackage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryImagePackage error: %v", err)
		return
	}
	golog.Infof("QueryImagePackage result count: %d", len(result))
}

