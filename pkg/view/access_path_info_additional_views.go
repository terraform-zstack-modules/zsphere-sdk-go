// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// AccessPathInfoView AccessPathInfo
type AccessPathInfoView struct {
	Name string `json:"name,omitempty"`
	AccessPathId int `json:"accessPathId,omitempty"`
	AccessPathIqn string `json:"accessPathIqn,omitempty"`
	TargetCount int `json:"targetCount,omitempty"`
	GatewayIps []string `json:"gatewayIps,omitempty"`
}

