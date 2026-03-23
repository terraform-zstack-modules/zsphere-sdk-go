// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryL2Network 测试查询二层网络
func TestQueryL2Network(t *testing.T) {
	testName := "查询二层网络"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryL2Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d L2 networks", len(result))
	skipIfNoResource(t, "L2Network", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetL2Network 测试获取二层网络详情
func TestGetL2Network(t *testing.T) {
	testName := "获取二层网络详情"
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
	list, err := cli.QueryL2Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No L2Network found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No L2Network found to test")
		return
	}

	// 通过UUID获取
	l2Network, err := cli.GetL2Network(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, l2Network.UUID, "L2Network UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, l2Network.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got L2 network: %s, name: %s, type: %s", l2Network.UUID, l2Network.Name, l2Network.Type)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateL2Network 测试更新二层网络
func TestUpdateL2Network(t *testing.T) {
	testName := "更新二层网络"
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
	list, err := cli.QueryL2Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No L2Network found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No L2Network found to test")
		return
	}

	// 准备更新参数
	originalL2Network := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateL2NetworkParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateL2NetworkParamDetail{
			Name:        originalL2Network.Name, // 保持名称不变
			Description: &newDescription,
		},
	}

	// 执行更新
	updatedL2Network, err := cli.UpdateL2Network(ctx, originalL2Network.UUID, updateParam)
	if !assert.NoError(t, err, "Update L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedL2Network.Description, "L2Network description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedL2Network.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated L2 network: %s, new description: %s", updatedL2Network.UUID, updatedL2Network.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDeleteL2Network 测试删除二层网络
func TestDeleteL2Network(t *testing.T) {
	testName := "删除二层网络"
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
	t.Skip("Skipping delete L2 network test as it deletes actual resources")

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryL2Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Query L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No L2Network found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No L2Network found to test")
		return
	}

	// 删除二层网络
	err = cli.DeleteL2Network(ctx, list[0].UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("L2 network deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestL2NetworkPageQuery 测试二层网络的分页查询
func TestL2NetworkPageQuery(t *testing.T) {
	testName := "二层网络的分页查询"
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

	l2Networks, total, err := cli.PageL2Network(ctx, &queryParam)
	if !assert.NoError(t, err, "Page L2 network should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d L2 networks out of %d total", len(l2Networks), total)

	if !assert.LessOrEqual(t, len(l2Networks), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的二层网络数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(l2Networks))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
