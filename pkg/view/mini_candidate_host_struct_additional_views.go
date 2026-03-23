// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// MiniCandidateHostStructView MiniCandidateHostStruct
type MiniCandidateHostStructView struct {
	HostName string `json:"hostName,omitempty"`
	Ipv4Address string `json:"ipv4Address,omitempty"`
	Ipv6Address string `json:"ipv6Address,omitempty"`
	Ipv4Interface string `json:"ipv4Interface,omitempty"`
	Ipv4CidrPrefix string `json:"ipv4CidrPrefix,omitempty"`
	Ipv4InterfaceBond string `json:"ipv4InterfaceBond,omitempty"`
	Ipv4Gateway string `json:"ipv4Gateway,omitempty"`
	ManagementVip string `json:"managementVip,omitempty"`
	IpmiIpv4Addr string `json:"ipmiIpv4Addr,omitempty"`
	IpmiIpv4Gateway string `json:"ipmiIpv4Gateway,omitempty"`
	IpmiVlan string `json:"ipmiVlan,omitempty"`
	Ipv6Interface string `json:"ipv6Interface,omitempty"`
	Manufacturer string `json:"manufacturer,omitempty"`
	Product string `json:"product,omitempty"`
	Sn string `json:"sn,omitempty"`
}

