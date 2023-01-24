package standardoperation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceSku struct {
	Capacity *int64  `json:"capacity,omitempty"`
	Family   *string `json:"family,omitempty"`
	Name     *string `json:"name,omitempty"`
	Size     *string `json:"size,omitempty"`
	Tier     *string `json:"tier,omitempty"`
}
