// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryImage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryImage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryImage error: %v", err)
		return
	}
	golog.Infof("QueryImage result count: %d", len(result))
}

