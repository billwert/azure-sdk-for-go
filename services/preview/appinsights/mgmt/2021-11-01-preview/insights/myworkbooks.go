package insights

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MyWorkbooksClient is the composite Swagger for Application Insights Management Client
type MyWorkbooksClient struct {
	BaseClient
}

// NewMyWorkbooksClient creates an instance of the MyWorkbooksClient client.
func NewMyWorkbooksClient(subscriptionID string) MyWorkbooksClient {
	return NewMyWorkbooksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMyWorkbooksClientWithBaseURI creates an instance of the MyWorkbooksClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMyWorkbooksClientWithBaseURI(baseURI string, subscriptionID string) MyWorkbooksClient {
	return MyWorkbooksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new private workbook.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// workbookProperties - properties that need to be specified to create a new private workbook.
// sourceID - azure Resource Id that will fetch all linked workbooks.
func (client MyWorkbooksClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, sourceID string) (result MyWorkbook, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: workbookProperties,
			Constraints: []validation.Constraint{{Target: "workbookProperties.MyWorkbookProperties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "workbookProperties.MyWorkbookProperties.DisplayName", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "workbookProperties.MyWorkbookProperties.SerializedData", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "workbookProperties.MyWorkbookProperties.Category", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, resourceName, workbookProperties, sourceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client MyWorkbooksClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, sourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(sourceID) > 0 {
		queryParameters["sourceId"] = autorest.Encode("query", sourceID)
	}

	workbookProperties.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}", pathParameters),
		autorest.WithJSON(workbookProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) CreateOrUpdateResponder(resp *http.Response) (result MyWorkbook, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete a private workbook.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
func (client MyWorkbooksClient) Delete(ctx context.Context, resourceGroupName string, resourceName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client MyWorkbooksClient) DeletePreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get a single private workbook by its resourceName.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
func (client MyWorkbooksClient) Get(ctx context.Context, resourceGroupName string, resourceName string) (result MyWorkbook, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, resourceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client MyWorkbooksClient) GetPreparer(ctx context.Context, resourceGroupName string, resourceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) GetResponder(resp *http.Response) (result MyWorkbook, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get all private workbooks defined within a specified resource group and category.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// category - category of workbook to return.
// tags - tags presents on each workbook returned.
// sourceID - azure Resource Id that will fetch all linked workbooks.
// canFetchContent - flag indicating whether or not to return the full content for each applicable workbook. If
// false, only return summary content for workbooks.
func (client MyWorkbooksClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, category CategoryType, tags []string, sourceID string, canFetchContent *bool) (result MyWorkbooksListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.mwlr.Response.Response != nil {
				sc = result.mwlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "ListByResourceGroup", err.Error())
	}

	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, category, tags, sourceID, canFetchContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.mwlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.mwlr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.mwlr.hasNextLink() && result.mwlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client MyWorkbooksClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, category CategoryType, tags []string, sourceID string, canFetchContent *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"category":    autorest.Encode("query", category),
	}
	if tags != nil && len(tags) > 0 {
		queryParameters["tags"] = autorest.Encode("query", tags, ",")
	}
	if len(sourceID) > 0 {
		queryParameters["sourceId"] = autorest.Encode("query", sourceID)
	}
	if canFetchContent != nil {
		queryParameters["canFetchContent"] = autorest.Encode("query", *canFetchContent)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) ListByResourceGroupResponder(resp *http.Response) (result MyWorkbooksListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client MyWorkbooksClient) listByResourceGroupNextResults(ctx context.Context, lastResults MyWorkbooksListResult) (result MyWorkbooksListResult, err error) {
	req, err := lastResults.myWorkbooksListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client MyWorkbooksClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, category CategoryType, tags []string, sourceID string, canFetchContent *bool) (result MyWorkbooksListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, category, tags, sourceID, canFetchContent)
	return
}

// ListBySubscription get all private workbooks defined within a specified subscription and category.
// Parameters:
// category - category of workbook to return.
// tags - tags presents on each workbook returned.
// canFetchContent - flag indicating whether or not to return the full content for each applicable workbook. If
// false, only return summary content for workbooks.
func (client MyWorkbooksClient) ListBySubscription(ctx context.Context, category CategoryType, tags []string, canFetchContent *bool) (result MyWorkbooksListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.mwlr.Response.Response != nil {
				sc = result.mwlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "ListBySubscription", err.Error())
	}

	result.fn = client.listBySubscriptionNextResults
	req, err := client.ListBySubscriptionPreparer(ctx, category, tags, canFetchContent)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListBySubscription", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.mwlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListBySubscription", resp, "Failure sending request")
		return
	}

	result.mwlr, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "ListBySubscription", resp, "Failure responding to request")
		return
	}
	if result.mwlr.hasNextLink() && result.mwlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListBySubscriptionPreparer prepares the ListBySubscription request.
func (client MyWorkbooksClient) ListBySubscriptionPreparer(ctx context.Context, category CategoryType, tags []string, canFetchContent *bool) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
		"category":    autorest.Encode("query", category),
	}
	if tags != nil && len(tags) > 0 {
		queryParameters["tags"] = autorest.Encode("query", tags, ",")
	}
	if canFetchContent != nil {
		queryParameters["canFetchContent"] = autorest.Encode("query", *canFetchContent)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/myWorkbooks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionSender sends the ListBySubscription request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) ListBySubscriptionSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionResponder handles the response to the ListBySubscription request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) ListBySubscriptionResponder(resp *http.Response) (result MyWorkbooksListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listBySubscriptionNextResults retrieves the next set of results, if any.
func (client MyWorkbooksClient) listBySubscriptionNextResults(ctx context.Context, lastResults MyWorkbooksListResult) (result MyWorkbooksListResult, err error) {
	req, err := lastResults.myWorkbooksListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listBySubscriptionNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListBySubscriptionSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listBySubscriptionNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListBySubscriptionResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "listBySubscriptionNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListBySubscriptionComplete enumerates all values, automatically crossing page boundaries as required.
func (client MyWorkbooksClient) ListBySubscriptionComplete(ctx context.Context, category CategoryType, tags []string, canFetchContent *bool) (result MyWorkbooksListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.ListBySubscription")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListBySubscription(ctx, category, tags, canFetchContent)
	return
}

// Update updates a private workbook that has already been added.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceName - the name of the Application Insights component resource.
// workbookProperties - properties that need to be specified to create a new private workbook.
// sourceID - azure Resource Id that will fetch all linked workbooks.
func (client MyWorkbooksClient) Update(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, sourceID string) (result MyWorkbook, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MyWorkbooksClient.Update")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("insights.MyWorkbooksClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, resourceName, workbookProperties, sourceID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "insights.MyWorkbooksClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client MyWorkbooksClient) UpdatePreparer(ctx context.Context, resourceGroupName string, resourceName string, workbookProperties MyWorkbook, sourceID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"resourceName":      autorest.Encode("path", resourceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-03-08"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(sourceID) > 0 {
		queryParameters["sourceId"] = autorest.Encode("query", sourceID)
	}

	workbookProperties.SystemData = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/myWorkbooks/{resourceName}", pathParameters),
		autorest.WithJSON(workbookProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client MyWorkbooksClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client MyWorkbooksClient) UpdateResponder(resp *http.Response) (result MyWorkbook, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
