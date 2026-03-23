// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryL2PortGroupNetwork(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryL2PortGroupNetwork(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryL2PortGroupNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2PortGroupNetwork result count: %d", len(result))
}

