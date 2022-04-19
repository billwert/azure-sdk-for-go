//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchangeanalysis_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/changeanalysis/armchangeanalysis"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesByResourceGroup.json
func ExampleChangesClient_NewListChangesByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchangeanalysis.NewChangesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListChangesByResourceGroupPager("<resource-group-name>",
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesClientListChangesByResourceGroupOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/changeanalysis/resource-manager/Microsoft.ChangeAnalysis/stable/2021-04-01/examples/ChangesListChangesBySubscription.json
func ExampleChangesClient_NewListChangesBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armchangeanalysis.NewChangesClient("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListChangesBySubscriptionPager(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-25T12:09:03.141Z"); return t }(),
		func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2021-04-26T12:09:03.141Z"); return t }(),
		&armchangeanalysis.ChangesClientListChangesBySubscriptionOptions{SkipToken: nil})
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
			return
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}
