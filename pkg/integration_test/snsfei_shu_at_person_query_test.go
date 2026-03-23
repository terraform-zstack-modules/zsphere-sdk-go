// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQuerySNSFeiShuAtPerson(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QuerySNSFeiShuAtPerson(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSFeiShuAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSFeiShuAtPerson result count: %d", len(result))
}

