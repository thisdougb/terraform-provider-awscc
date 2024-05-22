// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package connect

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
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
	registry.AddResourceFactory("awscc_connect_phone_number", phoneNumberResource)
}

// phoneNumberResource returns the Terraform awscc_connect_phone_number resource.
// This Terraform resource corresponds to the CloudFormation AWS::Connect::PhoneNumber resource.
func phoneNumberResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Address
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The phone number e164 address.",
		//	  "pattern": "^\\+[0-9]{2,15}",
		//	  "type": "string"
		//	}
		"address": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The phone number e164 address.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CountryCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The phone number country code.",
		//	  "pattern": "^[A-Z]{2}",
		//	  "type": "string"
		//	}
		"country_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The phone number country code.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[A-Z]{2}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the phone number.",
		//	  "maxLength": 500,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the phone number.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 500),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PhoneNumberArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The phone number ARN",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:phone-number/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"phone_number_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The phone number ARN",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Prefix
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The phone number prefix.",
		//	  "pattern": "^\\+[0-9]{1,15}",
		//	  "type": "string"
		//	}
		"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The phone number prefix.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^\\+[0-9]{1,15}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// Prefix is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: SourcePhoneNumberArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source phone number arn.",
		//	  "type": "string"
		//	}
		"source_phone_number_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source phone number arn.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "One or more tags.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Required:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthAtMost(256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the target the phone number is claimed to.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:(instance|traffic-distribution-group)/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"target_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the target the phone number is claimed to.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:(instance|traffic-distribution-group)/[-a-zA-Z0-9]*$"), ""),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Type
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The phone number type",
		//	  "pattern": "TOLL_FREE|DID|UIFN|SHARED|THIRD_PARTY_DID|THIRD_PARTY_TF|SHORT_CODE",
		//	  "type": "string"
		//	}
		"type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The phone number type",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("TOLL_FREE|DID|UIFN|SHARED|THIRD_PARTY_DID|THIRD_PARTY_TF|SHORT_CODE"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "Resource Type definition for AWS::Connect::PhoneNumber",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::PhoneNumber").WithTerraformTypeName("awscc_connect_phone_number")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":                 "Address",
		"country_code":            "CountryCode",
		"description":             "Description",
		"key":                     "Key",
		"phone_number_arn":        "PhoneNumberArn",
		"prefix":                  "Prefix",
		"source_phone_number_arn": "SourcePhoneNumberArn",
		"tags":                    "Tags",
		"target_arn":              "TargetArn",
		"type":                    "Type",
		"value":                   "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Prefix",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
