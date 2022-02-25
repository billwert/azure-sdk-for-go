//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbatch_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch"
)

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateList.json
func ExampleCertificateClient_ListByBatchAccount() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	pager := client.ListByBatchAccount("<resource-group-name>",
		"<account-name>",
		&armbatch.CertificateClientListByBatchAccountOptions{Maxresults: nil,
			Select: nil,
			Filter: nil,
		})
	for {
		nextResult := pager.NextPage(ctx)
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		if !nextResult {
			break
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("Pager result: %#v\n", v)
		}
	}
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateCreate_Full.json
func ExampleCertificateClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		armbatch.CertificateCreateOrUpdateParameters{
			Properties: &armbatch.CertificateCreateOrUpdateProperties{
				Format:              armbatch.CertificateFormatPfx.ToPtr(),
				Thumbprint:          to.StringPtr("<thumbprint>"),
				ThumbprintAlgorithm: to.StringPtr("<thumbprint-algorithm>"),
				Data:                to.StringPtr("<data>"),
				Password:            to.StringPtr("<password>"),
			},
		},
		&armbatch.CertificateClientCreateOptions{IfMatch: nil,
			IfNoneMatch: nil,
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientCreateResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateUpdate.json
func ExampleCertificateClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.Update(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		armbatch.CertificateCreateOrUpdateParameters{
			Properties: &armbatch.CertificateCreateOrUpdateProperties{
				Data:     to.StringPtr("<data>"),
				Password: to.StringPtr("<password>"),
			},
		},
		&armbatch.CertificateClientUpdateOptions{IfMatch: nil})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientUpdateResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateDelete.json
func ExampleCertificateClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateGet.json
func ExampleCertificateClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientGetResult)
}

// x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2021-06-01/examples/CertificateCancelDeletion.json
func ExampleCertificateClient_CancelDeletion() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armbatch.NewCertificateClient("<subscription-id>", cred, nil)
	res, err := client.CancelDeletion(ctx,
		"<resource-group-name>",
		"<account-name>",
		"<certificate-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.CertificateClientCancelDeletionResult)
}
