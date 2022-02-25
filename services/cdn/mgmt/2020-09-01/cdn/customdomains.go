package cdn

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

// CustomDomainsClient is the cdn Management Client
type CustomDomainsClient struct {
	BaseClient
}

// NewCustomDomainsClient creates an instance of the CustomDomainsClient client.
func NewCustomDomainsClient(subscriptionID string) CustomDomainsClient {
	return NewCustomDomainsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewCustomDomainsClientWithBaseURI creates an instance of the CustomDomainsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewCustomDomainsClientWithBaseURI(baseURI string, subscriptionID string) CustomDomainsClient {
	return CustomDomainsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Create creates a new custom domain within an endpoint.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
// customDomainName - name of the custom domain within an endpoint.
// customDomainProperties - properties required to create a new custom domain.
func (client CustomDomainsClient) Create(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters) (result CustomDomainsCreateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.Create")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}},
		{TargetValue: customDomainProperties,
			Constraints: []validation.Constraint{{Target: "customDomainProperties.CustomDomainPropertiesParameters", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "customDomainProperties.CustomDomainPropertiesParameters.HostName", Name: validation.Null, Rule: true, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "Create", err.Error())
	}

	req, err := client.CreatePreparer(ctx, resourceGroupName, profileName, endpointName, customDomainName, customDomainProperties)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Create", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Create", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreatePreparer prepares the Create request.
func (client CustomDomainsClient) CreatePreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainProperties CustomDomainParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithJSON(customDomainProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateSender sends the Create request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) CreateSender(req *http.Request) (future CustomDomainsCreateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateResponder handles the response to the Create request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) CreateResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes an existing custom domain within an endpoint.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
// customDomainName - name of the custom domain within an endpoint.
func (client CustomDomainsClient) Delete(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomainsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client CustomDomainsClient) DeletePreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) DeleteSender(req *http.Request) (future CustomDomainsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) DeleteResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DisableCustomHTTPS disable https delivery of the custom domain.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
// customDomainName - name of the custom domain within an endpoint.
func (client CustomDomainsClient) DisableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomainsDisableCustomHTTPSFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.DisableCustomHTTPS")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "DisableCustomHTTPS", err.Error())
	}

	req, err := client.DisableCustomHTTPSPreparer(ctx, resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS", nil, "Failure preparing request")
		return
	}

	result, err = client.DisableCustomHTTPSSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "DisableCustomHTTPS", result.Response(), "Failure sending request")
		return
	}

	return
}

// DisableCustomHTTPSPreparer prepares the DisableCustomHTTPS request.
func (client CustomDomainsClient) DisableCustomHTTPSPreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/disableCustomHttps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DisableCustomHTTPSSender sends the DisableCustomHTTPS request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) DisableCustomHTTPSSender(req *http.Request) (future CustomDomainsDisableCustomHTTPSFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DisableCustomHTTPSResponder handles the response to the DisableCustomHTTPS request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) DisableCustomHTTPSResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// EnableCustomHTTPS enable https delivery of the custom domain.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
// customDomainName - name of the custom domain within an endpoint.
// customDomainHTTPSParameters - the configuration specifying how to enable HTTPS for the custom domain - using
// CDN managed certificate or user's own certificate. If not specified, enabling ssl uses CDN managed
// certificate by default.
func (client CustomDomainsClient) EnableCustomHTTPS(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainHTTPSParameters *BasicCustomDomainHTTPSParameters) (result CustomDomainsEnableCustomHTTPSFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.EnableCustomHTTPS")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "EnableCustomHTTPS", err.Error())
	}

	req, err := client.EnableCustomHTTPSPreparer(ctx, resourceGroupName, profileName, endpointName, customDomainName, customDomainHTTPSParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS", nil, "Failure preparing request")
		return
	}

	result, err = client.EnableCustomHTTPSSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "EnableCustomHTTPS", result.Response(), "Failure sending request")
		return
	}

	return
}

// EnableCustomHTTPSPreparer prepares the EnableCustomHTTPS request.
func (client CustomDomainsClient) EnableCustomHTTPSPreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string, customDomainHTTPSParameters *BasicCustomDomainHTTPSParameters) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}/enableCustomHttps", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if customDomainHTTPSParameters != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(customDomainHTTPSParameters))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableCustomHTTPSSender sends the EnableCustomHTTPS request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) EnableCustomHTTPSSender(req *http.Request) (future CustomDomainsEnableCustomHTTPSFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// EnableCustomHTTPSResponder handles the response to the EnableCustomHTTPS request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) EnableCustomHTTPSResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets an existing custom domain within an endpoint.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
// customDomainName - name of the custom domain within an endpoint.
func (client CustomDomainsClient) Get(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (result CustomDomain, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.Get")
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
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, profileName, endpointName, customDomainName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client CustomDomainsClient) GetPreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string, customDomainName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"customDomainName":  autorest.Encode("path", customDomainName),
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains/{customDomainName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) GetResponder(resp *http.Response) (result CustomDomain, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpoint lists all of the existing custom domains within an endpoint.
// Parameters:
// resourceGroupName - name of the Resource group within the Azure subscription.
// profileName - name of the CDN profile which is unique within the resource group.
// endpointName - name of the endpoint under the profile which is unique globally.
func (client CustomDomainsClient) ListByEndpoint(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result CustomDomainListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.ListByEndpoint")
		defer func() {
			sc := -1
			if result.cdlr.Response.Response != nil {
				sc = result.cdlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("cdn.CustomDomainsClient", "ListByEndpoint", err.Error())
	}

	result.fn = client.listByEndpointNextResults
	req, err := client.ListByEndpointPreparer(ctx, resourceGroupName, profileName, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.cdlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure sending request")
		return
	}

	result.cdlr, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "ListByEndpoint", resp, "Failure responding to request")
		return
	}
	if result.cdlr.hasNextLink() && result.cdlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByEndpointPreparer prepares the ListByEndpoint request.
func (client CustomDomainsClient) ListByEndpointPreparer(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/customDomains", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByEndpointSender sends the ListByEndpoint request. The method will close the
// http.Response Body if it receives an error.
func (client CustomDomainsClient) ListByEndpointSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByEndpointResponder handles the response to the ListByEndpoint request. The method always
// closes the http.Response Body.
func (client CustomDomainsClient) ListByEndpointResponder(resp *http.Response) (result CustomDomainListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByEndpointNextResults retrieves the next set of results, if any.
func (client CustomDomainsClient) listByEndpointNextResults(ctx context.Context, lastResults CustomDomainListResult) (result CustomDomainListResult, err error) {
	req, err := lastResults.customDomainListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "listByEndpointNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "listByEndpointNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.CustomDomainsClient", "listByEndpointNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByEndpointComplete enumerates all values, automatically crossing page boundaries as required.
func (client CustomDomainsClient) ListByEndpointComplete(ctx context.Context, resourceGroupName string, profileName string, endpointName string) (result CustomDomainListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/CustomDomainsClient.ListByEndpoint")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByEndpoint(ctx, resourceGroupName, profileName, endpointName)
	return
}
