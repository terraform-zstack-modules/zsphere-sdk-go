// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryLocalStorageResourceRef(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryLocalStorageResourceRef(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryLocalStorageResourceRef error: %v", err)
		return
	}
	golog.Infof("QueryLocalStorageResourceRef result count: %d", len(result))
}

