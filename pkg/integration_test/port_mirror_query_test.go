// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryPortMirror(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryPortMirror(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryPortMirror error: %v", err)
		return
	}
	golog.Infof("QueryPortMirror result count: %d", len(result))
}

