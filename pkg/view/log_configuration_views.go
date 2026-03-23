// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// LogConfigurationInventoryView LogConfiguration
type LogConfigurationInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	ManagementNodeUuids []string `json:"managementNodeUuids,omitempty"`
}

// DeleteLogConfigurationEventView DeleteLogConfigurationEvent
type DeleteLogConfigurationEventView struct {
	Success bool `json:"success,omitempty"`
}

