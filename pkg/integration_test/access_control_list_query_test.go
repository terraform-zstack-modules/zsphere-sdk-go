// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryAccessControlList(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryAccessControlList(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessControlList error: %v", err)
		return
	}
	golog.Infof("QueryAccessControlList result count: %d", len(result))
}

