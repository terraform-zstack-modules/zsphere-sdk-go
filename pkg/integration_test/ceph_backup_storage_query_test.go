// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryCephBackupStorage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCephBackupStorage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryCephBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryCephBackupStorage result count: %d", len(result))
}

