// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// UpdateBlockPrimaryStorageParamDetail UpdateBlockPrimaryStorage detail param
type UpdateBlockPrimaryStorageParamDetail struct {
	VendorName *string `json:"vendorName,omitempty"`
	Metadata *string `json:"metadata,omitempty"`
}

// UpdateBlockPrimaryStorageParam UpdateBlockPrimaryStorage request param
type UpdateBlockPrimaryStorageParam struct {
	BaseParam
	Params UpdateBlockPrimaryStorageParamDetail `json:"updateBlockPrimaryStorage"`
}
// AddBlockPrimaryStorageParamDetail AddBlockPrimaryStorage detail param
type AddBlockPrimaryStorageParamDetail struct {
	VendorName string `json:"vendorName" validate:"required"`
	Metadata string `json:"metadata" validate:"required"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddBlockPrimaryStorageParam AddBlockPrimaryStorage request param
type AddBlockPrimaryStorageParam struct {
	BaseParam
	Params AddBlockPrimaryStorageParamDetail `json:"param"`
}
