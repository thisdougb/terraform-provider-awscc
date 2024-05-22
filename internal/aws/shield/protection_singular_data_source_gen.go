// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package shield

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_shield_protection", protectionDataSource)
}

// protectionDataSource returns the Terraform awscc_shield_protection data source.
// This Terraform data source corresponds to the CloudFormation AWS::Shield::Protection resource.
func protectionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: ApplicationLayerAutomaticResponseConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.",
		//	  "properties": {
		//	    "Action": {
		//	      "description": "Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.",
		//	      "properties": {
		//	        "Block": {
		//	          "additionalProperties": false,
		//	          "description": "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.\nYou must specify exactly one action, either `Block` or `Count`.",
		//	          "type": "object"
		//	        },
		//	        "Count": {
		//	          "additionalProperties": false,
		//	          "description": "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.\nYou must specify exactly one action, either `Block` or `Count`.",
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Status": {
		//	      "description": "Indicates whether automatic application layer DDoS mitigation is enabled for the protection.",
		//	      "enum": [
		//	        "ENABLED",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Action",
		//	    "Status"
		//	  ],
		//	  "type": "object"
		//	}
		"application_layer_automatic_response_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Action
				"action": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Block
						"block": schema.StringAttribute{ /*START ATTRIBUTE*/
							CustomType:  jsontypes.NormalizedType{},
							Description: "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Block` action.\nYou must specify exactly one action, either `Block` or `Count`.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: Count
						"count": schema.StringAttribute{ /*START ATTRIBUTE*/
							CustomType:  jsontypes.NormalizedType{},
							Description: "Specifies that Shield Advanced should configure its AWS WAF rules with the AWS WAF `Count` action.\nYou must specify exactly one action, either `Block` or `Count`.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Specifies the action setting that Shield Advanced should use in the AWS WAF rules that it creates on behalf of the protected resource in response to DDoS attacks. You specify this as part of the configuration for the automatic application layer DDoS mitigation feature, when you enable or update automatic mitigation. Shield Advanced creates the AWS WAF rules in a Shield Advanced-managed rule group, inside the web ACL that you have associated with the resource.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Status
				"status": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Indicates whether automatic application layer DDoS mitigation is enabled for the protection.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The automatic application layer DDoS mitigation settings for a Protection. This configuration determines whether Shield Advanced automatically manages rules in the web ACL in order to respond to application layer events that Shield Advanced determines to be DDoS attacks.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HealthCheckArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of the health check to associate with the protection.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 2048,
		//	    "minLength": 1,
		//	    "type": "string"
		//	  },
		//	  "maxItems": 1,
		//	  "type": "array"
		//	}
		"health_check_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Names (ARNs) of the health check to associate with the protection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Friendly name for the Protection.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[ a-zA-Z0-9_\\.\\-]*",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Friendly name for the Protection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProtectionArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon Resource Name) of the protection.",
		//	  "type": "string"
		//	}
		"protection_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon Resource Name) of the protection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProtectionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique identifier (ID) of the protection.",
		//	  "type": "string"
		//	}
		"protection_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique identifier (ID) of the protection.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN (Amazon Resource Name) of the resource to be protected.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN (Amazon Resource Name) of the resource to be protected.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tag key-value pairs for the Protection object.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A tag associated with an AWS resource. Tags are key:value pairs that you can use to categorize and manage your resources, for purposes like billing or other management. Typically, the tag key represents a category, such as \"environment\", and the tag value represents a specific value within that category, such as \"test,\" \"development,\" or \"production\". Or you might set the tag key to \"customer\" and the value to the customer name or ID. You can specify one or more tags to add to each AWS resource, up to 50 tags for a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
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
		//	  "maxItems": 200,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Part of the key:value pair that defines a tag. You can use a tag key to describe a category of information, such as \"customer.\" Tag keys are case-sensitive.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "Part of the key:value pair that defines a tag. You can use a tag value to describe a specific value within a category, such as \"companyA\" or \"companyB.\" Tag values are case-sensitive.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tag key-value pairs for the Protection object.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Shield::Protection",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Shield::Protection").WithTerraformTypeName("awscc_shield_protection")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action": "Action",
		"application_layer_automatic_response_configuration": "ApplicationLayerAutomaticResponseConfiguration",
		"block":             "Block",
		"count":             "Count",
		"health_check_arns": "HealthCheckArns",
		"key":               "Key",
		"name":              "Name",
		"protection_arn":    "ProtectionArn",
		"protection_id":     "ProtectionId",
		"resource_arn":      "ResourceArn",
		"status":            "Status",
		"tags":              "Tags",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
