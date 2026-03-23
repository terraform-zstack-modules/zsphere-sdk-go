// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// OvfCpuInfoView OvfCpuInfo
type OvfCpuInfoView struct {
	InstanceId string `json:"instanceId,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	CoresPerSocket int `json:"coresPerSocket,omitempty"`
}

