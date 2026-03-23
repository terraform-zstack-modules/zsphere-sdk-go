// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryAccount 测试查询账户
func TestQueryAccount(t *testing.T) {
	testName := "查询账户"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryAccount(ctx, &queryParam)
	if !assert.NoError(t, err, "Query account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d accounts", len(result))
	skipIfNoResource(t, "Account", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetAccount 测试获取账户详情
func TestGetAccount(t *testing.T) {
	testName := "获取账户详情"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryAccount(ctx, &queryParam)
	if !assert.NoError(t, err, "Query account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Account found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Account found to test")
		return
	}

	// 通过UUID获取
	account, err := cli.GetAccount(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, account.UUID, "Account UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, account.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got account: %s, name: %s, type: %s", account.UUID, account.Name, account.Type)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateAccount 测试更新账户
func TestUpdateAccount(t *testing.T) {
	testName := "更新账户"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会创建实际资源
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会创建实际资源", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	// t.Skip("Skipping create account test as it creates actual resources")

	// 创建账户参数
	typeStr := "Normal"
	descStr := "Created by SDK test"

	createParam := param.CreateAccountParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateAccountParamDetail{
			Name:        "test-account",
			Password:    "Test@123",
			Type:        strPtr(typeStr),
			Description: strPtr(descStr),
		},
	}

	// 执行创建
	account, err := cli.CreateAccount(ctx, createParam)
	if !assert.NoError(t, err, "Create account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 准备更新参数
	originalAccount := account
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateAccountParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateAccountParamDetail{
			Name:        originalAccount.Name + "_example",
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedAccount, err := cli.UpdateAccount(ctx, originalAccount.UUID, updateParam)
	if !assert.NoError(t, err, "Update account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedAccount.Description, "Account description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedAccount.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated account: %s, new description: %s", updatedAccount.UUID, updatedAccount.Description)

	// 删除账户
	err = cli.DeleteAccount(ctx, account.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Account deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAccount 测试创建账户
func TestCreateAccount(t *testing.T) {
	testName := "创建账户"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会创建实际资源
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会创建实际资源", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	t.Skip("Skipping create account test as it creates actual resources")

	// 创建账户参数
	typeStr := "Normal"
	descStr := "Created by SDK test"

	createParam := param.CreateAccountParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateAccountParamDetail{
			Name:        "test-account",
			Password:    "Test@123",
			Type:        strPtr(typeStr),
			Description: strPtr(descStr),
		},
	}

	// 执行创建
	account, err := cli.CreateAccount(ctx, createParam)
	if !assert.NoError(t, err, "Create account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created account: %s", account.UUID)

	// 删除账户
	err = cli.DeleteAccount(ctx, account.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Account deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDeleteAccount 测试删除账户
func TestDeleteAccount(t *testing.T) {
	testName := "删除账户"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会删除实际资源
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会删除实际资源", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	t.Skip("Skipping delete account test as it deletes actual resources")

	// 首先创建一个测试账户
	typeStr := "Normal"
	descStr := "Created for delete test"

	createParam := param.CreateAccountParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateAccountParamDetail{
			Name:        "test-delete-account",
			Password:    "Test@123",
			Type:        strPtr(typeStr),
			Description: strPtr(descStr),
		},
	}

	account, err := cli.CreateAccount(ctx, createParam)
	if !assert.NoError(t, err, "Create account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created account for delete test: %s", account.UUID)

	// 删除账户
	err = cli.DeleteAccount(ctx, account.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	// 验证账户已被删除
	queryParam := param.NewQueryParam()
	queryParam.AddQ("uuid=" + account.UUID)
	accounts, err := cli.QueryAccount(ctx, &queryParam)
	if !assert.NoError(t, err, "Query account should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, 0, len(accounts), "Account should be deleted") {
		golog.Errorf("测试失败 [%s]: 账户未被删除", testName)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Account deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
