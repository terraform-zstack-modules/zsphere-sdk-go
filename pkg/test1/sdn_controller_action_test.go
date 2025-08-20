// Copyright (c) ZStack.io, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"testing"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

func TestZSClient_QuerySdnController(t *testing.T) {
	sndControllers, err := accountLoginCli.QuerySdnController(param.NewQueryParam())
	if err != nil {
		t.Errorf("TestZSClient_QuerySdnController %v", err)
		return
	}
	t.Log(sndControllers)
}

func TestZSClient_GetSdnController(t *testing.T) {
	sdnController, err := accountLoginCli.GetSdnController("65589889039944b5a2efeb2ed4d67594")
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	t.Log(sdnController)
}
