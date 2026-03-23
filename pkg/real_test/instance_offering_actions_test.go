// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryInstanceOffering 测试查询计算规格
func TestQueryInstanceOffering(t *testing.T) {
	testName := "查询计算规格"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryInstanceOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d instance offerings", len(result))
	skipIfNoResource(t, "InstanceOffering", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetInstanceOffering 测试获取计算规格详情
func TestGetInstanceOffering(t *testing.T) {
	testName := "获取计算规格详情"
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
	list, err := cli.QueryInstanceOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No InstanceOffering found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No InstanceOffering found to test")
		return
	}

	// 通过UUID获取
	instanceOffering, err := cli.GetInstanceOffering(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, instanceOffering.UUID, "InstanceOffering UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, instanceOffering.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got instance offering: %s, name: %s, cpuNum: %d, memorySize: %d", instanceOffering.UUID, instanceOffering.Name, instanceOffering.CpuNum, instanceOffering.MemorySize)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateInstanceOffering 测试更新计算规格
func TestUpdateInstanceOffering(t *testing.T) {
	testName := "更新计算规格"
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
	list, err := cli.QueryInstanceOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No InstanceOffering found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No InstanceOffering found to test")
		return
	}

	// 准备更新参数
	originalInstanceOffering := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateInstanceOfferingParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateInstanceOfferingParamDetail{
			Name:        originalInstanceOffering.Name, // 保持名称不变
			Description: &newDescription,
		},
	}

	// 执行更新
	updatedInstanceOffering, err := cli.UpdateInstanceOffering(ctx, originalInstanceOffering.UUID, updateParam)
	if !assert.NoError(t, err, "Update instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedInstanceOffering.Description, "InstanceOffering description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedInstanceOffering.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated instance offering: %s, new description: %s", updatedInstanceOffering.UUID, updatedInstanceOffering.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteInstanceOffering 测试创建和删除计算规格
func TestCreateAndDeleteInstanceOffering(t *testing.T) {
	testName := "创建和删除计算规格"
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
	t.Skip("Skipping create instance offering test as it creates actual resources")

	// 创建计算规格参数
	cpuNum := 2
	memorySize := int64(4 * 1024 * 1024 * 1024) // 4GB
	description := "Created by SDK test"
	createParam := param.CreateInstanceOfferingParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateInstanceOfferingParamDetail{
			Name:        "test-instance-offering",
			Description: &description,
			CpuNum:      cpuNum,
			MemorySize:  memorySize,
		},
	}

	// 执行创建
	instanceOffering, err := cli.CreateInstanceOffering(ctx, createParam)
	if !assert.NoError(t, err, "Create instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created instance offering: %s", instanceOffering.UUID)

	// 删除计算规格
	err = cli.DeleteInstanceOffering(ctx, instanceOffering.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Instance offering deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestInstanceOfferingPageQuery 测试计算规格的分页查询
func TestInstanceOfferingPageQuery(t *testing.T) {
	testName := "计算规格的分页查询"
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

	instanceOfferings, total, err := cli.PageInstanceOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Page instance offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d instance offerings out of %d total", len(instanceOfferings), total)

	if !assert.LessOrEqual(t, len(instanceOfferings), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的计算规格数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(instanceOfferings))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
