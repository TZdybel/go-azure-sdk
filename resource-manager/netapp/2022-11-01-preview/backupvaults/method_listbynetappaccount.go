package backupvaults

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByNetAppAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BackupVault
}

type ListByNetAppAccountCompleteResult struct {
	Items []BackupVault
}

// ListByNetAppAccount ...
func (c BackupVaultsClient) ListByNetAppAccount(ctx context.Context, id NetAppAccountId) (result ListByNetAppAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/backupVaults", id.ID()),
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
		Values *[]BackupVault `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByNetAppAccountComplete retrieves all the results into a single object
func (c BackupVaultsClient) ListByNetAppAccountComplete(ctx context.Context, id NetAppAccountId) (ListByNetAppAccountCompleteResult, error) {
	return c.ListByNetAppAccountCompleteMatchingPredicate(ctx, id, BackupVaultOperationPredicate{})
}

// ListByNetAppAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c BackupVaultsClient) ListByNetAppAccountCompleteMatchingPredicate(ctx context.Context, id NetAppAccountId, predicate BackupVaultOperationPredicate) (result ListByNetAppAccountCompleteResult, err error) {
	items := make([]BackupVault, 0)

	resp, err := c.ListByNetAppAccount(ctx, id)
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

	result = ListByNetAppAccountCompleteResult{
		Items: items,
	}
	return
}
