//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"reflect"
)

// MarketplaceAgreementsClientListPager provides operations for iterating over paged responses.
type MarketplaceAgreementsClientListPager struct {
	client    *MarketplaceAgreementsClient
	current   MarketplaceAgreementsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MarketplaceAgreementsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MarketplaceAgreementsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MarketplaceAgreementsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.AgreementResourceListResponse.NextLink == nil || len(*p.current.AgreementResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MarketplaceAgreementsClientListResponse page.
func (p *MarketplaceAgreementsClientListPager) PageResponse() MarketplaceAgreementsClientListResponse {
	return p.current
}

// MonitorsClientListAPIKeysPager provides operations for iterating over paged responses.
type MonitorsClientListAPIKeysPager struct {
	client    *MonitorsClient
	current   MonitorsClientListAPIKeysResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListAPIKeysResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListAPIKeysPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListAPIKeysPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.APIKeyListResponse.NextLink == nil || len(*p.current.APIKeyListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listAPIKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListAPIKeysResponse page.
func (p *MonitorsClientListAPIKeysPager) PageResponse() MonitorsClientListAPIKeysResponse {
	return p.current
}

// MonitorsClientListByResourceGroupPager provides operations for iterating over paged responses.
type MonitorsClientListByResourceGroupPager struct {
	client    *MonitorsClient
	current   MonitorsClientListByResourceGroupResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListByResourceGroupResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListByResourceGroupPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListByResourceGroupPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MonitorResourceListResponse.NextLink == nil || len(*p.current.MonitorResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listByResourceGroupHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListByResourceGroupResponse page.
func (p *MonitorsClientListByResourceGroupPager) PageResponse() MonitorsClientListByResourceGroupResponse {
	return p.current
}

// MonitorsClientListHostsPager provides operations for iterating over paged responses.
type MonitorsClientListHostsPager struct {
	client    *MonitorsClient
	current   MonitorsClientListHostsResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListHostsResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListHostsPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListHostsPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.HostListResponse.NextLink == nil || len(*p.current.HostListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHostsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListHostsResponse page.
func (p *MonitorsClientListHostsPager) PageResponse() MonitorsClientListHostsResponse {
	return p.current
}

// MonitorsClientListLinkedResourcesPager provides operations for iterating over paged responses.
type MonitorsClientListLinkedResourcesPager struct {
	client    *MonitorsClient
	current   MonitorsClientListLinkedResourcesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListLinkedResourcesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListLinkedResourcesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListLinkedResourcesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.LinkedResourceListResponse.NextLink == nil || len(*p.current.LinkedResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listLinkedResourcesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListLinkedResourcesResponse page.
func (p *MonitorsClientListLinkedResourcesPager) PageResponse() MonitorsClientListLinkedResourcesResponse {
	return p.current
}

// MonitorsClientListMonitoredResourcesPager provides operations for iterating over paged responses.
type MonitorsClientListMonitoredResourcesPager struct {
	client    *MonitorsClient
	current   MonitorsClientListMonitoredResourcesResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListMonitoredResourcesResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListMonitoredResourcesPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListMonitoredResourcesPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MonitoredResourceListResponse.NextLink == nil || len(*p.current.MonitoredResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listMonitoredResourcesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListMonitoredResourcesResponse page.
func (p *MonitorsClientListMonitoredResourcesPager) PageResponse() MonitorsClientListMonitoredResourcesResponse {
	return p.current
}

// MonitorsClientListPager provides operations for iterating over paged responses.
type MonitorsClientListPager struct {
	client    *MonitorsClient
	current   MonitorsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, MonitorsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *MonitorsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *MonitorsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MonitorResourceListResponse.NextLink == nil || len(*p.current.MonitorResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current MonitorsClientListResponse page.
func (p *MonitorsClientListPager) PageResponse() MonitorsClientListResponse {
	return p.current
}

// OperationsClientListPager provides operations for iterating over paged responses.
type OperationsClientListPager struct {
	client    *OperationsClient
	current   OperationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, OperationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *OperationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *OperationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.OperationListResult.NextLink == nil || len(*p.current.OperationListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current OperationsClientListResponse page.
func (p *OperationsClientListPager) PageResponse() OperationsClientListResponse {
	return p.current
}

// SingleSignOnConfigurationsClientListPager provides operations for iterating over paged responses.
type SingleSignOnConfigurationsClientListPager struct {
	client    *SingleSignOnConfigurationsClient
	current   SingleSignOnConfigurationsClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, SingleSignOnConfigurationsClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *SingleSignOnConfigurationsClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *SingleSignOnConfigurationsClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SingleSignOnResourceListResponse.NextLink == nil || len(*p.current.SingleSignOnResourceListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current SingleSignOnConfigurationsClientListResponse page.
func (p *SingleSignOnConfigurationsClientListPager) PageResponse() SingleSignOnConfigurationsClientListResponse {
	return p.current
}

// TagRulesClientListPager provides operations for iterating over paged responses.
type TagRulesClientListPager struct {
	client    *TagRulesClient
	current   TagRulesClientListResponse
	err       error
	requester func(context.Context) (*policy.Request, error)
	advancer  func(context.Context, TagRulesClientListResponse) (*policy.Request, error)
}

// Err returns the last error encountered while paging.
func (p *TagRulesClientListPager) Err() error {
	return p.err
}

// NextPage returns true if the pager advanced to the next page.
// Returns false if there are no more pages or an error occurred.
func (p *TagRulesClientListPager) NextPage(ctx context.Context) bool {
	var req *policy.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.MonitoringTagRulesListResponse.NextLink == nil || len(*p.current.MonitoringTagRulesListResponse.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.pl.Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		p.err = runtime.NewResponseError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

// PageResponse returns the current TagRulesClientListResponse page.
func (p *TagRulesClientListPager) PageResponse() TagRulesClientListResponse {
	return p.current
}
