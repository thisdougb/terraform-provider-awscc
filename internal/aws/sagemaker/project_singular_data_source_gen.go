// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_sagemaker_project", projectDataSource)
}

// projectDataSource returns the Terraform awscc_sagemaker_project data source.
// This Terraform data source corresponds to the CloudFormation AWS::SageMaker::Project resource.
func projectDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time at which the project was created.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time at which the project was created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Project.",
		//	  "maxLength": 2048,
		//	  "minLength": 1,
		//	  "pattern": "arn:aws[a-z\\-]*:sagemaker:[a-z0-9\\-]*:[0-9]{12}:project.*",
		//	  "type": "string"
		//	}
		"project_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectDescription
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the project.",
		//	  "maxLength": 1024,
		//	  "pattern": ".*",
		//	  "type": "string"
		//	}
		"project_description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Project Id.",
		//	  "maxLength": 20,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*",
		//	  "type": "string"
		//	}
		"project_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Project Id.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the project.",
		//	  "maxLength": 32,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	  "type": "string"
		//	}
		"project_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ProjectStatus
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The status of a project.",
		//	  "enum": [
		//	    "Pending",
		//	    "CreateInProgress",
		//	    "CreateCompleted",
		//	    "CreateFailed",
		//	    "DeleteInProgress",
		//	    "DeleteFailed",
		//	    "DeleteCompleted"
		//	  ],
		//	  "type": "string"
		//	}
		"project_status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The status of a project.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceCatalogProvisionedProductDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Provisioned ServiceCatalog  Details",
		//	  "properties": {
		//	    "ProvisionedProductId": {
		//	      "description": "The identifier of the provisioning artifact (also known as a version).",
		//	      "maxLength": 100,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "ProvisionedProductStatusMessage": {
		//	      "description": "Provisioned Product Status Message",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"service_catalog_provisioned_product_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ProvisionedProductId
				"provisioned_product_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the provisioning artifact (also known as a version).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ProvisionedProductStatusMessage
				"provisioned_product_status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Provisioned Product Status Message",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Provisioned ServiceCatalog  Details",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ServiceCatalogProvisioningDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Input ServiceCatalog Provisioning Details",
		//	  "properties": {
		//	    "PathId": {
		//	      "description": "The path identifier of the product.",
		//	      "maxLength": 100,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "ProductId": {
		//	      "description": "Service Catalog product identifier.",
		//	      "maxLength": 100,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "ProvisioningArtifactId": {
		//	      "description": "The identifier of the provisioning artifact (also known as a version).",
		//	      "maxLength": 100,
		//	      "pattern": "^[a-zA-Z0-9](-*[a-zA-Z0-9])*$",
		//	      "type": "string"
		//	    },
		//	    "ProvisioningParameters": {
		//	      "description": "Parameters specified by the administrator that are required for provisioning the product.",
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Information about a parameter used to provision a product.",
		//	        "properties": {
		//	          "Key": {
		//	            "description": "The parameter key.",
		//	            "maxLength": 1000,
		//	            "minLength": 1,
		//	            "pattern": ".*",
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "description": "The parameter value.",
		//	            "maxLength": 4096,
		//	            "pattern": ".*",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "required": [
		//	          "Key",
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "required": [
		//	    "ProductId"
		//	  ],
		//	  "type": "object"
		//	}
		"service_catalog_provisioning_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PathId
				"path_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The path identifier of the product.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ProductId
				"product_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Service Catalog product identifier.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ProvisioningArtifactId
				"provisioning_artifact_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The identifier of the provisioning artifact (also known as a version).",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ProvisioningParameters
				"provisioning_parameters": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Key
							"key": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The parameter key.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "The parameter value.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Parameters specified by the administrator that are required for provisioning the product.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Input ServiceCatalog Provisioning Details",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
		//	        "maxLength": 256,
		//	        "pattern": "^([\\p{L}\\p{Z}\\p{N}_.:/=+\\-@]*)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 40,
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 127 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 1 to 255 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
		Description: "Data Source schema for AWS::SageMaker::Project",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::Project").WithTerraformTypeName("awscc_sagemaker_project")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"creation_time":                      "CreationTime",
		"key":                                "Key",
		"path_id":                            "PathId",
		"product_id":                         "ProductId",
		"project_arn":                        "ProjectArn",
		"project_description":                "ProjectDescription",
		"project_id":                         "ProjectId",
		"project_name":                       "ProjectName",
		"project_status":                     "ProjectStatus",
		"provisioned_product_id":             "ProvisionedProductId",
		"provisioned_product_status_message": "ProvisionedProductStatusMessage",
		"provisioning_artifact_id":           "ProvisioningArtifactId",
		"provisioning_parameters":            "ProvisioningParameters",
		"service_catalog_provisioned_product_details": "ServiceCatalogProvisionedProductDetails",
		"service_catalog_provisioning_details":        "ServiceCatalogProvisioningDetails",
		"tags":                                        "Tags",
		"value":                                       "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
