package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDeploymentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Deployment
}

type ListDeploymentsCompleteResult struct {
	Items []Deployment
}

// ListDeployments ...
func (c WebAppsClient) ListDeployments(ctx context.Context, id commonids.AppServiceId) (result ListDeploymentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/deployments", id.ID()),
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
		Values *[]Deployment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDeploymentsComplete retrieves all the results into a single object
func (c WebAppsClient) ListDeploymentsComplete(ctx context.Context, id commonids.AppServiceId) (ListDeploymentsCompleteResult, error) {
	return c.ListDeploymentsCompleteMatchingPredicate(ctx, id, DeploymentOperationPredicate{})
}

// ListDeploymentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListDeploymentsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate DeploymentOperationPredicate) (result ListDeploymentsCompleteResult, err error) {
	items := make([]Deployment, 0)

	resp, err := c.ListDeployments(ctx, id)
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

	result = ListDeploymentsCompleteResult{
		Items: items,
	}
	return
}
