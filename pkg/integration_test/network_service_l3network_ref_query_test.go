// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryNetworkServiceL3NetworkRef(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryNetworkServiceL3NetworkRef(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryNetworkServiceL3NetworkRef error: %v", err)
		return
	}
	golog.Infof("QueryNetworkServiceL3NetworkRef result count: %d", len(result))
}

