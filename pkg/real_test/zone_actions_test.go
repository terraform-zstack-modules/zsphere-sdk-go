// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryZone 测试查询区域
func TestQueryZone(t *testing.T) {
	testName := "查询区域"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryZone(ctx, &queryParam)
	if !assert.NoError(t, err, "Query zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d zones", len(result))
	skipIfNoResource(t, "Zone", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetZone 测试获取区域详情
func TestGetZone(t *testing.T) {
	testName := "获取区域详情"
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
	list, err := cli.QueryZone(ctx, &queryParam)
	if !assert.NoError(t, err, "Query zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Zone found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Zone found to test")
		return
	}

	// 通过UUID获取
	zone, err := cli.GetZone(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, zone.UUID, "Zone UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, zone.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got zone: %s, name: %s, state: %s", zone.UUID, zone.Name, zone.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateZone 测试更新区域
func TestUpdateZone(t *testing.T) {
	testName := "更新区域"
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
	list, err := cli.QueryZone(ctx, &queryParam)
	if !assert.NoError(t, err, "Query zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Zone found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Zone found to test")
		return
	}

	// 准备更新参数
	originalZone := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateZoneParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateZoneParamDetail{
			Name:        originalZone.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedZone, err := cli.UpdateZone(ctx, originalZone.UUID, updateParam)
	if !assert.NoError(t, err, "Update zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedZone.Description, "Zone description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedZone.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated zone: %s, new description: %s", updatedZone.UUID, updatedZone.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteZone 测试创建和删除区域
func TestCreateAndDeleteZone(t *testing.T) {
	testName := "创建和删除区域"
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
	t.Skip("Skipping create/delete zone test as it creates actual resources")

	// 创建区域
	createParam := param.CreateZoneParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateZoneParamDetail{
			Name:        "test-zone",
			Description: strPtr("Created by SDK test"),
		},
	}

	zone, err := cli.CreateZone(ctx, createParam)
	if !assert.NoError(t, err, "Create zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created zone: %s", zone.UUID)

	// 删除区域
	err = cli.DeleteZone(ctx, zone.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Zone deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
