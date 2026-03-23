// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVmInstancePciDeviceSpecRef(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstancePciDeviceSpecRef(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstancePciDeviceSpecRef error: %v", err)
		return
	}
	golog.Infof("QueryVmInstancePciDeviceSpecRef result count: %d", len(result))
}

