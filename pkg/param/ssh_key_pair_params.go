// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// CreateSshKeyPairParamDetail CreateSshKeyPair detail param
type CreateSshKeyPairParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	PublicKey string `json:"publicKey" validate:"required"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateSshKeyPairParam CreateSshKeyPair request param
type CreateSshKeyPairParam struct {
	BaseParam
	Params CreateSshKeyPairParamDetail `json:"params"`
}
// GenerateSshKeyPairParamDetail GenerateSshKeyPair detail param
type GenerateSshKeyPairParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// GenerateSshKeyPairParam GenerateSshKeyPair request param
type GenerateSshKeyPairParam struct {
	BaseParam
	Params GenerateSshKeyPairParamDetail `json:"params"`
}
// UpdateSshKeyPairParamDetail UpdateSshKeyPair detail param
type UpdateSshKeyPairParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateSshKeyPairParam UpdateSshKeyPair request param
type UpdateSshKeyPairParam struct {
	BaseParam
	Params UpdateSshKeyPairParamDetail `json:"updateSshKeyPair"`
}
// DeleteSshKeyPairParamDetail DeleteSshKeyPair detail param
type DeleteSshKeyPairParamDetail struct {
}

// DeleteSshKeyPairParam DeleteSshKeyPair request param
type DeleteSshKeyPairParam struct {
	BaseParam
	Params DeleteSshKeyPairParamDetail `json:"deleteSshKeyPair"`
}
