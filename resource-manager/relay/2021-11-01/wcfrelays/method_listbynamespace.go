package wcfrelays

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByNamespaceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WcfRelay
}

type ListByNamespaceCompleteResult struct {
	Items []WcfRelay
}

// ListByNamespace ...
func (c WCFRelaysClient) ListByNamespace(ctx context.Context, id NamespaceId) (result ListByNamespaceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/wcfRelays", id.ID()),
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
		Values *[]WcfRelay `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByNamespaceComplete retrieves all the results into a single object
func (c WCFRelaysClient) ListByNamespaceComplete(ctx context.Context, id NamespaceId) (ListByNamespaceCompleteResult, error) {
	return c.ListByNamespaceCompleteMatchingPredicate(ctx, id, WcfRelayOperationPredicate{})
}

// ListByNamespaceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WCFRelaysClient) ListByNamespaceCompleteMatchingPredicate(ctx context.Context, id NamespaceId, predicate WcfRelayOperationPredicate) (result ListByNamespaceCompleteResult, err error) {
	items := make([]WcfRelay, 0)

	resp, err := c.ListByNamespace(ctx, id)
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

	result = ListByNamespaceCompleteResult{
		Items: items,
	}
	return
}
