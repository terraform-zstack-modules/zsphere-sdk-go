// # Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/util/jsonutils"
)

func TestQueryManagementNode(t *testing.T) {
	data, err := accountLoginCli.QueryManagementNode(param.NewQueryParam())
	if err != nil {
		t.Errorf("error %v ", err)
	}
	golog.Infof("%v", jsonutils.Marshal(data))
}

func TestGetVersion(t *testing.T) {
	data, err := accountLoginCli.GetVersion()
	if err != nil {
		t.Errorf("error %v ", err)
	}
	golog.Infof("%v", data)
}
