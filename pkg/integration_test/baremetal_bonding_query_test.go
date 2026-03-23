// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryBaremetalBonding(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBaremetalBonding(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalBonding error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalBonding result count: %d", len(result))
}

