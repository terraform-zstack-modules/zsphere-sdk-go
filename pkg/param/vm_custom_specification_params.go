// Copyright (c) ZStack.io, Inc.

package param

import "time"

var _ = time.Now() // avoid unused import

// DeleteVmCustomSpecificationParamDetail DeleteVmCustomSpecification detail param
type DeleteVmCustomSpecificationParamDetail struct {
	DeleteMode *string `json:"deleteMode,omitempty"`
}

// DeleteVmCustomSpecificationParam DeleteVmCustomSpecification request param
type DeleteVmCustomSpecificationParam struct {
	BaseParam
	Params DeleteVmCustomSpecificationParamDetail `json:"deleteVmCustomSpecification"`
}
// UpdateVmCustomSpecificationParamDetail UpdateVmCustomSpecification detail param
type UpdateVmCustomSpecificationParamDetail struct {
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Hostname *string `json:"hostname,omitempty"`
	RootPassword *string `json:"rootPassword,omitempty"`
	GenerateSID *bool `json:"generateSID,omitempty"`
	DomainMode *string `json:"domainMode,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	DomainUsername *string `json:"domainUsername,omitempty"`
	DomainPassword *string `json:"domainPassword,omitempty"`
	Organization *string `json:"organization,omitempty"`
}

// UpdateVmCustomSpecificationParam UpdateVmCustomSpecification request param
type UpdateVmCustomSpecificationParam struct {
	BaseParam
	Params UpdateVmCustomSpecificationParamDetail `json:"updateVmCustomSpecification"`
}
// CreateVmCustomSpecificationParamDetail CreateVmCustomSpecification detail param
type CreateVmCustomSpecificationParamDetail struct {
	Name string `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
	Platform string `json:"platform" validate:"required"`
	Hostname *string `json:"hostname,omitempty"`
	RootPassword *string `json:"rootPassword,omitempty"`
	GenerateSID *bool `json:"generateSID,omitempty"`
	DomainMode *string `json:"domainMode,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
	DomainUsername *string `json:"domainUsername,omitempty"`
	DomainPassword *string `json:"domainPassword,omitempty"`
	Organization *string `json:"organization,omitempty"`
	ResourceUuid *string `json:"resourceUuid,omitempty"`
	TagUuids []string `json:"tagUuids,omitempty"`
}

// CreateVmCustomSpecificationParam CreateVmCustomSpecification request param
type CreateVmCustomSpecificationParam struct {
	BaseParam
	Params CreateVmCustomSpecificationParamDetail `json:"params"`
}
