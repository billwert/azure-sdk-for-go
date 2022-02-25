//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/domainservices/armdomainservices"
)

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesBySubscription.json
func ExampleClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	pager := client.List(nil)
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

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/ListDomainServicesByResourceGroup.json
func ExampleClient_ListByResourceGroup() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	pager := client.ListByResourceGroup("<resource-group-name>",
		nil)
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

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/CreateDomainService.json
func ExampleClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginCreateOrUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				DomainName: to.StringPtr("<domain-name>"),
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            armdomainservices.NtlmV1("Enabled").ToPtr(),
					SyncNtlmPasswords: armdomainservices.SyncNtlmPasswords("Enabled").ToPtr(),
					TLSV1:             armdomainservices.TLSV1("Disabled").ToPtr(),
				},
				FilteredSync: armdomainservices.FilteredSync("Enabled").ToPtr(),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         armdomainservices.ExternalAccess("Enabled").ToPtr(),
					Ldaps:                  armdomainservices.Ldaps("Enabled").ToPtr(),
					PfxCertificate:         to.StringPtr("<pfx-certificate>"),
					PfxCertificatePassword: to.StringPtr("<pfx-certificate-password>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.StringPtr("jicha@microsoft.com"),
						to.StringPtr("caalmont@microsoft.com")},
					NotifyDcAdmins:     armdomainservices.NotifyDcAdmins("Enabled").ToPtr(),
					NotifyGlobalAdmins: armdomainservices.NotifyGlobalAdmins("Enabled").ToPtr(),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.StringPtr("<location>"),
						SubnetID: to.StringPtr("<subnet-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientCreateOrUpdateResult)
}

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/GetDomainService.json
func ExampleClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientGetResult)
}

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/DeleteDomainService.json
func ExampleClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/domainservices/resource-manager/Microsoft.AAD/stable/2021-05-01/examples/UpdateDomainService.json
func ExampleClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armdomainservices.NewClient("<subscription-id>", cred, nil)
	poller, err := client.BeginUpdate(ctx,
		"<resource-group-name>",
		"<domain-service-name>",
		armdomainservices.DomainService{
			Properties: &armdomainservices.DomainServiceProperties{
				ConfigDiagnostics: &armdomainservices.ConfigDiagnostics{
					LastExecuted: to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC1123, "2021-05-05T12:00:23Z;"); return t }()),
					ValidatorResults: []*armdomainservices.ConfigDiagnosticsValidatorResult{
						{
							Issues: []*armdomainservices.ConfigDiagnosticsValidatorResultIssue{
								{
									DescriptionParams: []*string{},
									ID:                to.StringPtr("<id>"),
								}},
							ReplicaSetSubnetDisplayName: to.StringPtr("<replica-set-subnet-display-name>"),
							Status:                      armdomainservices.Status("Warning").ToPtr(),
							ValidatorID:                 to.StringPtr("<validator-id>"),
						}},
				},
				DomainSecuritySettings: &armdomainservices.DomainSecuritySettings{
					NtlmV1:            armdomainservices.NtlmV1("Enabled").ToPtr(),
					SyncNtlmPasswords: armdomainservices.SyncNtlmPasswords("Enabled").ToPtr(),
					TLSV1:             armdomainservices.TLSV1("Disabled").ToPtr(),
				},
				FilteredSync: armdomainservices.FilteredSync("Enabled").ToPtr(),
				LdapsSettings: &armdomainservices.LdapsSettings{
					ExternalAccess:         armdomainservices.ExternalAccess("Enabled").ToPtr(),
					Ldaps:                  armdomainservices.Ldaps("Enabled").ToPtr(),
					PfxCertificate:         to.StringPtr("<pfx-certificate>"),
					PfxCertificatePassword: to.StringPtr("<pfx-certificate-password>"),
				},
				NotificationSettings: &armdomainservices.NotificationSettings{
					AdditionalRecipients: []*string{
						to.StringPtr("jicha@microsoft.com"),
						to.StringPtr("caalmont@microsoft.com")},
					NotifyDcAdmins:     armdomainservices.NotifyDcAdmins("Enabled").ToPtr(),
					NotifyGlobalAdmins: armdomainservices.NotifyGlobalAdmins("Enabled").ToPtr(),
				},
				ReplicaSets: []*armdomainservices.ReplicaSet{
					{
						Location: to.StringPtr("<location>"),
						SubnetID: to.StringPtr("<subnet-id>"),
					},
					{
						Location: to.StringPtr("<location>"),
						SubnetID: to.StringPtr("<subnet-id>"),
					}},
			},
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.ClientUpdateResult)
}
