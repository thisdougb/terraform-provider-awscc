// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package cloudformation

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_cloudformation_module_version", moduleVersionDataSource)
}

// moduleVersionDataSource returns the Terraform awscc_cloudformation_module_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::CloudFormation::ModuleVersion resource.
func moduleVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the module.",
		//	  "pattern": "^arn:aws[A-Za-z0-9-]{0,64}:cloudformation:[A-Za-z0-9-]{1,64}:([0-9]{12})?:type/module/.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the module.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the registered module.",
		//	  "maxLength": 1024,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the registered module.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DocumentationUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of a page providing detailed documentation for this module.",
		//	  "maxLength": 4096,
		//	  "type": "string"
		//	}
		"documentation_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of a page providing detailed documentation for this module.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IsDefaultVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicator of whether this module version is the current default version",
		//	  "type": "boolean"
		//	}
		"is_default_version": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicator of whether this module version is the current default version",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModuleName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the module being registered.\n\nRecommended module naming pattern: company_or_organization::service::type::MODULE.",
		//	  "pattern": "^[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::[A-Za-z0-9]{2,64}::MODULE",
		//	  "type": "string"
		//	}
		"module_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the module being registered.\n\nRecommended module naming pattern: company_or_organization::service::type::MODULE.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ModulePackage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The url to the S3 bucket containing the schema and template fragment for the module you want to register.",
		//	  "type": "string"
		//	}
		"module_package": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The url to the S3 bucket containing the schema and template fragment for the module you want to register.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Schema
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The schema defining input parameters to and resources generated by the module.",
		//	  "maxLength": 16777216,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"schema": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The schema defining input parameters to and resources generated by the module.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TimeCreated
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time that the specified module version was registered.",
		//	  "type": "string"
		//	}
		"time_created": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time that the specified module version was registered.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VersionId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version ID of the module represented by this module instance.",
		//	  "pattern": "^[0-9]{8}$",
		//	  "type": "string"
		//	}
		"version_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version ID of the module represented by this module instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Visibility
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The scope at which the type is visible and usable in CloudFormation operations.\n\nThe only allowed value at present is:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.",
		//	  "enum": [
		//	    "PRIVATE"
		//	  ],
		//	  "type": "string"
		//	}
		"visibility": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The scope at which the type is visible and usable in CloudFormation operations.\n\nThe only allowed value at present is:\n\nPRIVATE: The type is only visible and usable within the account in which it is registered. Currently, AWS CloudFormation marks any types you register as PRIVATE.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::CloudFormation::ModuleVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CloudFormation::ModuleVersion").WithTerraformTypeName("awscc_cloudformation_module_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                "Arn",
		"description":        "Description",
		"documentation_url":  "DocumentationUrl",
		"is_default_version": "IsDefaultVersion",
		"module_name":        "ModuleName",
		"module_package":     "ModulePackage",
		"schema":             "Schema",
		"time_created":       "TimeCreated",
		"version_id":         "VersionId",
		"visibility":         "Visibility",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
