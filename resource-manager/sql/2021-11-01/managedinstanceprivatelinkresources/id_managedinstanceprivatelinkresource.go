package managedinstanceprivatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedInstancePrivateLinkResourceId{}

// ManagedInstancePrivateLinkResourceId is a struct representing the Resource ID for a Managed Instance Private Link Resource
type ManagedInstancePrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	ManagedInstanceName     string
	PrivateLinkResourceName string
}

// NewManagedInstancePrivateLinkResourceID returns a new ManagedInstancePrivateLinkResourceId struct
func NewManagedInstancePrivateLinkResourceID(subscriptionId string, resourceGroupName string, managedInstanceName string, privateLinkResourceName string) ManagedInstancePrivateLinkResourceId {
	return ManagedInstancePrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		ManagedInstanceName:     managedInstanceName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParseManagedInstancePrivateLinkResourceID parses 'input' into a ManagedInstancePrivateLinkResourceId
func ParseManagedInstancePrivateLinkResourceID(input string) (*ManagedInstancePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedInstancePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedInstancePrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateLinkResourceName", *parsed)
	}

	return &id, nil
}

// ParseManagedInstancePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a ManagedInstancePrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParseManagedInstancePrivateLinkResourceIDInsensitively(input string) (*ManagedInstancePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedInstancePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedInstancePrivateLinkResourceId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ManagedInstanceName, ok = parsed.Parsed["managedInstanceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", *parsed)
	}

	if id.PrivateLinkResourceName, ok = parsed.Parsed["privateLinkResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateLinkResourceName", *parsed)
	}

	return &id, nil
}

// ValidateManagedInstancePrivateLinkResourceID checks that 'input' can be parsed as a Managed Instance Private Link Resource ID
func ValidateManagedInstancePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstancePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "privateLinkResourceValue"),
	}
}

// String returns a human-readable description of this Managed Instance Private Link Resource ID
func (id ManagedInstancePrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Managed Instance Private Link Resource (%s)", strings.Join(components, "\n"))
}
