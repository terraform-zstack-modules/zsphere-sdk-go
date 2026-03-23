// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteImageReplicationGroupParamDetail DeleteImageReplicationGroup detail param
type DeleteImageReplicationGroupParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteImageReplicationGroupParam DeleteImageReplicationGroup request param
type DeleteImageReplicationGroupParam struct {
	BaseParam
	Params DeleteImageReplicationGroupParamDetail `json:"deleteImageReplicationGroup"`
}
// CreateImageReplicationGroupParamDetail CreateImageReplicationGroup detail param
type CreateImageReplicationGroupParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateImageReplicationGroupParam CreateImageReplicationGroup request param
type CreateImageReplicationGroupParam struct {
	BaseParam
	Params CreateImageReplicationGroupParamDetail `json:"params"`
}
