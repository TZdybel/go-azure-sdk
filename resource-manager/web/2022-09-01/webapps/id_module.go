package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ModuleId{}

// ModuleId is a struct representing the Resource ID for a Module
type ModuleId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	ProcessId         string
	ModuleName        string
}

// NewModuleID returns a new ModuleId struct
func NewModuleID(subscriptionId string, resourceGroupName string, siteName string, processId string, moduleName string) ModuleId {
	return ModuleId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		ProcessId:         processId,
		ModuleName:        moduleName,
	}
}

// ParseModuleID parses 'input' into a ModuleId
func ParseModuleID(input string) (*ModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(ModuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	return &id, nil
}

// ParseModuleIDInsensitively parses 'input' case-insensitively into a ModuleId
// note: this method should only be used for API response data and not user input
func ParseModuleIDInsensitively(input string) (*ModuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(ModuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ModuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	if id.ModuleName, ok = parsed.Parsed["moduleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "moduleName", *parsed)
	}

	return &id, nil
}

// ValidateModuleID checks that 'input' can be parsed as a Module ID
func ValidateModuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Module ID
func (id ModuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/processes/%s/modules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.ProcessId, id.ModuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Module ID
func (id ModuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticProcesses", "processes", "processes"),
		resourceids.UserSpecifiedSegment("processId", "processIdValue"),
		resourceids.StaticSegment("staticModules", "modules", "modules"),
		resourceids.UserSpecifiedSegment("moduleName", "moduleValue"),
	}
}

// String returns a human-readable description of this Module ID
func (id ModuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Process: %q", id.ProcessId),
		fmt.Sprintf("Module Name: %q", id.ModuleName),
	}
	return fmt.Sprintf("Module (%s)", strings.Join(components, "\n"))
}
