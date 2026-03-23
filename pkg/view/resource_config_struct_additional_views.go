// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceConfigStructView ResourceConfigStruct
type ResourceConfigStructView struct {
	Value string `json:"value,omitempty"`
	EffectiveConfigs []ResourceConfigInventoryView `json:"effectiveConfigs,omitempty"`
	Name string `json:"name,omitempty"`
}

