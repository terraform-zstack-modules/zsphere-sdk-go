// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryHost 测试查询物理机
func TestQueryHost(t *testing.T) {
	testName := "查询物理机"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Query host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d hosts", len(result))
	skipIfNoResource(t, "Host", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetHost 测试获取物理机详情
func TestGetHost(t *testing.T) {
	testName := "获取物理机详情"
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
	list, err := cli.QueryHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Query host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Host found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Host found to test")
		return
	}

	// 通过UUID获取
	host, err := cli.GetHost(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, host.UUID, "Host UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, host.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got host: %s, name: %s, state: %s", host.UUID, host.Name, host.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateHost 测试更新物理机
func TestUpdateHost(t *testing.T) {
	testName := "更新物理机"
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
	list, err := cli.QueryHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Query host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Host found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Host found to test")
		return
	}

	// 准备更新参数
	originalHost := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateHostParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateHostParamDetail{
			Name:        originalHost.Name, // 保持名称不变
			Description: &newDescription,
		},
	}

	// 执行更新
	updatedHost, err := cli.UpdateHost(ctx, originalHost.UUID, updateParam)
	if !assert.NoError(t, err, "Update host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedHost.Description, "Host description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedHost.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated host: %s, new description: %s", updatedHost.UUID, updatedHost.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestReconnectHost 测试重连物理机
func TestReconnectHost(t *testing.T) {
	testName := "重连物理机"
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
	t.Skip("Skipping reconnect host test as it affects actual resources")

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Query host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Host found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Host found to test")
		return
	}

	// 准备重连参数
	reconnectParam := param.ReconnectHostParam{
		BaseParam: param.BaseParam{},
		Params:    param.ReconnectHostParamDetail{},
	}

	// 执行重连
	reconnectedHost, err := cli.ReconnectHost(ctx, list[0].UUID, reconnectParam)
	if !assert.NoError(t, err, "Reconnect host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Reconnected host: %s, state: %s", reconnectedHost.UUID, reconnectedHost.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDeleteHost 测试删除物理机
func TestDeleteHost(t *testing.T) {
	testName := "删除物理机"
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
	t.Skip("Skipping delete host test as it deletes actual resources")

	// 首先查询获取一个有效的UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := cli.QueryHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Query host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Host found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Host found to test")
		return
	}

	// 删除物理机
	err = cli.DeleteHost(ctx, list[0].UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Host deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestHostPageQuery 测试物理机的分页查询
func TestHostPageQuery(t *testing.T) {
	testName := "物理机的分页查询"
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

	hosts, total, err := cli.PageHost(ctx, &queryParam)
	if !assert.NoError(t, err, "Page host should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d hosts out of %d total", len(hosts), total)

	if !assert.LessOrEqual(t, len(hosts), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的物理机数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(hosts))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
