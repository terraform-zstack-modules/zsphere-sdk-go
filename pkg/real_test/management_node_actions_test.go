// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryManagementNode 测试查询管理节点
func TestQueryManagementNode(t *testing.T) {
	testName := "查询管理节点"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryManagementNode(ctx, &queryParam)
	if !assert.NoError(t, err, "Query management node should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d management nodes", len(result))
	skipIfNoResource(t, "ManagementNode", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetManagementNode 测试获取管理节点详情
func TestGetManagementNode(t *testing.T) {
	testName := "获取管理节点详情"
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
	list, err := cli.QueryManagementNode(ctx, &queryParam)
	if !assert.NoError(t, err, "Query management node should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No ManagementNode found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No ManagementNode found to test")
		return
	}

	// 通过UUID获取
	managementNode, err := cli.GetManagementNode(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get management node should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, managementNode.UUID, "ManagementNode UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, managementNode.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got management node: %s, hostname: %s", managementNode.UUID, managementNode.HostName)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestManagementNodePageQuery 测试管理节点的分页查询
func TestManagementNodePageQuery(t *testing.T) {
	testName := "管理节点的分页查询"
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

	managementNodes, total, err := cli.PageManagementNode(ctx, &queryParam)
	if !assert.NoError(t, err, "Page management node should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Paged %d management nodes out of %d total", len(managementNodes), total)

	if !assert.LessOrEqual(t, len(managementNodes), pageSize, "Page size should be respected") {
		golog.Errorf("测试失败 [%s]: 返回的管理节点数量超过了页面大小 expected<=%d, actual=%d", testName, pageSize, len(managementNodes))
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
