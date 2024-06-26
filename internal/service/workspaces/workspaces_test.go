// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workspaces_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccWorkSpaces_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Directory": {
			acctest.CtBasic:               testAccDirectory_basic,
			"disappears":                  testAccDirectory_disappears,
			"ipGroupIds":                  testAccDirectory_ipGroupIDs,
			"selfServicePermissions":      testAccDirectory_selfServicePermissions,
			"subnetIDs":                   testAccDirectory_subnetIDs,
			names.AttrTags:                testAccDirectory_tags,
			"workspaceAccessProperties":   testAccDirectory_workspaceAccessProperties,
			"workspaceCreationProperties": testAccDirectory_workspaceCreationProperties,
			"workspaceCreationProperties_customSecurityGroupId_defaultOu": testAccDirectory_workspaceCreationProperties_customSecurityGroupId_defaultOu,
		},
		"IpGroup": {
			acctest.CtBasic:       testAccIPGroup_basic,
			"disappears":          testAccIPGroup_disappears,
			"multipleDirectories": testAccIPGroup_MultipleDirectories,
			names.AttrTags:        testAccIPGroup_tags,
		},
		"Workspace": {
			acctest.CtBasic:                           testAccWorkspace_basic,
			"recreate":                                testAccWorkspace_recreate,
			names.AttrTags:                            testAccWorkspace_tags,
			names.AttrTimeout:                         testAccWorkspace_timeout,
			"validateRootVolumeSize":                  testAccWorkspace_validateRootVolumeSize,
			"validateUserVolumeSize":                  testAccWorkspace_validateUserVolumeSize,
			"workspaceProperties":                     testAccWorkspace_workspaceProperties,
			"workspaceProperties_runningModeAlwaysOn": testAccWorkspace_workspaceProperties_runningModeAlwaysOn,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
