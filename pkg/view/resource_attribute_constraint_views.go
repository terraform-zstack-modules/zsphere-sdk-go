// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ResourceAttributeConstraintInventoryView ResourceAttributeConstraint
type ResourceAttributeConstraintInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id int64 `json:"id,omitempty"`
	KeyUuid string `json:"keyUuid,omitempty"`
	Type string `json:"type,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

