// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourceexplorer2_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccResourceExplorer2_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Index": {
			acctest.CtBasic: testAccIndex_basic,
			"disappears":    testAccIndex_disappears,
			names.AttrTags:  testAccIndex_tags,
			names.AttrType:  testAccIndex_type,
		},
		"View": {
			acctest.CtBasic:  testAccView_basic,
			"defaultView":    testAccView_defaultView,
			"disappears":     testAccView_disappears,
			names.AttrFilter: testAccView_filter,
			names.AttrTags:   testAccView_tags,
		},
		"SearchDataSource": {
			acctest.CtBasic: testAccSearchDataSource_basic,
			"indexType":     testAccSearchDataSource_IndexType,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
