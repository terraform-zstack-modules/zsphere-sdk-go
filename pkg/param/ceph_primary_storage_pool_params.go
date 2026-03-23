// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// AddCephPrimaryStoragePoolParamDetail AddCephPrimaryStoragePool detail param
type AddCephPrimaryStoragePoolParamDetail struct {
	PoolName string `json:"poolName" validate:"required"`
	AliasName *string `json:"aliasName,omitempty"`
	Description *string `json:"description,omitempty"`
	Type string `json:"type" validate:"required"`
	IsCreate bool `json:"isCreate,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// AddCephPrimaryStoragePoolParam AddCephPrimaryStoragePool request param
type AddCephPrimaryStoragePoolParam struct {
	BaseParam
	Params AddCephPrimaryStoragePoolParamDetail `json:"params"`
}
// DeleteCephPrimaryStoragePoolParamDetail DeleteCephPrimaryStoragePool detail param
type DeleteCephPrimaryStoragePoolParamDetail struct {
}

// DeleteCephPrimaryStoragePoolParam DeleteCephPrimaryStoragePool request param
type DeleteCephPrimaryStoragePoolParam struct {
	BaseParam
	Params DeleteCephPrimaryStoragePoolParamDetail `json:"deleteCephPrimaryStoragePool"`
}
// UpdateCephPrimaryStoragePoolParamDetail UpdateCephPrimaryStoragePool detail param
type UpdateCephPrimaryStoragePoolParamDetail struct {
	AliasName *string `json:"aliasName,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateCephPrimaryStoragePoolParam UpdateCephPrimaryStoragePool request param
type UpdateCephPrimaryStoragePoolParam struct {
	BaseParam
	Params UpdateCephPrimaryStoragePoolParamDetail `json:"updateCephPrimaryStoragePool"`
}
