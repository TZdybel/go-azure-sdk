package notificationrecipientemail

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = NotificationRecipientEmailId{}

// NotificationRecipientEmailId is a struct representing the Resource ID for a Notification Recipient Email
type NotificationRecipientEmailId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ServiceName        string
	WorkspaceId        string
	NotificationName   NotificationName
	RecipientEmailName string
}

// NewNotificationRecipientEmailID returns a new NotificationRecipientEmailId struct
func NewNotificationRecipientEmailID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, notificationName NotificationName, recipientEmailName string) NotificationRecipientEmailId {
	return NotificationRecipientEmailId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ServiceName:        serviceName,
		WorkspaceId:        workspaceId,
		NotificationName:   notificationName,
		RecipientEmailName: recipientEmailName,
	}
}

// ParseNotificationRecipientEmailID parses 'input' into a NotificationRecipientEmailId
func ParseNotificationRecipientEmailID(input string) (*NotificationRecipientEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(NotificationRecipientEmailId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NotificationRecipientEmailId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if v, ok := parsed.Parsed["notificationName"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "notificationName", *parsed)
		}

		notificationName, err := parseNotificationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.NotificationName = *notificationName
	}

	if id.RecipientEmailName, ok = parsed.Parsed["recipientEmailName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recipientEmailName", *parsed)
	}

	return &id, nil
}

// ParseNotificationRecipientEmailIDInsensitively parses 'input' case-insensitively into a NotificationRecipientEmailId
// note: this method should only be used for API response data and not user input
func ParseNotificationRecipientEmailIDInsensitively(input string) (*NotificationRecipientEmailId, error) {
	parser := resourceids.NewParserFromResourceIdType(NotificationRecipientEmailId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := NotificationRecipientEmailId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if v, ok := parsed.Parsed["notificationName"]; true {
		if !ok {
			return nil, resourceids.NewSegmentNotSpecifiedError(id, "notificationName", *parsed)
		}

		notificationName, err := parseNotificationName(v)
		if err != nil {
			return nil, fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.NotificationName = *notificationName
	}

	if id.RecipientEmailName, ok = parsed.Parsed["recipientEmailName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "recipientEmailName", *parsed)
	}

	return &id, nil
}

// ValidateNotificationRecipientEmailID checks that 'input' can be parsed as a Notification Recipient Email ID
func ValidateNotificationRecipientEmailID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNotificationRecipientEmailID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Notification Recipient Email ID
func (id NotificationRecipientEmailId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/notifications/%s/recipientEmails/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, string(id.NotificationName), id.RecipientEmailName)
}

// Segments returns a slice of Resource ID Segments which comprise this Notification Recipient Email ID
func (id NotificationRecipientEmailId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceId", "workspaceIdValue"),
		resourceids.StaticSegment("staticNotifications", "notifications", "notifications"),
		resourceids.ConstantSegment("notificationName", PossibleValuesForNotificationName(), "AccountClosedPublisher"),
		resourceids.StaticSegment("staticRecipientEmails", "recipientEmails", "recipientEmails"),
		resourceids.UserSpecifiedSegment("recipientEmailName", "recipientEmailValue"),
	}
}

// String returns a human-readable description of this Notification Recipient Email ID
func (id NotificationRecipientEmailId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Notification Name: %q", string(id.NotificationName)),
		fmt.Sprintf("Recipient Email Name: %q", id.RecipientEmailName),
	}
	return fmt.Sprintf("Notification Recipient Email (%s)", strings.Join(components, "\n"))
}
