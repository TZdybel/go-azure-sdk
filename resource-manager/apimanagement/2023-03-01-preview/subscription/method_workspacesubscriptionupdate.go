package subscription

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceSubscriptionUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SubscriptionContract
}

type WorkspaceSubscriptionUpdateOperationOptions struct {
	AppType *AppType
	IfMatch *string
	Notify  *bool
}

func DefaultWorkspaceSubscriptionUpdateOperationOptions() WorkspaceSubscriptionUpdateOperationOptions {
	return WorkspaceSubscriptionUpdateOperationOptions{}
}

func (o WorkspaceSubscriptionUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o WorkspaceSubscriptionUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkspaceSubscriptionUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.AppType != nil {
		out.Append("appType", fmt.Sprintf("%v", *o.AppType))
	}
	if o.Notify != nil {
		out.Append("notify", fmt.Sprintf("%v", *o.Notify))
	}
	return &out
}

// WorkspaceSubscriptionUpdate ...
func (c SubscriptionClient) WorkspaceSubscriptionUpdate(ctx context.Context, id WorkspaceSubscriptions2Id, input SubscriptionUpdateParameters, options WorkspaceSubscriptionUpdateOperationOptions) (result WorkspaceSubscriptionUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		Path:          id.ID(),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
