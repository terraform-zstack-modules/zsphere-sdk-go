// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorage(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySharedBlockGroupPrimaryStorage(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorage result count: %d", len(result))
}

