// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryUsbDevice(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryUsbDevice(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryUsbDevice error: %v", err)
		return
	}
	golog.Infof("QueryUsbDevice result count: %d", len(result))
}

