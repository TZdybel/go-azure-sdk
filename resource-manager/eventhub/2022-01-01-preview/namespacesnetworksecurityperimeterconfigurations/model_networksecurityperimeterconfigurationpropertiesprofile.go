package namespacesnetworksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationPropertiesProfile struct {
	AccessRules        *[]NspAccessRule `json:"accessRules,omitempty"`
	AccessRulesVersion *string          `json:"accessRulesVersion,omitempty"`
	Name               *string          `json:"name,omitempty"`
}
