// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// GlobalConfigOptionsView GlobalConfigOptions
type GlobalConfigOptionsView struct {
	ValidValue []string `json:"validValue,omitempty"`
	NumberGreaterThan int64 `json:"numberGreaterThan,omitempty"`
	NumberLessThan int64 `json:"numberLessThan,omitempty"`
	NumberGreaterThanOrEqual int64 `json:"numberGreaterThanOrEqual,omitempty"`
	NumberLessThanOrEqual int64 `json:"numberLessThanOrEqual,omitempty"`
}

