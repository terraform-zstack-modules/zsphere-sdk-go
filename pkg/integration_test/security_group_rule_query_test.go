// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySecurityGroupRule(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySecurityGroupRule(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("QuerySecurityGroupRule result count: %d", len(result))
}

