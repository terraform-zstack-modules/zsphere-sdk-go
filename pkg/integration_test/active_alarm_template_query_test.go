// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryActiveAlarmTemplate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryActiveAlarmTemplate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryActiveAlarmTemplate error: %v", err)
		return
	}
	golog.Infof("QueryActiveAlarmTemplate result count: %d", len(result))
}

