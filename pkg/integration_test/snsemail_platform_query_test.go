// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySNSEmailPlatform(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSEmailPlatform(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailPlatform error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailPlatform result count: %d", len(result))
}

