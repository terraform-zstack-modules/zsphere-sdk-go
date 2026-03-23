// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// IPsecConnectionInventoryView IPsecConnection
type IPsecConnectionInventoryView struct {
	BaseInfoView
	BaseTimeView
	Description string `json:"description,omitempty"`
	PeerAddress string `json:"peerAddress,omitempty"`
	AuthMode string `json:"authMode,omitempty"`
	AuthKey string `json:"authKey,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	IkeAuthAlgorithm string `json:"ikeAuthAlgorithm,omitempty"`
	IkeEncryptionAlgorithm string `json:"ikeEncryptionAlgorithm,omitempty"`
	IkeDhGroup int `json:"ikeDhGroup,omitempty"`
	PolicyAuthAlgorithm string `json:"policyAuthAlgorithm,omitempty"`
	PolicyEncryptionAlgorithm string `json:"policyEncryptionAlgorithm,omitempty"`
	Pfs string `json:"pfs,omitempty"`
	PolicyMode string `json:"policyMode,omitempty"`
	TransformProtocol string `json:"transformProtocol,omitempty"`
	IkeVersion string `json:"ikeVersion,omitempty"`
	IdType string `json:"idType,omitempty"`
	LocalId string `json:"localId,omitempty"`
	RemoteId string `json:"remoteId,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	IkeLifeTime int `json:"ikeLifeTime,omitempty"`
	LifeTime int `json:"lifeTime,omitempty"`
	PeerCidrs []IPsecPeerCidrInventoryView `json:"peerCidrs,omitempty"`
	L3NetworkRefs []IPsecL3NetworkRefInventoryView `json:"l3NetworkRefs,omitempty"`
}

// ReconnectIPsecConnectionEventView ReconnectIPsecConnectionEvent
type ReconnectIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// AddRemoteCidrsToIPsecConnectionEventView AddRemoteCidrsToIPsecConnectionEvent
type AddRemoteCidrsToIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// QueryIPSecConnectionView QueryIPSecConnection
type QueryIPSecConnectionView struct {
	Inventories []IPsecConnectionInventoryView `json:"inventories,omitempty"`
}

// ChangeIPsecConnectionEventView ChangeIPsecConnectionEvent
type ChangeIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// ChangeIPSecConnectionStateEventView ChangeIPSecConnectionStateEvent
type ChangeIPSecConnectionStateEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// DetachL3NetworksFromIPsecConnectionEventView DetachL3NetworksFromIPsecConnectionEvent
type DetachL3NetworksFromIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// AttachL3NetworksToIPsecConnectionEventView AttachL3NetworksToIPsecConnectionEvent
type AttachL3NetworksToIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// UpdateIPsecConnectionEventView UpdateIPsecConnectionEvent
type UpdateIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// RemoveRemoteCidrsFromIPsecConnectionEventView RemoveRemoteCidrsFromIPsecConnectionEvent
type RemoveRemoteCidrsFromIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

// DeleteIPsecConnectionEventView DeleteIPsecConnectionEvent
type DeleteIPsecConnectionEventView struct {
	Success bool `json:"success,omitempty"`
}

// GetVpcAttachedIpsecView GetVpcAttachedIpsec
type GetVpcAttachedIpsecView struct {
	Inventories []IPsecConnectionInventoryView `json:"inventories,omitempty"`
}

// CreateIPsecConnectionEventView CreateIPsecConnectionEvent
type CreateIPsecConnectionEventView struct {
	Inventory IPsecConnectionInventoryView `json:"inventory,omitempty"`
}

