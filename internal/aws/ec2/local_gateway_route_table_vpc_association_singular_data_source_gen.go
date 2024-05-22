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
	registry.AddDataSourceFactory("awscc_ec2_local_gateway_route_table_vpc_association", localGatewayRouteTableVPCAssociationDataSource)
}

// localGatewayRouteTableVPCAssociationDataSource returns the Terraform awscc_ec2_local_gateway_route_table_vpc_association data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::LocalGatewayRouteTableVPCAssociation resource.
func localGatewayRouteTableVPCAssociationDataSource(ctx context.Context) (datasource.DataSource, error) {
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
		// Property: LocalGatewayRouteTableVpcAssociationId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the association.",
		//	  "type": "string"
		//	}
		"local_gateway_route_table_vpc_association_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The state of the association.",
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The state of the association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags for the association.",
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
			Description: "The tags for the association.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VpcId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the VPC.",
		//	  "type": "string"
		//	}
		"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the VPC.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::LocalGatewayRouteTableVPCAssociation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::LocalGatewayRouteTableVPCAssociation").WithTerraformTypeName("awscc_ec2_local_gateway_route_table_vpc_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"key":                          "Key",
		"local_gateway_id":             "LocalGatewayId",
		"local_gateway_route_table_id": "LocalGatewayRouteTableId",
		"local_gateway_route_table_vpc_association_id": "LocalGatewayRouteTableVpcAssociationId",
		"state":  "State",
		"tags":   "Tags",
		"value":  "Value",
		"vpc_id": "VpcId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
