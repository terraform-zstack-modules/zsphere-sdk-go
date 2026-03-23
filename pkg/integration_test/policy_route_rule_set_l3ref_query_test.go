// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPolicyRouteRuleSetL3Ref(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPolicyRouteRuleSetL3Ref(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPolicyRouteRuleSetL3Ref error: %v", err)
		return
	}
	golog.Infof("QueryPolicyRouteRuleSetL3Ref result count: %d", len(result))
}

