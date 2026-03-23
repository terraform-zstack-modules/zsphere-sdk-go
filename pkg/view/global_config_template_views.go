// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GlobalConfigTemplateInventoryView GlobalConfigTemplate
type GlobalConfigTemplateInventoryView struct {
	BaseInfoView
	BaseTimeView
	Type string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
}

// QueryGlobalConfigTemplateView QueryGlobalConfigTemplate
type QueryGlobalConfigTemplateView struct {
	Inventories []GlobalConfigTemplateInventoryView `json:"inventories,omitempty"`
}

