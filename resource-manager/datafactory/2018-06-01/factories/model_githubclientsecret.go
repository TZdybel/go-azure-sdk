package factories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GitHubClientSecret struct {
	ByoaSecretAkvUrl *string `json:"byoaSecretAkvUrl,omitempty"`
	ByoaSecretName   *string `json:"byoaSecretName,omitempty"`
}
