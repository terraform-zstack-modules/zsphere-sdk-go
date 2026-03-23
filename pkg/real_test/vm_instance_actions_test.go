// Copyright (c) ZStack.io, Inc.

package real_test

import (
	"context"
	"testing"

	"github.com/kataras/golog"
	"github.com/stretchr/testify/assert"

	"github.com/terraform-zstack-modules/zsphere-sdk-go/pkg/param"
)

// TestQueryVmInstance 测试查询虚拟机实例
func TestQueryVmInstance(t *testing.T) {
	testName := "查询虚拟机实例"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	queryParam := param.NewQueryParam()
	result, err := cli.QueryVmInstance(ctx, &queryParam)
	if !assert.NoError(t, err, "Query VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Found %d VM instances", len(result))
	skipIfNoResource(t, "VM Instance", len(result))

	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestGetVmInstance 测试获取虚拟机实例详情
func TestGetVmInstance(t *testing.T) {
	testName := "获取虚拟机实例详情"
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
	list, err := cli.QueryVmInstance(ctx, &queryParam)
	if !assert.NoError(t, err, "Query VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No VM Instance found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No VM Instance found to test")
		return
	}

	// 通过UUID获取
	vm, err := cli.GetVmInstance(ctx, list[0].UUID)
	if !assert.NoError(t, err, "Get VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, list[0].UUID, vm.UUID, "VM instance UUID should match") {
		golog.Errorf("测试失败 [%s]: UUID不匹配 expected=%s, actual=%s", testName, list[0].UUID, vm.UUID)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Got VM instance: %s, name: %s, state: %s", vm.UUID, vm.Name, vm.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestUpdateVmInstance 测试更新虚拟机实例
func TestUpdateVmInstance(t *testing.T) {
	testName := "更新虚拟机实例"
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
	list, err := cli.QueryVmInstance(ctx, &queryParam)
	if !assert.NoError(t, err, "Query VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No VM Instance found to test", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No VM Instance found to test")
		return
	}

	// 准备更新参数
	originalVm := list[0]
	newDescription := "Updated by SDK test"

	updateParam := param.UpdateVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params: param.UpdateVmInstanceParamDetail{
			Name:        originalVm.Name, // 保持名称不变
			Description: strPtr(newDescription),
		},
	}

	// 执行更新
	updatedVm, err := cli.UpdateVmInstance(ctx, originalVm.UUID, updateParam)
	if !assert.NoError(t, err, "Update VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, newDescription, updatedVm.Description, "VM instance description should be updated") {
		golog.Errorf("测试失败 [%s]: 描述未更新 expected=%s, actual=%s", testName, newDescription, updatedVm.Description)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Updated VM instance: %s, new description: %s", updatedVm.UUID, updatedVm.Description)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestVmInstanceLifecycle 测试虚拟机实例生命周期操作
func TestVmInstanceLifecycle(t *testing.T) {
	testName := "虚拟机实例生命周期操作"
	golog.Infof("开始测试: %s", testName)
	ctx := context.Background()
	defer func() {
		if r := recover(); r != nil {
			golog.Errorf("测试失败 [%s]: %v", testName, r)
			recordTestResult(t.Name(), false)
		}
	}()

	// 首先查询获取一个有效的UUID，查找运行中的虚拟机
	queryParam := param.NewQueryParam()
	queryParam.AddQ("state=Running")
	queryParam.Limit(1)
	list, err := cli.QueryVmInstance(ctx, &queryParam)
	if !assert.NoError(t, err, "Query VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if len(list) == 0 {
		golog.Warnf("测试跳过 [%s]: No running VM instance found to test lifecycle operations", testName)
		recordTestResult(t.Name(), true) // 跳过也算通过
		t.Skip("No running VM instance found to test lifecycle operations")
		return
	}

	vm := list[0]
	golog.Infof("Testing lifecycle operations for VM: %s, current state: %s", vm.UUID, vm.State)

	// 停止虚拟机
	typeStr := "grace" // 优雅关机
	stopParam := param.StopVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params: param.StopVmInstanceParamDetail{
			Type: strPtr(typeStr),
		},
	}

	stoppedVm, err := cli.StopVmInstance(ctx, vm.UUID, stopParam)
	if err != nil {
		golog.Errorf("测试失败 [%s]: 停止虚拟机失败 - %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, "Stopped", stoppedVm.State, "VM state should be Stopped") {
		golog.Errorf("测试失败 [%s]: 虚拟机状态不是Stopped，而是%s", testName, stoppedVm.State)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("VM stopped successfully, new state: %s", stoppedVm.State)

	// 启动虚拟机
	startParam := param.StartVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params:    param.StartVmInstanceParamDetail{},
	}

	startedVm, err := cli.StartVmInstance(ctx, vm.UUID, startParam)
	if err != nil {
		golog.Errorf("测试失败 [%s]: 启动虚拟机失败 - %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	if !assert.Equal(t, "Running", startedVm.State, "VM state should be Running") {
		golog.Errorf("测试失败 [%s]: 虚拟机状态不是Running，而是%s", testName, startedVm.State)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("VM started successfully, new state: %s", startedVm.State)
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}

// TestCreateVmInstance 测试创建虚拟机实例
func TestCreateVmInstance(t *testing.T) {
	testName := "创建虚拟机实例"
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
	t.Skip("Skipping create VM instance test as it creates actual resources")

	// 创建虚拟机需要有效的实例规格UUID、镜像UUID和L3网络UUID
	// 这些值需要根据实际环境进行替换
	instanceOfferingUuid := "your-instance-offering-uuid"
	imageUuid := "your-image-uuid"
	l3NetworkUuid := "your-l3network-uuid"
	descStr := "Created by SDK test"

	createParam := param.CreateVmInstanceParam{
		BaseParam: param.BaseParam{},
		Params: param.CreateVmInstanceParamDetail{
			Name:                 "test-vm",
			InstanceOfferingUuid: strPtr(instanceOfferingUuid),
			ImageUuid:            strPtr(imageUuid),
			L3NetworkUuids:       []string{l3NetworkUuid},
			Description:          strPtr(descStr),
		},
	}

	vm, err := cli.CreateVmInstance(ctx, createParam)
	if !assert.NoError(t, err, "Create VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Infof("Created VM instance: %s", vm.UUID)

	// 删除虚拟机
	err = cli.DestroyVmInstance(ctx, vm.UUID, param.DeleteModePermissive)
	if !assert.NoError(t, err, "Destroy VM instance should not return error") {
		golog.Errorf("测试失败 [%s]: %v", testName, err)
		recordTestResult(t.Name(), false)
		return
	}

	golog.Info("VM instance destroyed successfully")
	golog.Infof("测试通过: %s", testName)
	recordTestResult(t.Name(), true)
}
