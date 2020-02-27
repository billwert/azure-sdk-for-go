package servicemap

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// ProcessesClient is the service Map API Reference
type ProcessesClient struct {
	BaseClient
}

// NewProcessesClient creates an instance of the ProcessesClient client.
func NewProcessesClient(subscriptionID string) ProcessesClient {
	return NewProcessesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewProcessesClientWithBaseURI creates an instance of the ProcessesClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewProcessesClientWithBaseURI(baseURI string, subscriptionID string) ProcessesClient {
	return ProcessesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get returns the specified process.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// processName - process resource name.
// timestamp - UTC date and time specifying a time instance relative to which to evaluate a resource. When not
// specified, the service uses DateTime.UtcNow.
func (client ProcessesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, timestamp *date.Time) (result Process, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: processName,
			Constraints: []validation.Constraint{{Target: "processName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "processName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.ProcessesClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, workspaceName, machineName, processName, timestamp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ProcessesClient) GetPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, timestamp *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"processName":       autorest.Encode("path", processName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if timestamp != nil {
		queryParameters["timestamp"] = autorest.Encode("query", *timestamp)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ProcessesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ProcessesClient) GetResponder(resp *http.Response) (result Process, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetLiveness obtains the liveness status of the process during the specified time interval.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// processName - process resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client ProcessesClient) GetLiveness(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (result Liveness, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.GetLiveness")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: processName,
			Constraints: []validation.Constraint{{Target: "processName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "processName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.ProcessesClient", "GetLiveness", err.Error())
	}

	req, err := client.GetLivenessPreparer(ctx, resourceGroupName, workspaceName, machineName, processName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "GetLiveness", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetLivenessSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "GetLiveness", resp, "Failure sending request")
		return
	}

	result, err = client.GetLivenessResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "GetLiveness", resp, "Failure responding to request")
	}

	return
}

// GetLivenessPreparer prepares the GetLiveness request.
func (client ProcessesClient) GetLivenessPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"processName":       autorest.Encode("path", processName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/liveness", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetLivenessSender sends the GetLiveness request. The method will close the
// http.Response Body if it receives an error.
func (client ProcessesClient) GetLivenessSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetLivenessResponder handles the response to the GetLiveness request. The method always
// closes the http.Response Body.
func (client ProcessesClient) GetLivenessResponder(resp *http.Response) (result Liveness, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListAcceptingPorts returns a collection of ports on which this process is accepting
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// processName - process resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client ProcessesClient) ListAcceptingPorts(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (result PortCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.ListAcceptingPorts")
		defer func() {
			sc := -1
			if result.pc.Response.Response != nil {
				sc = result.pc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: processName,
			Constraints: []validation.Constraint{{Target: "processName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "processName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.ProcessesClient", "ListAcceptingPorts", err.Error())
	}

	result.fn = client.listAcceptingPortsNextResults
	req, err := client.ListAcceptingPortsPreparer(ctx, resourceGroupName, workspaceName, machineName, processName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListAcceptingPorts", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListAcceptingPortsSender(req)
	if err != nil {
		result.pc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListAcceptingPorts", resp, "Failure sending request")
		return
	}

	result.pc, err = client.ListAcceptingPortsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListAcceptingPorts", resp, "Failure responding to request")
	}

	return
}

// ListAcceptingPortsPreparer prepares the ListAcceptingPorts request.
func (client ProcessesClient) ListAcceptingPortsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"processName":       autorest.Encode("path", processName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/acceptingPorts", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListAcceptingPortsSender sends the ListAcceptingPorts request. The method will close the
// http.Response Body if it receives an error.
func (client ProcessesClient) ListAcceptingPortsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListAcceptingPortsResponder handles the response to the ListAcceptingPorts request. The method always
// closes the http.Response Body.
func (client ProcessesClient) ListAcceptingPortsResponder(resp *http.Response) (result PortCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listAcceptingPortsNextResults retrieves the next set of results, if any.
func (client ProcessesClient) listAcceptingPortsNextResults(ctx context.Context, lastResults PortCollection) (result PortCollection, err error) {
	req, err := lastResults.portCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listAcceptingPortsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListAcceptingPortsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listAcceptingPortsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListAcceptingPortsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listAcceptingPortsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListAcceptingPortsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProcessesClient) ListAcceptingPortsComplete(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (result PortCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.ListAcceptingPorts")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListAcceptingPorts(ctx, resourceGroupName, workspaceName, machineName, processName, startTime, endTime)
	return
}

// ListConnections returns a collection of connections terminating or originating at the specified process
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// machineName - machine resource name.
// processName - process resource name.
// startTime - UTC date and time specifying the start time of an interval. When not specified the service uses
// DateTime.UtcNow - 10m
// endTime - UTC date and time specifying the end time of an interval. When not specified the service uses
// DateTime.UtcNow
func (client ProcessesClient) ListConnections(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (result ConnectionCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.ListConnections")
		defer func() {
			sc := -1
			if result.cc.Response.Response != nil {
				sc = result.cc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}},
		{TargetValue: machineName,
			Constraints: []validation.Constraint{{Target: "machineName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "machineName", Name: validation.MinLength, Rule: 3, Chain: nil}}},
		{TargetValue: processName,
			Constraints: []validation.Constraint{{Target: "processName", Name: validation.MaxLength, Rule: 128, Chain: nil},
				{Target: "processName", Name: validation.MinLength, Rule: 3, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.ProcessesClient", "ListConnections", err.Error())
	}

	result.fn = client.listConnectionsNextResults
	req, err := client.ListConnectionsPreparer(ctx, resourceGroupName, workspaceName, machineName, processName, startTime, endTime)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListConnections", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.cc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListConnections", resp, "Failure sending request")
		return
	}

	result.cc, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "ListConnections", resp, "Failure responding to request")
	}

	return
}

// ListConnectionsPreparer prepares the ListConnections request.
func (client ProcessesClient) ListConnectionsPreparer(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"machineName":       autorest.Encode("path", machineName),
		"processName":       autorest.Encode("path", processName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if startTime != nil {
		queryParameters["startTime"] = autorest.Encode("query", *startTime)
	}
	if endTime != nil {
		queryParameters["endTime"] = autorest.Encode("query", *endTime)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/machines/{machineName}/processes/{processName}/connections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListConnectionsSender sends the ListConnections request. The method will close the
// http.Response Body if it receives an error.
func (client ProcessesClient) ListConnectionsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListConnectionsResponder handles the response to the ListConnections request. The method always
// closes the http.Response Body.
func (client ProcessesClient) ListConnectionsResponder(resp *http.Response) (result ConnectionCollection, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listConnectionsNextResults retrieves the next set of results, if any.
func (client ProcessesClient) listConnectionsNextResults(ctx context.Context, lastResults ConnectionCollection) (result ConnectionCollection, err error) {
	req, err := lastResults.connectionCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listConnectionsNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListConnectionsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listConnectionsNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListConnectionsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.ProcessesClient", "listConnectionsNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListConnectionsComplete enumerates all values, automatically crossing page boundaries as required.
func (client ProcessesClient) ListConnectionsComplete(ctx context.Context, resourceGroupName string, workspaceName string, machineName string, processName string, startTime *date.Time, endTime *date.Time) (result ConnectionCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ProcessesClient.ListConnections")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListConnections(ctx, resourceGroupName, workspaceName, machineName, processName, startTime, endTime)
	return
}
