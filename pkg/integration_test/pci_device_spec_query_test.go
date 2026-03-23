// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPciDeviceSpec(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPciDeviceSpec(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDeviceSpec error: %v", err)
		return
	}
	golog.Infof("QueryPciDeviceSpec result count: %d", len(result))
}

