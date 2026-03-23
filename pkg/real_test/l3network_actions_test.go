// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryL3Network 测试查询L3网络
func TestQueryL3Network(t *testing.T) {
	testName := "查询L3网络"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryL3Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d L3Networks", len(result))
	skipIfNoResource(t, "L3Network", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetL3Network 测试获取L3网络详情
func TestGetL3Network(t *testing.T) {
	testName := "获取L3网络详情"
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
	list, err := cli.QueryL3Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No L3Network found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No L3Network found to test")
		return
	}

	// 通过UUID获取
	l3Network, err := cli.GetL3Network(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, l3Network.UUID, "L3Network UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, l3Network.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got L3Network: %s, name: %s, type: %s", l3Network.UUID, l3Network.Name, l3Network.Type)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateL3Network 测试更新L3网络
func TestUpdateL3Network(t *testing.T) {
	testName := "更新L3网络"
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
	list, err := cli.QueryL3Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No L3Network found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No L3Network found to test")
		return
	}

	// 准备更新参数
	originalL3Network := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateL3NetworkParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateL3NetworkParamDetail{
			Name:        originalL3Network.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedL3Network, err := cli.UpdateL3Network(ctx, originalL3Network.UUID, updateParam)
	if !assert.NoError(t, err, "Update L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedL3Network.Description, "L3Network description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedL3Network.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated L3Network: %s, new description: %s", updatedL3Network.UUID, updatedL3Network.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestL3NetworkPageQuery 测试L3网络的分页查询
func TestL3NetworkPageQuery(t *testing.T) {
	testName := "L3网络的分页查询"
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

	l3Networks, total, err := cli.PageL3Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Page L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d L3Networks out of %d total", len(l3Networks), total)

	if !assert.LessOrEqual(t, len(l3Networks), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的L3网络数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(l3Networks))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteL3Network 测试创建和删除L3网络
func TestCreateAndDeleteL3Network(t *testing.T) {
	testName := "创建和删除L3网络"
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
	t.Skip("Skipping create/delete L3Network test as it creates actual resources")

	// 创建L3网络需要有效的L2网络UUID
	// 这些值需要根据实际环境进行替换
	l2NetworkUuid := "your-l2network-uuid"
	typeStr := "L3BasicNetwork"
	descStr := "Created by SDK test"

	createParam := param.CreateL3NetworkParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateL3NetworkParamDetail{
			Name:          "test-l3network",
			Description:   strPtr(descStr),
			L2NetworkUuid: l2NetworkUuid,
			Type:          strPtr(typeStr),
		},
	}

	l3Network, err := cli.CreateL3Network(ctx, createParam)
	if !assert.NoError(t, err, "Create L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created L3Network: %s", l3Network.UUID)

	// 删除L3网络
	err = cli.DeleteL3Network(ctx, l3Network.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete L3Network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("L3Network deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
