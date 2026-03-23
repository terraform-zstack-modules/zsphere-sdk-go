// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceBindableConfigStructView ResourceBindableConfigStruct
type ResourceBindableConfigStructView struct {
	Name string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	BindResourceTypes []string `json:"bindResourceTypes,omitempty"`
}

