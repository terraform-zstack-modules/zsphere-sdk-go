// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryActiveAlarm(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryActiveAlarm(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryActiveAlarm error: %v", err)
		return
	}
	golog.Infof("QueryActiveAlarm result count: %d", len(result))
}

