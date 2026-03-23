// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryUserTag(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryUserTag(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryUserTag error: %v", err)
		return
	}
	golog.Infof("QueryUserTag result count: %d", len(result))
}

