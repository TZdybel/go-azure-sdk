package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartNetworkTraceOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type StartNetworkTraceOperationOptions struct {
	DurationInSeconds *int64
	MaxFrameLength    *int64
	SasUrl            *string
}

func DefaultStartNetworkTraceOperationOptions() StartNetworkTraceOperationOptions {
	return StartNetworkTraceOperationOptions{}
}

func (o StartNetworkTraceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StartNetworkTraceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o StartNetworkTraceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.DurationInSeconds != nil {
		out.Append("durationInSeconds", fmt.Sprintf("%v", *o.DurationInSeconds))
	}
	if o.MaxFrameLength != nil {
		out.Append("maxFrameLength", fmt.Sprintf("%v", *o.MaxFrameLength))
	}
	if o.SasUrl != nil {
		out.Append("sasUrl", fmt.Sprintf("%v", *o.SasUrl))
	}
	return &out
}

// StartNetworkTrace ...
func (c WebAppsClient) StartNetworkTrace(ctx context.Context, id commonids.AppServiceId, options StartNetworkTraceOperationOptions) (result StartNetworkTraceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/startNetworkTrace", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// StartNetworkTraceThenPoll performs StartNetworkTrace then polls until it's completed
func (c WebAppsClient) StartNetworkTraceThenPoll(ctx context.Context, id commonids.AppServiceId, options StartNetworkTraceOperationOptions) error {
	result, err := c.StartNetworkTrace(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing StartNetworkTrace: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after StartNetworkTrace: %+v", err)
	}

	return nil
}
