// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_transit_gateway_multicast_group_source", transitGatewayMulticastGroupSourceDataSource)
}

// transitGatewayMulticastGroupSourceDataSource returns the Terraform awscc_ec2_transit_gateway_multicast_group_source data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::TransitGatewayMulticastGroupSource resource.
func transitGatewayMulticastGroupSourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupIpAddress
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IP address assigned to the transit gateway multicast group.",
		//	  "type": "string"
		//	}
		"group_ip_address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IP address assigned to the transit gateway multicast group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupMember
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that the resource is a transit gateway multicast group member.",
		//	  "type": "boolean"
		//	}
		"group_member": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GroupSource
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates that the resource is a transit gateway multicast group member.",
		//	  "type": "boolean"
		//	}
		"group_source": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates that the resource is a transit gateway multicast group member.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInterfaceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway attachment.",
		//	  "type": "string"
		//	}
		"network_interface_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the resource.",
		//	  "type": "string"
		//	}
		"resource_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of resource, for example a VPC attachment.",
		//	  "type": "string"
		//	}
		"resource_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of resource, for example a VPC attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source type.",
		//	  "type": "string"
		//	}
		"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SubnetId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the subnet.",
		//	  "type": "string"
		//	}
		"subnet_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the subnet.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayAttachmentId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway attachment.",
		//	  "type": "string"
		//	}
		"transit_gateway_attachment_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway attachment.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TransitGatewayMulticastDomainId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the transit gateway multicast domain.",
		//	  "type": "string"
		//	}
		"transit_gateway_multicast_domain_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the transit gateway multicast domain.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::TransitGatewayMulticastGroupSource",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::TransitGatewayMulticastGroupSource").WithTerraformTypeName("awscc_ec2_transit_gateway_multicast_group_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_ip_address":                    "GroupIpAddress",
		"group_member":                        "GroupMember",
		"group_source":                        "GroupSource",
		"network_interface_id":                "NetworkInterfaceId",
		"resource_id":                         "ResourceId",
		"resource_type":                       "ResourceType",
		"source_type":                         "SourceType",
		"subnet_id":                           "SubnetId",
		"transit_gateway_attachment_id":       "TransitGatewayAttachmentId",
		"transit_gateway_multicast_domain_id": "TransitGatewayMulticastDomainId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
