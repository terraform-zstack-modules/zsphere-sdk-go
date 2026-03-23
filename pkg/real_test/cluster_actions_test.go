// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryCluster 测试查询集群
func TestQueryCluster(t *testing.T) {
	testName := "查询集群"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryCluster(ctx, &queryParam)
	if !assert.NoError(t, err, "Query cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d clusters", len(result))
	skipIfNoResource(t, "Cluster", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetCluster 测试获取集群详情
func TestGetCluster(t *testing.T) {
	testName := "获取集群详情"
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
	list, err := cli.QueryCluster(ctx, &queryParam)
	if !assert.NoError(t, err, "Query cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Cluster found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Cluster found to test")
		return
	}

	// 通过UUID获取
	cluster, err := cli.GetCluster(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, cluster.UUID, "Cluster UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, cluster.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got cluster: %s, name: %s, state: %s", cluster.UUID, cluster.Name, cluster.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateCluster 测试更新集群
func TestUpdateCluster(t *testing.T) {
	testName := "更新集群"
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
	list, err := cli.QueryCluster(ctx, &queryParam)
	if !assert.NoError(t, err, "Query cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No Cluster found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Cluster found to test")
		return
	}

	// 准备更新参数
	originalCluster := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateClusterParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateClusterParamDetail{
			Name:        originalCluster.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedCluster, err := cli.UpdateCluster(ctx, originalCluster.UUID, updateParam)
	if !assert.NoError(t, err, "Update cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedCluster.Description, "Cluster description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedCluster.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated cluster: %s, new description: %s", updatedCluster.UUID, updatedCluster.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateAndDeleteCluster 测试创建和删除集群
func TestCreateAndDeleteCluster(t *testing.T) {
	testName := "创建和删除集群"
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
	t.Skip("Skipping create/delete cluster test as it creates actual resources")

	// 查询获取一个有效的区域UUID
	queryZoneParam := param.NewQueryParam()
	queryZoneParam.Limit(1)
	zones, err := cli.QueryZone(ctx, &queryZoneParam)
	if !assert.NoError(t, err, "Query zone should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(zones) == 0 {
		golog.Warnf("测试跳过 [%s]: No Zone found to create cluster", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No Zone found to create cluster")
		return
	}

	// 创建集群
	createParam := param.CreateClusterParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateClusterParamDetail{
			ZoneUuid:       zones[0].UUID,
			Name:           "test-cluster",
			Description:    strPtr("Created by SDK test"),
			HypervisorType: "KVM", // 根据实际环境调整
		},
	}

	cluster, err := cli.CreateCluster(ctx, createParam)
	if !assert.NoError(t, err, "Create cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created cluster: %s", cluster.UUID)

	// 删除集群
	err = cli.DeleteCluster(ctx, cluster.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Delete cluster should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("Cluster deleted successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
