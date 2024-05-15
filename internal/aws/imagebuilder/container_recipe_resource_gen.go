// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package imagebuilder

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_imagebuilder_container_recipe", containerRecipeResource)
}

// containerRecipeResource returns the Terraform awscc_imagebuilder_container_recipe resource.
// This Terraform resource corresponds to the CloudFormation AWS::ImageBuilder::ContainerRecipe resource.
func containerRecipeResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the container recipe.",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the container recipe.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Components
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Components for build and test that are included in the container recipe.",
		//	  "insertionOrder": true,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Configuration details of the component.",
		//	    "properties": {
		//	      "ComponentArn": {
		//	        "description": "The Amazon Resource Name (ARN) of the component.",
		//	        "type": "string"
		//	      },
		//	      "Parameters": {
		//	        "description": "A group of parameter settings that are used to configure the component for a specific recipe.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "description": "Contains a key/value pair that sets the named component parameter.",
		//	          "properties": {
		//	            "Name": {
		//	              "description": "The name of the component parameter to set.",
		//	              "type": "string"
		//	            },
		//	            "Value": {
		//	              "description": "Sets the value for the named component parameter.",
		//	              "insertionOrder": true,
		//	              "items": {
		//	                "type": "string"
		//	              },
		//	              "type": "array"
		//	            }
		//	          },
		//	          "required": [
		//	            "Name",
		//	            "Value"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"components": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: ComponentArn
					"component_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The Amazon Resource Name (ARN) of the component.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Parameters
					"parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Name
								"name": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The name of the component parameter to set.",
									Required:    true,
								}, /*END ATTRIBUTE*/
								// Property: Value
								"value": schema.ListAttribute{ /*START ATTRIBUTE*/
									ElementType: types.StringType,
									Description: "Sets the value for the named component parameter.",
									Required:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "A group of parameter settings that are used to configure the component for a specific recipe.",
						Optional:    true,
						Computed:    true,
						PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
							generic.Multiset(),
							listplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Components for build and test that are included in the container recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
				listplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ContainerType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the type of container, such as Docker.",
		//	  "enum": [
		//	    "DOCKER"
		//	  ],
		//	  "type": "string"
		//	}
		"container_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the type of container, such as Docker.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"DOCKER",
				),
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
		//	  "description": "The description of the container recipe.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the container recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DockerfileTemplateData
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.",
		//	  "type": "string"
		//	}
		"dockerfile_template_data": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Dockerfiles are text documents that are used to build Docker containers, and ensure that they contain all of the elements required by the application running inside. The template data consists of contextual variables where Image Builder places build information or scripts, based on your container image recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// DockerfileTemplateData is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: DockerfileTemplateUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The S3 URI for the Dockerfile that will be used to build your container image.",
		//	  "type": "string"
		//	}
		"dockerfile_template_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The S3 URI for the Dockerfile that will be used to build your container image.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// DockerfileTemplateUri is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ImageOsVersionOverride
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the operating system version for the source image.",
		//	  "type": "string"
		//	}
		"image_os_version_override": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the operating system version for the source image.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// ImageOsVersionOverride is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: InstanceConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A group of options that can be used to configure an instance for building and testing container images.",
		//	  "properties": {
		//	    "BlockDeviceMappings": {
		//	      "description": "Defines the block devices to attach for building an instance from this Image Builder AMI.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Defines block device mappings for the instance used to configure your image. ",
		//	        "properties": {
		//	          "DeviceName": {
		//	            "description": "The device to which these mappings apply.",
		//	            "type": "string"
		//	          },
		//	          "Ebs": {
		//	            "additionalProperties": false,
		//	            "description": "Use to manage Amazon EBS-specific configuration for this mapping.",
		//	            "properties": {
		//	              "DeleteOnTermination": {
		//	                "description": "Use to configure delete on termination of the associated device.",
		//	                "type": "boolean"
		//	              },
		//	              "Encrypted": {
		//	                "description": "Use to configure device encryption.",
		//	                "type": "boolean"
		//	              },
		//	              "Iops": {
		//	                "description": "Use to configure device IOPS.",
		//	                "type": "integer"
		//	              },
		//	              "KmsKeyId": {
		//	                "description": "Use to configure the KMS key to use when encrypting the device.",
		//	                "type": "string"
		//	              },
		//	              "SnapshotId": {
		//	                "description": "The snapshot that defines the device contents.",
		//	                "type": "string"
		//	              },
		//	              "Throughput": {
		//	                "description": "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
		//	                "type": "integer"
		//	              },
		//	              "VolumeSize": {
		//	                "description": "Use to override the device's volume size.",
		//	                "type": "integer"
		//	              },
		//	              "VolumeType": {
		//	                "description": "Use to override the device's volume type.",
		//	                "enum": [
		//	                  "standard",
		//	                  "io1",
		//	                  "io2",
		//	                  "gp2",
		//	                  "gp3",
		//	                  "sc1",
		//	                  "st1"
		//	                ],
		//	                "type": "string"
		//	              }
		//	            },
		//	            "type": "object"
		//	          },
		//	          "NoDevice": {
		//	            "description": "Use to remove a mapping from the parent image.",
		//	            "type": "string"
		//	          },
		//	          "VirtualName": {
		//	            "description": "Use to manage instance ephemeral devices.",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    },
		//	    "Image": {
		//	      "description": "The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"instance_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BlockDeviceMappings
				"block_device_mappings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: DeviceName
							"device_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The device to which these mappings apply.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Ebs
							"ebs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: DeleteOnTermination
									"delete_on_termination": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Description: "Use to configure delete on termination of the associated device.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Encrypted
									"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
										Description: "Use to configure device encryption.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
											boolplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Iops
									"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Use to configure device IOPS.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: KmsKeyId
									"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Use to configure the KMS key to use when encrypting the device.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: SnapshotId
									"snapshot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "The snapshot that defines the device contents.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Throughput
									"throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "For GP3 volumes only - The throughput in MiB/s that the volume supports.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: VolumeSize
									"volume_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
										Description: "Use to override the device's volume size.",
										Optional:    true,
										Computed:    true,
										PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
											int64planmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: VolumeType
									"volume_type": schema.StringAttribute{ /*START ATTRIBUTE*/
										Description: "Use to override the device's volume type.",
										Optional:    true,
										Computed:    true,
										Validators: []validator.String{ /*START VALIDATORS*/
											stringvalidator.OneOf(
												"standard",
												"io1",
												"io2",
												"gp2",
												"gp3",
												"sc1",
												"st1",
											),
										}, /*END VALIDATORS*/
										PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
											stringplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Description: "Use to manage Amazon EBS-specific configuration for this mapping.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
									objectplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: NoDevice
							"no_device": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use to remove a mapping from the parent image.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: VirtualName
							"virtual_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Use to manage instance ephemeral devices.",
								Optional:    true,
								Computed:    true,
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Defines the block devices to attach for building an instance from this Image Builder AMI.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Image
				"image": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The AMI ID to use as the base image for a container build and test instance. If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "A group of options that can be used to configure an instance for building and testing container images.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Identifies which KMS key is used to encrypt the container image.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Identifies which KMS key is used to encrypt the container image.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the container recipe.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the container recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ParentImage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The source image for the container recipe.",
		//	  "type": "string"
		//	}
		"parent_image": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The source image for the container recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: PlatformOverride
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Specifies the operating system platform when you use a custom source image.",
		//	  "enum": [
		//	    "Windows",
		//	    "Linux"
		//	  ],
		//	  "type": "string"
		//	}
		"platform_override": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies the operating system platform when you use a custom source image.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"Windows",
					"Linux",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// PlatformOverride is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Tags that are attached to the container recipe.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "Tags that are attached to the container recipe.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
				mapplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TargetRepository
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The destination repository for the container image.",
		//	  "properties": {
		//	    "RepositoryName": {
		//	      "description": "The name of the container repository where the output container image is stored. This name is prefixed by the repository location.",
		//	      "type": "string"
		//	    },
		//	    "Service": {
		//	      "description": "Specifies the service in which this image was registered.",
		//	      "enum": [
		//	        "ECR"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"target_repository": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: RepositoryName
				"repository_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the container repository where the output container image is stored. This name is prefixed by the repository location.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Service
				"service": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the service in which this image was registered.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.OneOf(
							"ECR",
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The destination repository for the container image.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
				objectplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Version
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The semantic version of the container recipe (\u003cmajor\u003e.\u003cminor\u003e.\u003cpatch\u003e).",
		//	  "type": "string"
		//	}
		"version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The semantic version of the container recipe (<major>.<minor>.<patch>).",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: WorkingDirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The working directory to be used during build and test workflows.",
		//	  "type": "string"
		//	}
		"working_directory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The working directory to be used during build and test workflows.",
			Optional:    true,
			Computed:    true,
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
		Description: "Resource schema for AWS::ImageBuilder::ContainerRecipe",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::ImageBuilder::ContainerRecipe").WithTerraformTypeName("awscc_imagebuilder_container_recipe")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"block_device_mappings":     "BlockDeviceMappings",
		"component_arn":             "ComponentArn",
		"components":                "Components",
		"container_type":            "ContainerType",
		"delete_on_termination":     "DeleteOnTermination",
		"description":               "Description",
		"device_name":               "DeviceName",
		"dockerfile_template_data":  "DockerfileTemplateData",
		"dockerfile_template_uri":   "DockerfileTemplateUri",
		"ebs":                       "Ebs",
		"encrypted":                 "Encrypted",
		"image":                     "Image",
		"image_os_version_override": "ImageOsVersionOverride",
		"instance_configuration":    "InstanceConfiguration",
		"iops":                      "Iops",
		"kms_key_id":                "KmsKeyId",
		"name":                      "Name",
		"no_device":                 "NoDevice",
		"parameters":                "Parameters",
		"parent_image":              "ParentImage",
		"platform_override":         "PlatformOverride",
		"repository_name":           "RepositoryName",
		"service":                   "Service",
		"snapshot_id":               "SnapshotId",
		"tags":                      "Tags",
		"target_repository":         "TargetRepository",
		"throughput":                "Throughput",
		"value":                     "Value",
		"version":                   "Version",
		"virtual_name":              "VirtualName",
		"volume_size":               "VolumeSize",
		"volume_type":               "VolumeType",
		"working_directory":         "WorkingDirectory",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/DockerfileTemplateData",
		"/properties/DockerfileTemplateUri",
		"/properties/ImageOsVersionOverride",
		"/properties/PlatformOverride",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
