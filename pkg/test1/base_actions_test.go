// # Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
)

const (
	// The ZStack Cloud Community Edition only supports login authentication for super admin accounts.
	// The ZStack Cloud Basic Edition supports login authentication for AccessKey, super admin, and sub-accounts.
	// The ZStack Cloud Enterprise Edition supports login authentication for AccessKey, super admin, sub-accounts, and enterprise users.

	accountLoginHostname        = "test-zsphere.zstack.io" //ZStack Cloud API endpoint IP address
	accountLoginAccountName     = "admin"
	accountLoginAccountPassword = "zsv@2025"

	accountLoginMasterHostname = "IPOfCloudAPIEndpoint"
	accountLoginSlaveHostname  = "IPOfCloudAPIEndpoint"

	accessKeyAuthHostname        = "test-zsphere.zstack.io"
	accessKeyAuthAccessKeyId     = "8bwZmAKXzY3ICbyAV6Dk"
	accessKeyAuthAccessKeySecret = "jIxkhbgBZhglg8DvRkPD1Y4TW9FCthS3hYJrk1pg"

	readOnly = false
	debug    = false
)

var accountLoginCli = client.NewZSClient(
	client.DefaultZSConfig(accountLoginHostname).
		LoginAccount(accountLoginAccountName, accountLoginAccountPassword).
		ReadOnly(readOnly).
		Debug(true),
)

var accessKeyAuthCli = client.NewZSClient(
	client.DefaultZSConfig(accessKeyAuthHostname).
		AccessKey(accessKeyAuthAccessKeyId, accessKeyAuthAccessKeySecret).
		ReadOnly(readOnly).
		Debug(debug),
)

func TestMain(m *testing.M) {
	_, err := accountLoginCli.Login()
	if err != nil {
		golog.Errorf("TestMain err %v", err)
	}
	defer accountLoginCli.Logout()

	m.Run()
}
