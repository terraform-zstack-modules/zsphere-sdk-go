// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RefreshNvmeTargetParamDetail RefreshNvmeTarget detail param
type RefreshNvmeTargetParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	NvmeLunUuids []string `json:"nvmeLunUuids,omitempty"`
}

// RefreshNvmeTargetParam RefreshNvmeTarget request param
type RefreshNvmeTargetParam struct {
	BaseParam
	Params RefreshNvmeTargetParamDetail `json:"params"`
}
