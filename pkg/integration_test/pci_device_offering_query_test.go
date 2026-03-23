// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPciDeviceOffering(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPciDeviceOffering(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPciDeviceOffering error: %v", err)
		return
	}
	golog.Infof("QueryPciDeviceOffering result count: %d", len(result))
}

