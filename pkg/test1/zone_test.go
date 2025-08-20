// # Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestQueryZone(t *testing.T) {
	zone, err := accountLoginCli.QueryZone(param.NewQueryParam())
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(zone)
}
