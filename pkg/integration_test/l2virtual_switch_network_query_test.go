// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package integration_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryL2VirtualSwitchNetwork(t *testing.T) {
	ctx := context.Background()
	queryParam := param.NewQueryParam()
	result, err := testCli.QueryL2VirtualSwitchNetwork(ctx, &queryParam)
	if err != nil {
		t.Errorf("TestQueryL2VirtualSwitchNetwork error: %v", err)
		return
	}
	golog.Infof("QueryL2VirtualSwitchNetwork result count: %d", len(result))
}

