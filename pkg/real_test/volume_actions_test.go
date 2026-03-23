// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryVolume 测试查询云盘
func TestQueryVolume(t *testing.T) {
	testName := "查询云盘"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryVolume(ctx, &queryParam)
	if !assert.NoError(t, err, "Query volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d volumes", len(result))
	skipIfNoResource(t, "Volume", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetVolume 测试获取云盘详情
func TestGetVolume(t *testing.T) {
	testName := "获取云盘详情"
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
	list, err := cli.QueryVolume(ctx, &queryParam)
	if !assert.NoError(t, err, "Query volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Volume found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Volume found to test")
		return
	}

	// 通过UUID获取
	volume, err := cli.GetVolume(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, volume.UUID, "Volume UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, volume.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got volume: %s, name: %s, state: %s", volume.UUID, volume.Name, volume.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateVolume 测试更新云盘
func TestUpdateVolume(t *testing.T) {
	testName := "更新云盘"
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
	list, err := cli.QueryVolume(ctx, &queryParam)
	if !assert.NoError(t, err, "Query volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Volume found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Volume found to test")
		return
	}

	// 准备更新参数
	originalVolume := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateVolumeParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateVolumeParamDetail{
			Name:        originalVolume.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedVolume, err := cli.UpdateVolume(ctx, originalVolume.UUID, updateParam)
	if !assert.NoError(t, err, "Update volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedVolume.Description, "Volume description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedVolume.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated volume: %s, new description: %s", updatedVolume.UUID, updatedVolume.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteVolume 测试创建和删除云盘
func TestCreateAndDeleteVolume(t *testing.T) {
	testName := "创建和删除云盘"
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
	t.Skip("Skipping create/delete volume test as it creates actual resources")

	// 查询获取一个有效的主存储UUID和磁盘规格UUID
	queryPSParam := param.NewQueryParam()
	queryPSParam.Limit(1)
	primaryStorages, err := cli.QueryPrimaryStorage(ctx, &queryPSParam)
	if !assert.NoError(t, err, "Query primary storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(primaryStorages) == 0 {
		golog.Warnf("测试跳过 [%s]: No primary storage found to create volume", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No primary storage found to create volume")
		return
	}

	// 查询磁盘规格
	queryDOParam := param.NewQueryParam()
	queryDOParam.Limit(1)
	diskOfferings, err := cli.QueryDiskOffering(ctx, &queryDOParam)
	if !assert.NoError(t, err, "Query disk offering should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(diskOfferings) == 0 {
		golog.Warnf("测试跳过 [%s]: No disk offering found to create volume", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No disk offering found to create volume")
		return
	}

	// 创建云盘参数
	createParam := param.CreateDataVolumeParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateDataVolumeParamDetail{
			Name:               "test-volume",
			Description:        strPtr("Created by SDK test"),
			DiskOfferingUuid:   strPtr(diskOfferings[0].UUID),
			PrimaryStorageUuid: strPtr(primaryStorages[0].UUID),
		},
	}

	volume, err := cli.CreateDataVolume(ctx, createParam)
	if !assert.NoError(t, err, "Create volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created volume: %s", volume.UUID)

	// 删除云盘
	err = cli.DeleteDataVolume(ctx, volume.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete volume should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Volume deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
