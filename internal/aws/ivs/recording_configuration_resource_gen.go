// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ivs

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ivs_recording_configuration", recordingConfigurationResource)
}

// recordingConfigurationResource returns the Terraform awscc_ivs_recording_configuration resource.
// This Terraform resource corresponds to the CloudFormation AWS::IVS::RecordingConfiguration resource.
func recordingConfigurationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^arn:aws[-a-z]*:ivs:[a-z0-9-]+:[0-9]+:recording-configuration/[a-zA-Z0-9-]+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Recording Configuration ARN is automatically generated on creation and assigned as the unique identifier.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DestinationConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Recording Destination Configuration.",
		//	  "properties": {
		//	    "S3": {
		//	      "additionalProperties": false,
		//	      "description": "Recording S3 Destination Configuration.",
		//	      "properties": {
		//	        "BucketName": {
		//	          "maxLength": 63,
		//	          "minLength": 3,
		//	          "pattern": "^[a-z0-9-.]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "BucketName"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"destination_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: S3
				"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: BucketName
						"bucket_name": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(3, 63),
								stringvalidator.RegexMatches(regexp.MustCompile("^[a-z0-9-.]+$"), ""),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.RequiresReplace(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Recording S3 Destination Configuration.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
						objectplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Recording Destination Configuration.",
			Required:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Recording Configuration Name.",
		//	  "maxLength": 128,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z0-9-_]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Recording Configuration Name.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(0, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9-_]*$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RecordingReconnectWindowSeconds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": 0,
		//	  "description": "Recording Reconnect Window Seconds. (0 means disabled)",
		//	  "maximum": 300,
		//	  "minimum": 0,
		//	  "type": "integer"
		//	}
		"recording_reconnect_window_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Recording Reconnect Window Seconds. (0 means disabled)",
			Optional:    true,
			Computed:    true,
			Default:     int64default.StaticInt64(0),
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(0, 300),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
				int64planmodifier.UseStateForUnknown(),
				int64planmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RenditionConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Rendition Configuration describes which renditions should be recorded for a stream.",
		//	  "properties": {
		//	    "RenditionSelection": {
		//	      "default": "ALL",
		//	      "description": "Resolution Selection indicates which set of renditions are recorded for a stream.",
		//	      "enum": [
		//	        "ALL",
		//	        "NONE",
		//	        "CUSTOM"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Renditions": {
		//	      "description": "Renditions indicates which renditions are recorded for a stream.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "FULL_HD",
		//	          "HD",
		//	          "SD",
		//	          "LOWEST_RESOLUTION"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 4,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"rendition_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RenditionSelection
				"rendition_selection": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Resolution Selection indicates which set of renditions are recorded for a stream.",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString("ALL"),
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ALL",
							"NONE",
							"CUSTOM",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Renditions
				"renditions": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Renditions indicates which renditions are recorded for a stream.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 4),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"FULL_HD",
								"HD",
								"SD",
								"LOWEST_RESOLUTION",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
						setplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Rendition Configuration describes which renditions should be recorded for a stream.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Recording Configuration State.",
		//	  "enum": [
		//	    "CREATING",
		//	    "CREATE_FAILED",
		//	    "ACTIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Recording Configuration State.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of key-value pairs that contain metadata for the asset model.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
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
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of key-value pairs that contain metadata for the asset model.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ThumbnailConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Recording Thumbnail Configuration.",
		//	  "properties": {
		//	    "RecordingMode": {
		//	      "default": "INTERVAL",
		//	      "description": "Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.",
		//	      "enum": [
		//	        "INTERVAL",
		//	        "DISABLED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Resolution": {
		//	      "description": "Resolution indicates the desired resolution of recorded thumbnails.",
		//	      "enum": [
		//	        "FULL_HD",
		//	        "HD",
		//	        "SD",
		//	        "LOWEST_RESOLUTION"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "Storage": {
		//	      "description": "Storage indicates the format in which thumbnails are recorded.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "enum": [
		//	          "SEQUENTIAL",
		//	          "LATEST"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "maxItems": 2,
		//	      "minItems": 0,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "TargetIntervalSeconds": {
		//	      "default": 60,
		//	      "description": "Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.",
		//	      "maximum": 60,
		//	      "minimum": 1,
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"thumbnail_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RecordingMode
				"recording_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Thumbnail Recording Mode, which determines whether thumbnails are recorded at an interval or are disabled.",
					Optional:    true,
					Computed:    true,
					Default:     stringdefault.StaticString("INTERVAL"),
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"INTERVAL",
							"DISABLED",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Resolution
				"resolution": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Resolution indicates the desired resolution of recorded thumbnails.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"FULL_HD",
							"HD",
							"SD",
							"LOWEST_RESOLUTION",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Storage
				"storage": schema.SetAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "Storage indicates the format in which thumbnails are recorded.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Set{ /*START VALIDATORS*/
						setvalidator.SizeBetween(0, 2),
						setvalidator.ValueStringsAre(
							stringvalidator.OneOf(
								"SEQUENTIAL",
								"LATEST",
							),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
						setplanmodifier.UseStateForUnknown(),
						setplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TargetIntervalSeconds
				"target_interval_seconds": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Description: "Target Interval Seconds defines the interval at which thumbnails are recorded. This field is required if RecordingMode is INTERVAL.",
					Optional:    true,
					Computed:    true,
					Default:     int64default.StaticInt64(60),
					Validators: []validator.Int64{ /*START VALIDATORS*/
						int64validator.Between(1, 60),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
						int64planmodifier.UseStateForUnknown(),
						int64planmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Recording Thumbnail Configuration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
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
		Description: "Resource Type definition for AWS::IVS::RecordingConfiguration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IVS::RecordingConfiguration").WithTerraformTypeName("awscc_ivs_recording_configuration")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                                "Arn",
		"bucket_name":                        "BucketName",
		"destination_configuration":          "DestinationConfiguration",
		"key":                                "Key",
		"name":                               "Name",
		"recording_mode":                     "RecordingMode",
		"recording_reconnect_window_seconds": "RecordingReconnectWindowSeconds",
		"rendition_configuration":            "RenditionConfiguration",
		"rendition_selection":                "RenditionSelection",
		"renditions":                         "Renditions",
		"resolution":                         "Resolution",
		"s3":                                 "S3",
		"state":                              "State",
		"storage":                            "Storage",
		"tags":                               "Tags",
		"target_interval_seconds":            "TargetIntervalSeconds",
		"thumbnail_configuration":            "ThumbnailConfiguration",
		"value":                              "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
