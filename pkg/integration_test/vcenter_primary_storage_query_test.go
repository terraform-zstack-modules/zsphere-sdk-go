// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVCenterPrimaryStorage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVCenterPrimaryStorage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVCenterPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QueryVCenterPrimaryStorage result count: %d", len(result))
}

