// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RemoveAccessControlListEntryParamDetail RemoveAccessControlListEntry detail param
type RemoveAccessControlListEntryParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// RemoveAccessControlListEntryParam RemoveAccessControlListEntry request param
type RemoveAccessControlListEntryParam struct {
	BaseParam
	Params RemoveAccessControlListEntryParamDetail `json:"removeAccessControlListEntry"`
}
// AddAccessControlListEntryParamDetail AddAccessControlListEntry detail param
type AddAccessControlListEntryParamDetail struct {
	Entries string `json:"entries" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddAccessControlListEntryParam AddAccessControlListEntry request param
type AddAccessControlListEntryParam struct {
	BaseParam
	Params AddAccessControlListEntryParamDetail `json:"params"`
}
