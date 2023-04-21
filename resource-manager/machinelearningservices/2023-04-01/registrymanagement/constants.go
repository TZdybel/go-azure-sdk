package registrymanagement

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointServiceConnectionStatus string

const (
	EndpointServiceConnectionStatusApproved     EndpointServiceConnectionStatus = "Approved"
	EndpointServiceConnectionStatusDisconnected EndpointServiceConnectionStatus = "Disconnected"
	EndpointServiceConnectionStatusPending      EndpointServiceConnectionStatus = "Pending"
	EndpointServiceConnectionStatusRejected     EndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForEndpointServiceConnectionStatus() []string {
	return []string{
		string(EndpointServiceConnectionStatusApproved),
		string(EndpointServiceConnectionStatusDisconnected),
		string(EndpointServiceConnectionStatusPending),
		string(EndpointServiceConnectionStatusRejected),
	}
}

func parseEndpointServiceConnectionStatus(input string) (*EndpointServiceConnectionStatus, error) {
	vals := map[string]EndpointServiceConnectionStatus{
		"approved":     EndpointServiceConnectionStatusApproved,
		"disconnected": EndpointServiceConnectionStatusDisconnected,
		"pending":      EndpointServiceConnectionStatusPending,
		"rejected":     EndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointServiceConnectionStatus(input)
	return &out, nil
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}