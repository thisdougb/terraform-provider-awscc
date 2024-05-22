// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_mediatailor_live_source", liveSourceResource)
}

// liveSourceResource returns the Terraform awscc_mediatailor_live_source resource.
// This Terraform resource corresponds to the CloudFormation AWS::MediaTailor::LiveSource resource.
func liveSourceResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eThe ARN of the live source.\u003c/p\u003e",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "<p>The ARN of the live source.</p>",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: HttpPackageConfigurations
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "\u003cp\u003eA list of HTTP package configuration parameters for this live source.\u003c/p\u003e",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "\u003cp\u003eThe HTTP package configuration properties for the requested VOD source.\u003c/p\u003e",
		//	    "properties": {
		//	      "Path": {
		//	        "description": "\u003cp\u003eThe relative path to the URL for this VOD source. This is combined with \u003ccode\u003eSourceLocation::HttpConfiguration::BaseUrl\u003c/code\u003e to form a valid URL.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "SourceGroup": {
		//	        "description": "\u003cp\u003eThe name of the source group. This has to match one of the \u003ccode\u003eChannel::Outputs::SourceGroup\u003c/code\u003e.\u003c/p\u003e",
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "DASH",
		//	          "HLS"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Path",
		//	      "SourceGroup",
		//	      "Type"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"http_package_configurations": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Path
					"path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The relative path to the URL for this VOD source. This is combined with <code>SourceLocation::HttpConfiguration::BaseUrl</code> to form a valid URL.</p>",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: SourceGroup
					"source_group": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "<p>The name of the source group. This has to match one of the <code>Channel::Outputs::SourceGroup</code>.</p>",
						Required:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"DASH",
								"HLS",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "<p>A list of HTTP package configuration parameters for this live source.</p>",
			Required:    true,
		}, /*END ATTRIBUTE*/
		// Property: LiveSourceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"live_source_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceLocationName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"source_location_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to assign to the live source.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to assign to the live source.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::MediaTailor::LiveSource Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaTailor::LiveSource").WithTerraformTypeName("awscc_mediatailor_live_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                         "Arn",
		"http_package_configurations": "HttpPackageConfigurations",
		"key":                         "Key",
		"live_source_name":            "LiveSourceName",
		"path":                        "Path",
		"source_group":                "SourceGroup",
		"source_location_name":        "SourceLocationName",
		"tags":                        "Tags",
		"type":                        "Type",
		"value":                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
