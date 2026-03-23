// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryCertificate(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryCertificate(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryCertificate error: %v", err)
		return
	}
	golog.Infof("QueryCertificate result count: %d", len(result))
}

