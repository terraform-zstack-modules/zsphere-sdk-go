// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// RefreshFiberChannelStorageParamDetail RefreshFiberChannelStorage detail param
type RefreshFiberChannelStorageParamDetail struct {
	ZoneUuid string `json:"zoneUuid" validate:"required"`
	ScsiLunUuids []string `json:"scsiLunUuids,omitempty"`
}

// RefreshFiberChannelStorageParam RefreshFiberChannelStorage request param
type RefreshFiberChannelStorageParam struct {
	BaseParam
	Params RefreshFiberChannelStorageParamDetail `json:"params"`
}
