// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53resolver

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_route53resolver_outpost_resolver", outpostResolverDataSource)
}

// outpostResolverDataSource returns the Terraform awscc_route53resolver_outpost_resolver data source.
// This Terraform data source corresponds to the CloudFormation AWS::Route53Resolver::OutpostResolver resource.
func outpostResolverDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver ARN.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver creation time",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver creation time",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatorRequestId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The id of the creator request.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"creator_request_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The id of the creator request.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Id",
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"outpost_resolver_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Id",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of OutpostResolvers.",
		//	  "maximum": 256,
		//	  "minimum": 4,
		//	  "type": "integer"
		//	}
		"instance_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of OutpostResolvers.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModificationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver last modified time",
		//	  "maxLength": 40,
		//	  "minLength": 20,
		//	  "type": "string"
		//	}
		"modification_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver last modified time",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver name.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver name.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Outpost ARN.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Outpost ARN.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PreferredInstanceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver instance type.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"preferred_instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver instance type.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver status, possible values are CREATING, OPERATIONAL, UPDATING, DELETING, ACTION_NEEDED, FAILED_CREATION and FAILED_DELETION.",
		//	  "enum": [
		//	    "CREATING",
		//	    "OPERATIONAL",
		//	    "DELETING",
		//	    "UPDATING",
		//	    "ACTION_NEEDED",
		//	    "FAILED_CREATION",
		//	    "FAILED_DELETION"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver status, possible values are CREATING, OPERATIONAL, UPDATING, DELETING, ACTION_NEEDED, FAILED_CREATION and FAILED_DELETION.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: StatusMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The OutpostResolver status message.",
		//	  "type": "string"
		//	}
		"status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The OutpostResolver status message.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
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
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Route53Resolver::OutpostResolver",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53Resolver::OutpostResolver").WithTerraformTypeName("awscc_route53resolver_outpost_resolver")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                     "Arn",
		"creation_time":           "CreationTime",
		"creator_request_id":      "CreatorRequestId",
		"instance_count":          "InstanceCount",
		"key":                     "Key",
		"modification_time":       "ModificationTime",
		"name":                    "Name",
		"outpost_arn":             "OutpostArn",
		"outpost_resolver_id":     "Id",
		"preferred_instance_type": "PreferredInstanceType",
		"status":                  "Status",
		"status_message":          "StatusMessage",
		"tags":                    "Tags",
		"value":                   "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
