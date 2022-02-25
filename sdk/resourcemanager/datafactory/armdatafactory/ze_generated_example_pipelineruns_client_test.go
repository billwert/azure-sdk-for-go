//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datafactory/armdatafactory"
)

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_QueryByFactory.json
func ExamplePipelineRunsClient_QueryByFactory() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewPipelineRunsClient("<subscription-id>", cred, nil)
	res, err := client.QueryByFactory(ctx,
		"<resource-group-name>",
		"<factory-name>",
		armdatafactory.RunFilterParameters{
			Filters: []*armdatafactory.RunQueryFilter{
				{
					Operand:  armdatafactory.RunQueryFilterOperand("PipelineName").ToPtr(),
					Operator: armdatafactory.RunQueryFilterOperator("Equals").ToPtr(),
					Values: []*string{
						to.StringPtr("examplePipeline")},
				}},
			LastUpdatedAfter:  to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:36:44.3345758Z"); return t }()),
			LastUpdatedBefore: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-06-16T00:49:48.3686473Z"); return t }()),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineRunsClientQueryByFactoryResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Get.json
func ExamplePipelineRunsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewPipelineRunsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<run-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.PipelineRunsClientGetResult)
}

// x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/PipelineRuns_Cancel.json
func ExamplePipelineRunsClient_Cancel() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdatafactory.NewPipelineRunsClient("<subscription-id>", cred, nil)
	_, err = client.Cancel(ctx,
		"<resource-group-name>",
		"<factory-name>",
		"<run-id>",
		&armdatafactory.PipelineRunsClientCancelOptions{IsRecursive: nil})
	if err != nil {
		log.Fatal(err)
	}
}
