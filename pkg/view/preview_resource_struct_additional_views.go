// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// PreviewResourceStructView PreviewResourceStruct
type PreviewResourceStructView struct {
	Actions []ActionStructView `json:"actions,omitempty"`
	Conditions map[string]bool `json:"conditions,omitempty"`
}

