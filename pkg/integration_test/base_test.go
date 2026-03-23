// Copyright (c) ZStack.io, Inc.
// Auto-generated integration test infrastructure. DO NOT EDIT.

package integration_test

import (
	"context"
	"os"
	"testing"

	"github.com/kataras/golog"
	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/client"
)

const (
	defaultHostname = "localhost"
	defaultAccount  = "admin"
	defaultPassword = "password"
)

var testCli *client.ZSClient

func TestMain(m *testing.M) {
	ctx := context.Background()
	hostname := os.Getenv("ZSTACK_HOST")
	if hostname == "" {
		hostname = defaultHostname
	}
	account := os.Getenv("ZSTACK_ACCOUNT")
	if account == "" {
		account = defaultAccount
	}
	password := os.Getenv("ZSTACK_PASSWORD")
	if password == "" {
		password = defaultPassword
	}

	config := client.DefaultZSConfig(hostname).
		LoginAccount(account, password).
		Debug(true)
	testCli = client.NewZSClient(config)

	_, err := testCli.Login(ctx)
	if err != nil {
		golog.Errorf("Integration test login failed: %v", err)
		os.Exit(1)
	}
	defer testCli.Logout(ctx)

	os.Exit(m.Run())
}
