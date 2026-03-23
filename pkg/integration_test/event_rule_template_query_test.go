// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryEventRuleTemplate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryEventRuleTemplate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryEventRuleTemplate error: %v", err)
		return
	}
	golog.Infof("QueryEventRuleTemplate result count: %d", len(result))
}

