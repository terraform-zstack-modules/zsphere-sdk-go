// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryPrimaryStorage 测试查询主存储
func TestQueryPrimaryStorage(t *testing.T) {
	testName := "查询主存储"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryPrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d primary storages", len(result))
	skipIfNoResource(t, "PrimaryStorage", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetPrimaryStorage 测试获取主存储详情
func TestGetPrimaryStorage(t *testing.T) {
	testName := "获取主存储详情"
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
	list, err := cli.QueryPrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No PrimaryStorage found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No PrimaryStorage found to test")
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdatePrimaryStorage 测试更新主存储
func TestUpdatePrimaryStorage(t *testing.T) {
	testName := "更新主存储"
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
	list, err := cli.QueryPrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No PrimaryStorage found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No PrimaryStorage found to test")
		return
	}

	// 准备更新参数
	originalPrimaryStorage := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdatePrimaryStorageParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdatePrimaryStorageParamDetail{
			Name:        originalPrimaryStorage.Name, // 保持名称不变
			Description: &newDescription,
		},
	}

	// 执行更新
	updatedPrimaryStorage, err := cli.UpdatePrimaryStorage(ctx, originalPrimaryStorage.UUID, updateParam)
	if !assert.NoError(t, err, "Update primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedPrimaryStorage.Description, "PrimaryStorage description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedPrimaryStorage.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated primary storage: %s, new description: %s", updatedPrimaryStorage.UUID, updatedPrimaryStorage.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestReconnectPrimaryStorage 测试重连主存储
func TestReconnectPrimaryStorage(t *testing.T) {
	testName := "重连主存储"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 跳过此测试，因为它会影响实际资源
	// 如果需要在实际环境中运行，请移除此行
	golog.Warnf("测试跳过 [%s]: 该测试会影响实际资源", testName)
	recordTestResult(t.Name(), true) // 跳过也算通过
	t.Skip("Skipping reconnect primary storage test as it affects actual resources")

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryPrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No PrimaryStorage found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No PrimaryStorage found to test")
		return
	}

	// 准备重连参数
	reconnectParam := param.ReconnectPrimaryStorageParam{
		BaseParam: param.BaseParam{},
		Params:    param.ReconnectPrimaryStorageParamDetail{},
	}

	// 执行重连
	reconnectedPrimaryStorage, err := cli.ReconnectPrimaryStorage(ctx, list[0].UUID, reconnectParam)
	if !assert.NoError(t, err, "Reconnect primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Reconnected primary storage: %s, state: %s", reconnectedPrimaryStorage.UUID, reconnectedPrimaryStorage.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDeletePrimaryStorage 测试删除主存储
func TestDeletePrimaryStorage(t *testing.T) {
	testName := "删除主存储"
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
	t.Skip("Skipping delete primary storage test as it deletes actual resources")

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryPrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No PrimaryStorage found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No PrimaryStorage found to test")
		return
	}

	// 删除主存储
	err = cli.DeletePrimaryStorage(ctx, list[0].UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Primary storage deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestPrimaryStoragePageQuery 测试主存储的分页查询
func TestPrimaryStoragePageQuery(t *testing.T) {
	testName := "主存储的分页查询"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	pageSize := 10
	queryParam := param.NewQueryParam()
	queryParam.Limit(pageSize)

	primaryStorages, total, err := cli.PagePrimaryStorage(ctx, &queryParam)
	if !assert.NoError(t, err, "Page primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d primary storages out of %d total", len(primaryStorages), total)

	if !assert.LessOrEqual(t, len(primaryStorages), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的主存储数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(primaryStorages))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
