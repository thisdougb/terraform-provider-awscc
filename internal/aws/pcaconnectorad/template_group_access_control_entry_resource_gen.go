// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_pcaconnectorad_template_group_access_control_entry", templateGroupAccessControlEntryResource)
}

// templateGroupAccessControlEntryResource returns the Terraform awscc_pcaconnectorad_template_group_access_control_entry resource.
// This Terraform resource corresponds to the CloudFormation AWS::PCAConnectorAD::TemplateGroupAccessControlEntry resource.
func templateGroupAccessControlEntryResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccessRights
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "AutoEnroll": {
		//	      "enum": [
		//	        "ALLOW",
		//	        "DENY"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Enroll": {
		//	      "enum": [
		//	        "ALLOW",
		//	        "DENY"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"access_rights": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoEnroll
				"auto_enroll": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ALLOW",
							"DENY",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Enroll
				"enroll": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ALLOW",
							"DENY",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			// AccessRights is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: GroupDisplayName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 0,
		//	  "pattern": "^[\\x20-\\x7E]+$",
		//	  "type": "string"
		//	}
		"group_display_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\x20-\\x7E]+$"), ""),
			}, /*END VALIDATORS*/
			// GroupDisplayName is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: GroupSecurityIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 256,
		//	  "minLength": 7,
		//	  "pattern": "^S-[0-9]-([0-9]+-){1,14}[0-9]+$",
		//	  "type": "string"
		//	}
		"group_security_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(7, 256),
				stringvalidator.RegexMatches(regexp.MustCompile("^S-[0-9]-([0-9]+-){1,14}[0-9]+$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TemplateArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)\\/template(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"template_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)\\/template(\\/[\\w-]+)$"), ""),
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
		Description: "Definition of AWS::PCAConnectorAD::TemplateGroupAccessControlEntry Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCAConnectorAD::TemplateGroupAccessControlEntry").WithTerraformTypeName("awscc_pcaconnectorad_template_group_access_control_entry")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"access_rights":             "AccessRights",
		"auto_enroll":               "AutoEnroll",
		"enroll":                    "Enroll",
		"group_display_name":        "GroupDisplayName",
		"group_security_identifier": "GroupSecurityIdentifier",
		"template_arn":              "TemplateArn",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/AccessRights",
		"/properties/GroupDisplayName",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
