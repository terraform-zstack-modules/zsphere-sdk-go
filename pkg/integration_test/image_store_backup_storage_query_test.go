// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryImageStoreBackupStorage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryImageStoreBackupStorage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryImageStoreBackupStorage error: %v", err)
		return
	}
	golog.Infof("QueryImageStoreBackupStorage result count: %d", len(result))
}

