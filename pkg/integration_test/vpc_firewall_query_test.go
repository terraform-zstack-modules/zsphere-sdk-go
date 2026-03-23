// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryVpcFirewall(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryVpcFirewall(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryVpcFirewall error: %v", err)
		return
	}
	golog.Infof("QueryVpcFirewall result count: %d", len(result))
}

