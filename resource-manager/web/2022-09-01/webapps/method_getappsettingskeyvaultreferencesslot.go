package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAppSettingsKeyVaultReferencesSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ApiKVReference
}

type GetAppSettingsKeyVaultReferencesSlotCompleteResult struct {
	Items []ApiKVReference
}

// GetAppSettingsKeyVaultReferencesSlot ...
func (c WebAppsClient) GetAppSettingsKeyVaultReferencesSlot(ctx context.Context, id SlotId) (result GetAppSettingsKeyVaultReferencesSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/config/configReferences/appSettings", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]ApiKVReference `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetAppSettingsKeyVaultReferencesSlotComplete retrieves all the results into a single object
func (c WebAppsClient) GetAppSettingsKeyVaultReferencesSlotComplete(ctx context.Context, id SlotId) (GetAppSettingsKeyVaultReferencesSlotCompleteResult, error) {
	return c.GetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate(ctx, id, ApiKVReferenceOperationPredicate{})
}

// GetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) GetAppSettingsKeyVaultReferencesSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate ApiKVReferenceOperationPredicate) (result GetAppSettingsKeyVaultReferencesSlotCompleteResult, err error) {
	items := make([]ApiKVReference, 0)

	resp, err := c.GetAppSettingsKeyVaultReferencesSlot(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = GetAppSettingsKeyVaultReferencesSlotCompleteResult{
		Items: items,
	}
	return
}
