// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryBackupStorage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryBackupStorage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryBackupStorage result count: %d", len(result))
}

