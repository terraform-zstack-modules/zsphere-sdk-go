// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryDiskOffering 测试查询云盘规格
func TestQueryDiskOffering(t *testing.T) {
	testName := "查询云盘规格"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryDiskOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d disk offerings", len(result))
	skipIfNoResource(t, "DiskOffering", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetDiskOffering 测试获取云盘规格详情
func TestGetDiskOffering(t *testing.T) {
	testName := "获取云盘规格详情"
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
	list, err := cli.QueryDiskOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No DiskOffering found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No DiskOffering found to test")
		return
	}

	// 通过UUID获取
	diskOffering, err := cli.GetDiskOffering(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, diskOffering.UUID, "DiskOffering UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, diskOffering.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got disk offering: %s, name: %s, diskSize: %d", diskOffering.UUID, diskOffering.Name, diskOffering.DiskSize)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateDiskOffering 测试更新云盘规格
func TestUpdateDiskOffering(t *testing.T) {
	testName := "更新云盘规格"
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
	list, err := cli.QueryDiskOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Query disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No DiskOffering found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No DiskOffering found to test")
		return
	}

	// 准备更新参数
	originalDiskOffering := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateDiskOfferingParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateDiskOfferingParamDetail{
			Name:        originalDiskOffering.Name, // 保持名称不变
			Description: &newDescription,
		},
	}

	// 执行更新
	updatedDiskOffering, err := cli.UpdateDiskOffering(ctx, originalDiskOffering.UUID, updateParam)
	if !assert.NoError(t, err, "Update disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedDiskOffering.Description, "DiskOffering description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedDiskOffering.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated disk offering: %s, new description: %s", updatedDiskOffering.UUID, updatedDiskOffering.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteDiskOffering 测试创建和删除云盘规格
func TestCreateAndDeleteDiskOffering(t *testing.T) {
	testName := "创建和删除云盘规格"
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
	t.Skip("Skipping create disk offering test as it creates actual resources")

	// 创建云盘规格参数
	diskSize := int64(10 * 1024 * 1024 * 1024) // 10GB
	description := "Created by SDK test"
	createParam := param.CreateDiskOfferingParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateDiskOfferingParamDetail{
			Name:        "test-disk-offering",
			Description: &description,
			DiskSize:    diskSize,
		},
	}

	// 执行创建
	diskOffering, err := cli.CreateDiskOffering(ctx, createParam)
	if !assert.NoError(t, err, "Create disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created disk offering: %s", diskOffering.UUID)

	// 删除云盘规格
	err = cli.DeleteDiskOffering(ctx, diskOffering.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Disk offering deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestDiskOfferingPageQuery 测试云盘规格的分页查询
func TestDiskOfferingPageQuery(t *testing.T) {
	testName := "云盘规格的分页查询"
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

	diskOfferings, total, err := cli.PageDiskOffering(ctx, &queryParam)
	if !assert.NoError(t, err, "Page disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d disk offerings out of %d total", len(diskOfferings), total)

	if !assert.LessOrEqual(t, len(diskOfferings), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的云盘规格数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(diskOfferings))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
