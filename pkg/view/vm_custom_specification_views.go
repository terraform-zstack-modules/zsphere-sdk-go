// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// VmCustomSpecificationInventoryView VmCustomSpecification
type VmCustomSpecificationInventoryView struct {
	BaseInfoView
	BaseTimeView
	VmInstanceUuid string `json:"vmInstanceUuid,omitempty"`
	Description string `json:"description,omitempty"`
	Platform string `json:"platform,omitempty"`
	Hostname string `json:"hostname,omitempty"`
	GenerateSID bool `json:"generateSID,omitempty"`
	DomainMode string `json:"domainMode,omitempty"`
	DomainName string `json:"domainName,omitempty"`
	DomainUsername string `json:"domainUsername,omitempty"`
	Organization string `json:"organization,omitempty"`
}

// QueryVmCustomSpecificationView QueryVmCustomSpecification
type QueryVmCustomSpecificationView struct {
	Inventories []VmCustomSpecificationInventoryView `json:"inventories,omitempty"`
}

// DeleteVmCustomSpecificationEventView DeleteVmCustomSpecificationEvent
type DeleteVmCustomSpecificationEventView struct {
	Success bool `json:"success,omitempty"`
}

// UpdateVmCustomSpecificationEventView UpdateVmCustomSpecificationEvent
type UpdateVmCustomSpecificationEventView struct {
	Inventory VmCustomSpecificationInventoryView `json:"inventory,omitempty"`
}

// CreateVmCustomSpecificationEventView CreateVmCustomSpecificationEvent
type CreateVmCustomSpecificationEventView struct {
	Inventory VmCustomSpecificationInventoryView `json:"inventory,omitempty"`
}

