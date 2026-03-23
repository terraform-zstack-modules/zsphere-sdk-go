// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SecurityGroupL3NetworkRefInventoryView SecurityGroupL3NetworkRef
type SecurityGroupL3NetworkRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	L3NetworkUuid string `json:"l3NetworkUuid,omitempty"`
	SecurityGroupUuid string `json:"securityGroupUuid,omitempty"`
}

