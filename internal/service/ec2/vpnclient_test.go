// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfsync "github.com/hashicorp/terraform-provider-aws/internal/experimental/sync"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// This is part of an experimental feature, do not use this as a starting point for tests
//
//	"This place is not a place of honor... no highly esteemed deed is commemorated here... nothing valued is here.
//	What is here was dangerous and repulsive to us. This message is a warning about danger."
//	--  https://hyperallergic.com/312318/a-nuclear-warning-designed-to-last-10000-years/
func TestAccClientVPNEndpoint_serial(t *testing.T) {
	t.Parallel()

	semaphore := tfsync.GetSemaphore("ClientVPN", "AWS_EC2_CLIENT_VPN_LIMIT", 5)
	testCases := map[string]map[string]func(*testing.T, tfsync.Semaphore){
		"Endpoint": {
			acctest.CtBasic:                testAccClientVPNEndpoint_basic,
			"disappears":                   testAccClientVPNEndpoint_disappears,
			"msADAuth":                     testAccClientVPNEndpoint_msADAuth,
			"msADAuthAndMutualAuth":        testAccClientVPNEndpoint_msADAuthAndMutualAuth,
			"federatedAuth":                testAccClientVPNEndpoint_federatedAuth,
			"federatedAuthWithSelfService": testAccClientVPNEndpoint_federatedAuthWithSelfServiceProvider,
			"withClientConnect":            testAccClientVPNEndpoint_withClientConnectOptions,
			"withClientLoginBanner":        testAccClientVPNEndpoint_withClientLoginBannerOptions,
			"withLogGroup":                 testAccClientVPNEndpoint_withConnectionLogOptions,
			"withDNSServers":               testAccClientVPNEndpoint_withDNSServers,
			names.AttrTags:                 testAccClientVPNEndpoint_tags,
			"simpleAttributesUpdate":       testAccClientVPNEndpoint_simpleAttributesUpdate,
			"selfServicePortal":            testAccClientVPNEndpoint_selfServicePortal,
			"vpcNoSecurityGroups":          testAccClientVPNEndpoint_vpcNoSecurityGroups,
			"vpcSecurityGroups":            testAccClientVPNEndpoint_vpcSecurityGroups,
			"basicDataSource":              testAccClientVPNEndpointDataSource_basic,
		},
		"AuthorizationRule": {
			acctest.CtBasic:      testAccClientVPNAuthorizationRule_basic,
			"groups":             testAccClientVPNAuthorizationRule_groups,
			names.AttrSubnets:    testAccClientVPNAuthorizationRule_subnets,
			"disappears":         testAccClientVPNAuthorizationRule_disappears,
			"disappearsEndpoint": testAccClientVPNAuthorizationRule_Disappears_endpoint,
		},
		"NetworkAssociation": {
			acctest.CtBasic:   testAccClientVPNNetworkAssociation_basic,
			"multipleSubnets": testAccClientVPNNetworkAssociation_multipleSubnets,
			"disappears":      testAccClientVPNNetworkAssociation_disappears,
		},
		"Route": {
			acctest.CtBasic:       testAccClientVPNRoute_basic,
			names.AttrDescription: testAccClientVPNRoute_description,
			"disappears":          testAccClientVPNRoute_disappears,
		},
	}

	acctest.RunLimitedConcurrencyTests2Levels(t, semaphore, testCases)
}

func testAccPreCheckClientVPNSyncronize(t *testing.T, semaphore tfsync.Semaphore) {
	tfsync.TestAccPreCheckSyncronize(t, semaphore, "Client VPN")
}
