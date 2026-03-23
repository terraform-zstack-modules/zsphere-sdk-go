// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// UserTagInventoryView UserTag
type UserTagInventoryView struct {
	BaseInfoView
	BaseTimeView
	TagPatternUuid string `json:"tagPatternUuid,omitempty"`
	TagPattern TagPatternInventoryView `json:"tagPattern,omitempty"`
	ResourceUuid string `json:"resourceUuid,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Tag string `json:"tag,omitempty"`
	Type string `json:"type,omitempty"`
}

// QueryUserTagView QueryUserTag
type QueryUserTagView struct {
	Inventories []UserTagInventoryView `json:"inventories,omitempty"`
}

