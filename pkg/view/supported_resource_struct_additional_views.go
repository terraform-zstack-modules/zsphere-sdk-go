// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// SupportedResourceStructView SupportedResourceStruct
type SupportedResourceStructView struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	ActionName string `json:"actionName,omitempty"`
	Resources []string `json:"resources,omitempty"`
}

