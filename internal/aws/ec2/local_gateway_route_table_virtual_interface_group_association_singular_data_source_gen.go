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
	registry.AddDataSourceFactory("awscc_ec2_local_gateway_route_table_virtual_interface_group_association", localGatewayRouteTableVirtualInterfaceGroupAssociationDataSource)
}

// localGatewayRouteTableVirtualInterfaceGroupAssociationDataSource returns the Terraform awscc_ec2_local_gateway_route_table_virtual_interface_group_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation resource.
func localGatewayRouteTableVirtualInterfaceGroupAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: LocalGatewayId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway.",
		//	  "type": "string"
		//	}
		"local_gateway_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayRouteTableArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the local gateway route table.",
		//	  "type": "string"
		//	}
		"local_gateway_route_table_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the local gateway route table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayRouteTableId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway route table.",
		//	  "type": "string"
		//	}
		"local_gateway_route_table_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway route table.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayRouteTableVirtualInterfaceGroupAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway route table virtual interface group association.",
		//	  "type": "string"
		//	}
		"local_gateway_route_table_virtual_interface_group_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway route table virtual interface group association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LocalGatewayVirtualInterfaceGroupId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the local gateway route table virtual interface group.",
		//	  "type": "string"
		//	}
		"local_gateway_virtual_interface_group_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the local gateway route table virtual interface group.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OwnerId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The owner of the local gateway route table virtual interface group association.",
		//	  "type": "string"
		//	}
		"owner_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The owner of the local gateway route table virtual interface group association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the local gateway route table virtual interface group association.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the local gateway route table virtual interface group association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the local gateway route table virtual interface group association.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 127,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 255,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags for the local gateway route table virtual interface group association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVirtualInterfaceGroupAssociation").WithTerraformTypeName("awscc_ec2_local_gateway_route_table_virtual_interface_group_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                           "Key",
		"local_gateway_id":              "LocalGatewayId",
		"local_gateway_route_table_arn": "LocalGatewayRouteTableArn",
		"local_gateway_route_table_id":  "LocalGatewayRouteTableId",
		"local_gateway_route_table_virtual_interface_group_association_id": "LocalGatewayRouteTableVirtualInterfaceGroupAssociationId",
		"local_gateway_virtual_interface_group_id":                         "LocalGatewayVirtualInterfaceGroupId",
		"owner_id": "OwnerId",
		"state":    "State",
		"tags":     "Tags",
		"value":    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
