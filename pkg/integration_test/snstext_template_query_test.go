// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySNSTextTemplate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSTextTemplate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSTextTemplate error: %v", err)
		return
	}
	golog.Infof("QuerySNSTextTemplate result count: %d", len(result))
}

