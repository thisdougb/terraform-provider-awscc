// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package osis

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_osis_pipeline", pipelineResource)
}

// pipelineResource returns the Terraform awscc_osis_pipeline resource.
// This Terraform resource corresponds to the CloudFormation AWS::OSIS::Pipeline resource.
func pipelineResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BufferOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Key-value pairs to configure buffering.",
		//	  "properties": {
		//	    "PersistentBufferEnabled": {
		//	      "description": "Whether persistent buffering should be enabled.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "required": [
		//	    "PersistentBufferEnabled"
		//	  ],
		//	  "type": "object"
		//	}
		"buffer_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PersistentBufferEnabled
				"persistent_buffer_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether persistent buffering should be enabled.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.Bool{ /*START VALIDATORS*/
						fwvalidators.NotNullBool(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Key-value pairs to configure buffering.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EncryptionAtRestOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Key-value pairs to configure encryption at rest.",
		//	  "properties": {
		//	    "KmsKeyArn": {
		//	      "description": "The KMS key to use for encrypting data. By default an AWS owned key is used",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "KmsKeyArn"
		//	  ],
		//	  "type": "object"
		//	}
		"encryption_at_rest_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The KMS key to use for encrypting data. By default an AWS owned key is used",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						fwvalidators.NotNullString(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Key-value pairs to configure encryption at rest.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IngestEndpointUrls
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of endpoints that can be used for ingesting data into a pipeline",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"ingest_endpoint_urls": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of endpoints that can be used for ingesting data into a pipeline",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LogPublishingOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Key-value pairs to configure log publishing.",
		//	  "properties": {
		//	    "CloudWatchLogDestination": {
		//	      "additionalProperties": false,
		//	      "description": "The destination for OpenSearch Ingestion Service logs sent to Amazon CloudWatch.",
		//	      "properties": {
		//	        "LogGroup": {
		//	          "maxLength": 512,
		//	          "minLength": 1,
		//	          "pattern": "\\/aws\\/vendedlogs\\/[\\.\\-_/#A-Za-z0-9]+",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "LogGroup"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "IsLoggingEnabled": {
		//	      "description": "Whether logs should be published.",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"log_publishing_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CloudWatchLogDestination
				"cloudwatch_log_destination": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: LogGroup
						"log_group": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.LengthBetween(1, 512),
								stringvalidator.RegexMatches(regexp.MustCompile("\\/aws\\/vendedlogs\\/[\\.\\-_/#A-Za-z0-9]+"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "The destination for OpenSearch Ingestion Service logs sent to Amazon CloudWatch.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: IsLoggingEnabled
				"is_logging_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether logs should be published.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Key-value pairs to configure log publishing.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MaxUnits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The maximum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
		//	  "maximum": 384,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"max_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The maximum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
			Required:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 384),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: MinUnits
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The minimum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
		//	  "maximum": 384,
		//	  "minimum": 1,
		//	  "type": "integer"
		//	}
		"min_units": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The minimum pipeline capacity, in Ingestion OpenSearch Compute Units (OCUs).",
			Required:    true,
			Validators: []validator.Int64{ /*START VALIDATORS*/
				int64validator.Between(1, 384),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the pipeline.",
		//	  "maxLength": 76,
		//	  "minLength": 46,
		//	  "pattern": "^arn:(aws|aws\\-cn|aws\\-us\\-gov|aws\\-iso|aws\\-iso\\-b):osis:.+:pipeline\\/.+$",
		//	  "type": "string"
		//	}
		"pipeline_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the pipeline.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineConfigurationBody
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Data Prepper pipeline configuration.",
		//	  "maxLength": 24000,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"pipeline_configuration_body": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Data Prepper pipeline configuration.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 24000),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PipelineName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Name of the OpenSearch Ingestion Service pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region.",
		//	  "maxLength": 28,
		//	  "minLength": 3,
		//	  "pattern": "[a-z][a-z0-9\\-]+",
		//	  "type": "string"
		//	}
		"pipeline_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Name of the OpenSearch Ingestion Service pipeline to create. Pipeline names are unique across the pipelines owned by an account within an AWS Region.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(3, 28),
				stringvalidator.RegexMatches(regexp.MustCompile("[a-z][a-z0-9\\-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
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
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcEndpointService
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The VPC endpoint service name for the pipeline.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"vpc_endpoint_service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The VPC endpoint service name for the pipeline.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcEndpoints
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The VPC interface endpoints that have access to the pipeline.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "An OpenSearch Ingestion Service-managed VPC endpoint that will access one or more pipelines.",
		//	    "properties": {
		//	      "VpcEndpointId": {
		//	        "description": "The unique identifier of the endpoint.",
		//	        "type": "string"
		//	      },
		//	      "VpcId": {
		//	        "description": "The ID for your VPC. AWS Privatelink generates this value when you create a VPC.",
		//	        "type": "string"
		//	      },
		//	      "VpcOptions": {
		//	        "additionalProperties": false,
		//	        "description": "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
		//	        "properties": {
		//	          "SecurityGroupIds": {
		//	            "description": "A list of security groups associated with the VPC endpoint.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 20,
		//	              "minLength": 11,
		//	              "pattern": "sg-\\w{8}(\\w{9})?",
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "SubnetIds": {
		//	            "description": "A list of subnet IDs associated with the VPC endpoint.",
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "maxLength": 24,
		//	              "minLength": 15,
		//	              "pattern": "subnet-\\w{8}(\\w{9})?",
		//	              "type": "string"
		//	            },
		//	            "type": "array"
		//	          },
		//	          "VpcAttachmentOptions": {
		//	            "additionalProperties": false,
		//	            "description": "Options for attaching a VPC to the pipeline.",
		//	            "properties": {
		//	              "AttachToVpc": {
		//	                "description": "Whether the pipeline should be attached to the provided VPC",
		//	                "type": "boolean"
		//	              },
		//	              "CidrBlock": {
		//	                "description": "The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).",
		//	                "pattern": "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(3[0-2]|[12]?[0-9])$",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "AttachToVpc",
		//	              "CidrBlock"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "VpcEndpointManagement": {
		//	            "description": "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
		//	            "enum": [
		//	              "CUSTOMER",
		//	              "SERVICE"
		//	            ],
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "SubnetIds"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"vpc_endpoints": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: VpcEndpointId
					"vpc_endpoint_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The unique identifier of the endpoint.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VpcId
					"vpc_id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID for your VPC. AWS Privatelink generates this value when you create a VPC.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: VpcOptions
					"vpc_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: SecurityGroupIds
							"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of security groups associated with the VPC endpoint.",
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: SubnetIds
							"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Description: "A list of subnet IDs associated with the VPC endpoint.",
								Computed:    true,
								PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
									generic.Multiset(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: VpcAttachmentOptions
							"vpc_attachment_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: AttachToVpc
									"attach_to_vpc": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Description: "Whether the pipeline should be attached to the provided VPC",
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: CidrBlock
									"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).",
										Computed:    true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Options for attaching a VPC to the pipeline.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: VpcEndpointManagement
							"vpc_endpoint_management": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The VPC interface endpoints that have access to the pipeline.",
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VpcOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "description": "A list of security groups associated with the VPC endpoint.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 20,
		//	        "minLength": 11,
		//	        "pattern": "sg-\\w{8}(\\w{9})?",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "SubnetIds": {
		//	      "description": "A list of subnet IDs associated with the VPC endpoint.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "maxLength": 24,
		//	        "minLength": 15,
		//	        "pattern": "subnet-\\w{8}(\\w{9})?",
		//	        "type": "string"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "VpcAttachmentOptions": {
		//	      "additionalProperties": false,
		//	      "description": "Options for attaching a VPC to the pipeline.",
		//	      "properties": {
		//	        "AttachToVpc": {
		//	          "description": "Whether the pipeline should be attached to the provided VPC",
		//	          "type": "boolean"
		//	        },
		//	        "CidrBlock": {
		//	          "description": "The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).",
		//	          "pattern": "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(3[0-2]|[12]?[0-9])$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "AttachToVpc",
		//	        "CidrBlock"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "VpcEndpointManagement": {
		//	      "description": "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
		//	      "enum": [
		//	        "CUSTOMER",
		//	        "SERVICE"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "SubnetIds"
		//	  ],
		//	  "type": "object"
		//	}
		"vpc_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of security groups associated with the VPC endpoint.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(11, 20),
							stringvalidator.RegexMatches(regexp.MustCompile("sg-\\w{8}(\\w{9})?"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: SubnetIds
				"subnet_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list of subnet IDs associated with the VPC endpoint.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(15, 24),
							stringvalidator.RegexMatches(regexp.MustCompile("subnet-\\w{8}(\\w{9})?"), ""),
						),
						fwvalidators.NotNullList(),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VpcAttachmentOptions
				"vpc_attachment_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AttachToVpc
						"attach_to_vpc": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Description: "Whether the pipeline should be attached to the provided VPC",
							Optional:    true,
							Computed:    true,
							Validators: []validator.Bool{ /*START VALIDATORS*/
								fwvalidators.NotNullBool(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: CidrBlock
						"cidr_block": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "The CIDR block to be reserved for OpenSearch Ingestion to create elastic network interfaces (ENIs).",
							Optional:    true,
							Computed:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/(3[0-2]|[12]?[0-9])$"), ""),
								fwvalidators.NotNullString(),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Options for attaching a VPC to the pipeline.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: VpcEndpointManagement
				"vpc_endpoint_management": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Defines whether you or Amazon OpenSearch Ingestion service create and manage the VPC endpoint configured for the pipeline.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"CUSTOMER",
							"SERVICE",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Container for the values required to configure VPC access for the pipeline. If you don't specify these values, OpenSearch Ingestion Service creates the pipeline with a public endpoint.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
			// VpcOptions is a write-only property.
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
		Description: "An OpenSearch Ingestion Service Data Prepper pipeline running Data Prepper.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::OSIS::Pipeline").WithTerraformTypeName("awscc_osis_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attach_to_vpc":               "AttachToVpc",
		"buffer_options":              "BufferOptions",
		"cidr_block":                  "CidrBlock",
		"cloudwatch_log_destination":  "CloudWatchLogDestination",
		"encryption_at_rest_options":  "EncryptionAtRestOptions",
		"ingest_endpoint_urls":        "IngestEndpointUrls",
		"is_logging_enabled":          "IsLoggingEnabled",
		"key":                         "Key",
		"kms_key_arn":                 "KmsKeyArn",
		"log_group":                   "LogGroup",
		"log_publishing_options":      "LogPublishingOptions",
		"max_units":                   "MaxUnits",
		"min_units":                   "MinUnits",
		"persistent_buffer_enabled":   "PersistentBufferEnabled",
		"pipeline_arn":                "PipelineArn",
		"pipeline_configuration_body": "PipelineConfigurationBody",
		"pipeline_name":               "PipelineName",
		"security_group_ids":          "SecurityGroupIds",
		"subnet_ids":                  "SubnetIds",
		"tags":                        "Tags",
		"value":                       "Value",
		"vpc_attachment_options":      "VpcAttachmentOptions",
		"vpc_endpoint_id":             "VpcEndpointId",
		"vpc_endpoint_management":     "VpcEndpointManagement",
		"vpc_endpoint_service":        "VpcEndpointService",
		"vpc_endpoints":               "VpcEndpoints",
		"vpc_id":                      "VpcId",
		"vpc_options":                 "VpcOptions",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/VpcOptions",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
