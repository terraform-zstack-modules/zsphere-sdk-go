// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddExternalBackupStorageParamDetail AddExternalBackupStorage detail param
type AddExternalBackupStorageParamDetail struct {
	Identity string `json:"identity" validate:"required"`
	Url string `json:"url" validate:"required"`
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	ImportImages *bool `json:"importImages,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddExternalBackupStorageParam AddExternalBackupStorage request param
type AddExternalBackupStorageParam struct {
	BaseParam
	Params AddExternalBackupStorageParamDetail `json:"params"`
}
