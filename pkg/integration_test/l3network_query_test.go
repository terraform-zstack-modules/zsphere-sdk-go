// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryL3Network(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryL3Network(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryL3Network error: %v", err)
		return
	}
	golog.Infof("QueryL3Network result count: %d", len(result))
}

