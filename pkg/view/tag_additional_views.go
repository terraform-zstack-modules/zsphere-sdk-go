// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// TagView Tag
type TagView struct {
	Name string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

