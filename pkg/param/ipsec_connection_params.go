// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// ReconnectIPsecConnectionParamDetail ReconnectIPsecConnection detail param
type ReconnectIPsecConnectionParamDetail struct {
}

// ReconnectIPsecConnectionParam ReconnectIPsecConnection request param
type ReconnectIPsecConnectionParam struct {
	BaseParam
	Params ReconnectIPsecConnectionParamDetail `json:"reconnectIPsecConnection"`
}
// ChangeIPsecConnectionParamDetail ChangeIPsecConnection detail param
type ChangeIPsecConnectionParamDetail struct {
	PeerAddress string `json:"peerAddress" validate:"required"`
	AuthMode *string `json:"authMode,omitempty"`
	AuthKey string `json:"authKey" validate:"required"`
	IdType *string `json:"idType,omitempty"`
	LocalId *string `json:"localId,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	IkeVersion *string `json:"ikeVersion,omitempty"`
	IkeAuthAlgorithm *string `json:"ikeAuthAlgorithm,omitempty"`
	IkeEncryptionAlgorithm *string `json:"ikeEncryptionAlgorithm,omitempty"`
	IkeDhGroup int `json:"ikeDhGroup,omitempty"`
	IkeLifeTime *int `json:"ikeLifeTime,omitempty"`
	PolicyAuthAlgorithm *string `json:"policyAuthAlgorithm,omitempty"`
	PolicyEncryptionAlgorithm *string `json:"policyEncryptionAlgorithm,omitempty"`
	Pfs *string `json:"pfs,omitempty"`
	PolicyMode *string `json:"policyMode,omitempty"`
	TransformProtocol *string `json:"transformProtocol,omitempty"`
	LifeTime *int `json:"lifeTime,omitempty"`
}

// ChangeIPsecConnectionParam ChangeIPsecConnection request param
type ChangeIPsecConnectionParam struct {
	BaseParam
	Params ChangeIPsecConnectionParamDetail `json:"changeIPsecConnection"`
}
// UpdateIPsecConnectionParamDetail UpdateIPsecConnection detail param
type UpdateIPsecConnectionParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// UpdateIPsecConnectionParam UpdateIPsecConnection request param
type UpdateIPsecConnectionParam struct {
	BaseParam
	Params UpdateIPsecConnectionParamDetail `json:"updateIPsecConnection"`
}
// DeleteIPsecConnectionParamDetail DeleteIPsecConnection detail param
type DeleteIPsecConnectionParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteIPsecConnectionParam DeleteIPsecConnection request param
type DeleteIPsecConnectionParam struct {
	BaseParam
	Params DeleteIPsecConnectionParamDetail `json:"deleteIPsecConnection"`
}
// CreateIPsecConnectionParamDetail CreateIPsecConnection detail param
type CreateIPsecConnectionParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	VipUuid string `json:"vipUuid" validate:"required"`
	PeerAddress string `json:"peerAddress" validate:"required"`
	AuthMode *string `json:"authMode,omitempty"`
	AuthKey string `json:"authKey" validate:"required"`
	IdType *string `json:"idType,omitempty"`
	LocalId *string `json:"localId,omitempty"`
	RemoteId *string `json:"remoteId,omitempty"`
	L3NetworkUuid *string `json:"l3NetworkUuid,omitempty"`
	PeerCidrs []string `json:"peerCidrs,omitempty"`
	IkeVersion *string `json:"ikeVersion,omitempty"`
	IkeAuthAlgorithm *string `json:"ikeAuthAlgorithm,omitempty"`
	IkeEncryptionAlgorithm *string `json:"ikeEncryptionAlgorithm,omitempty"`
	IkeDhGroup int `json:"ikeDhGroup,omitempty"`
	IkeLifeTime *int `json:"ikeLifeTime,omitempty"`
	PolicyAuthAlgorithm *string `json:"policyAuthAlgorithm,omitempty"`
	PolicyEncryptionAlgorithm *string `json:"policyEncryptionAlgorithm,omitempty"`
	Pfs *string `json:"pfs,omitempty"`
	PolicyMode *string `json:"policyMode,omitempty"`
	TransformProtocol *string `json:"transformProtocol,omitempty"`
	LifeTime *int `json:"lifeTime,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateIPsecConnectionParam CreateIPsecConnection request param
type CreateIPsecConnectionParam struct {
	BaseParam
	Params CreateIPsecConnectionParamDetail `json:"params"`
}
