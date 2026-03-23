// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVmInstanceResourceMetadataGroup(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVmInstanceResourceMetadataGroup(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVmInstanceResourceMetadataGroup error: %v", err)
		return
	}
	golog.Infof("QueryVmInstanceResourceMetadataGroup result count: %d", len(result))
}

