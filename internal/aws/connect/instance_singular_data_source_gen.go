// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_instance", instanceDataSource)
}

// instanceDataSource returns the Terraform awscc_connect_instance data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::Instance resource.
func instanceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An instanceArn is automatically generated on creation based on instanceId.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An instanceArn is automatically generated on creation based on instanceId.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Attributes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The attributes for the instance.",
		//	  "properties": {
		//	    "AutoResolveBestVoices": {
		//	      "description": "Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.",
		//	      "type": "boolean"
		//	    },
		//	    "ContactLens": {
		//	      "description": "Boolean flag which enables CONTACT_LENS on an instance.",
		//	      "type": "boolean"
		//	    },
		//	    "ContactflowLogs": {
		//	      "description": "Boolean flag which enables CONTACTFLOW_LOGS on an instance.",
		//	      "type": "boolean"
		//	    },
		//	    "EarlyMedia": {
		//	      "description": "Boolean flag which enables EARLY_MEDIA on an instance.",
		//	      "type": "boolean"
		//	    },
		//	    "InboundCalls": {
		//	      "description": "Mandatory element which enables inbound calls on new instance.",
		//	      "type": "boolean"
		//	    },
		//	    "OutboundCalls": {
		//	      "description": "Mandatory element which enables outbound calls on new instance.",
		//	      "type": "boolean"
		//	    },
		//	    "UseCustomTTSVoices": {
		//	      "description": "Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "InboundCalls",
		//	    "OutboundCalls"
		//	  ],
		//	  "type": "object"
		//	}
		"attributes": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AutoResolveBestVoices
				"auto_resolve_best_voices": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag which enables AUTO_RESOLVE_BEST_VOICES on an instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ContactLens
				"contact_lens": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag which enables CONTACT_LENS on an instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ContactflowLogs
				"contactflow_logs": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag which enables CONTACTFLOW_LOGS on an instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: EarlyMedia
				"early_media": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag which enables EARLY_MEDIA on an instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: InboundCalls
				"inbound_calls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Mandatory element which enables inbound calls on new instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: OutboundCalls
				"outbound_calls": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Mandatory element which enables outbound calls on new instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UseCustomTTSVoices
				"use_custom_tts_voices": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Boolean flag which enables USE_CUSTOM_TTS_VOICES on an instance.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The attributes for the instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CreatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Timestamp of instance creation logged as part of instance creation.",
		//	  "format": "date-time",
		//	  "type": "string"
		//	}
		"created_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  timetypes.RFC3339Type{},
			Description: "Timestamp of instance creation logged as part of instance creation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DirectoryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Existing directoryId user wants to map to the new Connect instance.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^d-[0-9a-f]{10}$",
		//	  "type": "string"
		//	}
		"directory_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Existing directoryId user wants to map to the new Connect instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An instanceId is automatically generated on creation and assigned as the unique identifier.",
		//	  "type": "string"
		//	}
		"instance_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "An instanceId is automatically generated on creation and assigned as the unique identifier.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdentityManagementType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the type of directory integration for new instance.",
		//	  "enum": [
		//	    "SAML",
		//	    "CONNECT_MANAGED",
		//	    "EXISTING_DIRECTORY"
		//	  ],
		//	  "type": "string"
		//	}
		"identity_management_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the type of directory integration for new instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceAlias
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Alias of the new directory created as part of new instance creation.",
		//	  "maxLength": 62,
		//	  "minLength": 1,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"instance_alias": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Alias of the new directory created as part of new instance creation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the creation status of new instance.",
		//	  "enum": [
		//	    "CREATION_IN_PROGRESS",
		//	    "CREATION_FAILED",
		//	    "ACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"instance_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the creation status of new instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceRole
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Service linked role created as part of instance creation.",
		//	  "type": "string"
		//	}
		"service_role": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Service linked role created as part of instance creation.",
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
		Description: "Data Source schema for AWS::Connect::Instance",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::Instance").WithTerraformTypeName("awscc_connect_instance")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                      "Arn",
		"attributes":               "Attributes",
		"auto_resolve_best_voices": "AutoResolveBestVoices",
		"contact_lens":             "ContactLens",
		"contactflow_logs":         "ContactflowLogs",
		"created_time":             "CreatedTime",
		"directory_id":             "DirectoryId",
		"early_media":              "EarlyMedia",
		"identity_management_type": "IdentityManagementType",
		"inbound_calls":            "InboundCalls",
		"instance_alias":           "InstanceAlias",
		"instance_id":              "Id",
		"instance_status":          "InstanceStatus",
		"key":                      "Key",
		"outbound_calls":           "OutboundCalls",
		"service_role":             "ServiceRole",
		"tags":                     "Tags",
		"use_custom_tts_voices":    "UseCustomTTSVoices",
		"value":                    "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
