// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ElaborationCategoryView ElaborationCategory
type ElaborationCategoryView struct {
	Category string `json:"category,omitempty"`
	Num int `json:"num,omitempty"`
}

