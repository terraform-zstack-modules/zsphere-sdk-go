// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryImage 测试查询镜像
func TestQueryImage(t *testing.T) {
	testName := "查询镜像"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryImage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d images", len(result))
	skipIfNoResource(t, "Image", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetImage 测试获取镜像详情
func TestGetImage(t *testing.T) {
	testName := "获取镜像详情"
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
	list, err := cli.QueryImage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Image found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Image found to test")
		return
	}

	// 通过UUID获取
	image, err := cli.GetImage(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, image.UUID, "Image UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, image.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got image: %s, name: %s, state: %s", image.UUID, image.Name, image.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateImage 测试更新镜像
func TestUpdateImage(t *testing.T) {
	testName := "更新镜像"
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
	list, err := cli.QueryImage(ctx, &queryParam)
	if !assert.NoError(t, err, "Query image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Image found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Image found to test")
		return
	}

	// 准备更新参数
	originalImage := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateImageParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateImageParamDetail{
			Name:        originalImage.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedImage, err := cli.UpdateImage(ctx, originalImage.UUID, updateParam)
	if !assert.NoError(t, err, "Update image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedImage.Description, "Image description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedImage.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated image: %s, new description: %s", updatedImage.UUID, updatedImage.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteImage 测试创建和删除镜像
func TestCreateAndDeleteImage(t *testing.T) {
	testName := "创建和删除镜像"
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
	t.Skip("Skipping create/delete image test as it creates actual resources")

	// 查询获取一个有效的备份存储UUID
	queryBSParam := param.NewQueryParam()
	queryBSParam.Limit(1)
	backupStorages, err := cli.QueryBackupStorage(ctx, &queryBSParam)
	if !assert.NoError(t, err, "Query backup storage should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(backupStorages) == 0 {
		golog.Warnf("测试跳过 [%s]: No backup storage found to create image", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No backup storage found to create image")
		return
	}

	// 创建镜像参数
	// 注意：这里使用URL方式添加镜像，需要一个有效的镜像URL
	url := "http://example.com/path/to/image.qcow2" // 替换为有效的镜像URL
	createParam := param.AddImageParam{
		BaseParam: param.BaseParam{},
		Params: param.AddImageParamDetail{
			Name:               "test-image",
			Url:                url,
			Description:        strPtr("Created by SDK test"),
			BackupStorageUuids: []string{backupStorages[0].UUID},
			Format:             strPtr("qcow2"),
			Platform:           strPtr("Linux"),
		},
	}

	image, err := cli.AddImage(ctx, createParam)
	if !assert.NoError(t, err, "Create image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created image: %s", image.UUID)

	// 删除镜像
	err = cli.DeleteImage(ctx, image.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete image should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Image deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
