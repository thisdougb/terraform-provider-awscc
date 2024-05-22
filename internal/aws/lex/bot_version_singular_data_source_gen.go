// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lex

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lex_bot_version", botVersionDataSource)
}

// botVersionDataSource returns the Terraform awscc_lex_bot_version data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lex::BotVersion resource.
func botVersionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Unique ID of resource",
		//	  "maxLength": 10,
		//	  "minLength": 10,
		//	  "pattern": "^[0-9a-zA-Z]+$",
		//	  "type": "string"
		//	}
		"bot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Unique ID of resource",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The version of a bot.",
		//	  "maxLength": 5,
		//	  "minLength": 1,
		//	  "pattern": "^(DRAFT|[0-9]+)$",
		//	  "type": "string"
		//	}
		"bot_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The version of a bot.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BotVersionLocaleSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "BotVersionLocaleDetails": {
		//	        "additionalProperties": false,
		//	        "description": "The version of a bot used for a bot locale.",
		//	        "properties": {
		//	          "SourceBotVersion": {
		//	            "description": "The version of a bot.",
		//	            "maxLength": 5,
		//	            "minLength": 1,
		//	            "pattern": "^(DRAFT|[0-9]+)$",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "SourceBotVersion"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "LocaleId": {
		//	        "description": "The identifier of the language and locale that the bot will be used in.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "LocaleId",
		//	      "BotVersionLocaleDetails"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"bot_version_locale_specification": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: BotVersionLocaleDetails
					"bot_version_locale_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SourceBotVersion
							"source_bot_version": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The version of a bot.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The version of a bot used for a bot locale.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: LocaleId
					"locale_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The identifier of the language and locale that the bot will be used in.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Specifies the locales that Amazon Lex adds to this version. You can choose the Draft version or any other previously published version for each locale.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the version. Use the description to help identify the version in lists.",
		//	  "maxLength": 200,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the version. Use the description to help identify the version in lists.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lex::BotVersion",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lex::BotVersion").WithTerraformTypeName("awscc_lex_bot_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bot_id":                           "BotId",
		"bot_version":                      "BotVersion",
		"bot_version_locale_details":       "BotVersionLocaleDetails",
		"bot_version_locale_specification": "BotVersionLocaleSpecification",
		"description":                      "Description",
		"locale_id":                        "LocaleId",
		"source_bot_version":               "SourceBotVersion",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
