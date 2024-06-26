// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKDataSource("aws_ec2_coip_pool")
func DataSourceCoIPPool() *schema.Resource {
	return &schema.Resource{
		ReadWithoutTimeout: dataSourceCoIPPoolRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"local_gateway_route_table_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"pool_cidrs": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
				Set:      schema.HashString,
			},

			"pool_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			names.AttrARN: {
				Type:     schema.TypeString,
				Computed: true,
			},

			names.AttrTags: tftags.TagsSchemaComputed(),

			names.AttrFilter: customFiltersSchema(),
		},
	}
}

func dataSourceCoIPPoolRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	req := &ec2.DescribeCoipPoolsInput{}

	if v, ok := d.GetOk("pool_id"); ok {
		req.PoolIds = []*string{aws.String(v.(string))}
	}

	filters := map[string]string{}

	if v, ok := d.GetOk("local_gateway_route_table_id"); ok {
		filters["coip-pool.local-gateway-route-table-id"] = v.(string)
	}

	req.Filters = newAttributeFilterList(filters)

	if tags, tagsOk := d.GetOk(names.AttrTags); tagsOk {
		req.Filters = append(req.Filters, newTagFilterList(
			Tags(tftags.New(ctx, tags.(map[string]interface{}))),
		)...)
	}

	req.Filters = append(req.Filters, newCustomFilterList(
		d.Get(names.AttrFilter).(*schema.Set),
	)...)
	if len(req.Filters) == 0 {
		// Don't send an empty filters list; the EC2 API won't accept it.
		req.Filters = nil
	}

	log.Printf("[DEBUG] Reading AWS COIP Pool: %s", req)
	resp, err := conn.DescribeCoipPoolsWithContext(ctx, req)
	if err != nil {
		return sdkdiag.AppendErrorf(diags, "describing EC2 COIP Pools: %s", err)
	}
	if resp == nil || len(resp.CoipPools) == 0 {
		return sdkdiag.AppendErrorf(diags, "no matching COIP Pool found")
	}
	if len(resp.CoipPools) > 1 {
		return sdkdiag.AppendErrorf(diags, "multiple Coip Pools matched; use additional constraints to reduce matches to a single COIP Pool")
	}

	coip := resp.CoipPools[0]

	d.SetId(aws.StringValue(coip.PoolId))

	d.Set("local_gateway_route_table_id", coip.LocalGatewayRouteTableId)
	d.Set(names.AttrARN, coip.PoolArn)

	if err := d.Set("pool_cidrs", aws.StringValueSlice(coip.PoolCidrs)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting pool_cidrs: %s", err)
	}

	d.Set("pool_id", coip.PoolId)

	if err := d.Set(names.AttrTags, KeyValueTags(ctx, coip.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	return diags
}
