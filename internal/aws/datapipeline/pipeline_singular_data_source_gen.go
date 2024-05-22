// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package datapipeline

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_datapipeline_pipeline", pipelineDataSource)
}

// pipelineDataSource returns the Terraform awscc_datapipeline_pipeline data source.
// This Terraform data source corresponds to the CloudFormation AWS::DataPipeline::Pipeline resource.
func pipelineDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Activate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
		//	  "type": "boolean"
		//	}
		"activate": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether to validate and start the pipeline or stop an active pipeline. By default, the value is set to true.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description of the pipeline.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description of the pipeline.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the pipeline.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the pipeline.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ParameterObjects
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parameter objects used with the pipeline.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Attributes": {
		//	        "description": "The attributes of the parameter object.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "description": "The field identifier.",
		//	              "type": "string"
		//	            },
		//	            "StringValue": {
		//	              "description": "The field value, expressed as a String.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key",
		//	            "StringValue"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "Id": {
		//	        "description": "The ID of the parameter object.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Attributes",
		//	      "Id"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"parameter_objects": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Attributes
					"attributes": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The field identifier.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: StringValue
								"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The field value, expressed as a String.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "The attributes of the parameter object.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the parameter object.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The parameter objects used with the pipeline.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ParameterValues
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The parameter values used with the pipeline.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Id": {
		//	        "description": "The ID of the parameter value.",
		//	        "type": "string"
		//	      },
		//	      "StringValue": {
		//	        "description": "The field value, expressed as a String.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Id",
		//	      "StringValue"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"parameter_values": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the parameter value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: StringValue
					"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The field value, expressed as a String.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The parameter values used with the pipeline.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PipelineId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"pipeline_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PipelineObjects
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Fields": {
		//	        "description": "Key-value pairs that define the properties of the object.",
		//	        "insertionOrder": false,
		//	        "items": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Key": {
		//	              "description": "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
		//	              "type": "string"
		//	            },
		//	            "RefValue": {
		//	              "description": "A field value that you specify as an identifier of another object in the same pipeline definition.",
		//	              "type": "string"
		//	            },
		//	            "StringValue": {
		//	              "description": "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Key"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "type": "array",
		//	        "uniqueItems": false
		//	      },
		//	      "Id": {
		//	        "description": "The ID of the object.",
		//	        "type": "string"
		//	      },
		//	      "Name": {
		//	        "description": "The name of the object.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Fields",
		//	      "Id",
		//	      "Name"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"pipeline_objects": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Fields
					"fields": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Key
								"key": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "Specifies the name of a field for a particular object. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: RefValue
								"ref_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field value that you specify as an identifier of another object in the same pipeline definition.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
								// Property: StringValue
								"string_value": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "A field value that you specify as a string. To view valid values for a particular field, see Pipeline Object Reference in the AWS Data Pipeline Developer Guide.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Description: "Key-value pairs that define the properties of the object.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Id
					"id": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The ID of the object.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Name
					"name": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The name of the object.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The objects that define the pipeline. These objects overwrite the existing pipeline definition. Not all objects, fields, and values can be updated. For information about restrictions, see Editing Your Pipeline in the AWS Data Pipeline Developer Guide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PipelineTags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of a tag.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value to associate with the key name.",
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
		//	  "uniqueItems": false
		//	}
		"pipeline_tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of a tag.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value to associate with the key name.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "A list of arbitrary tags (key-value pairs) to associate with the pipeline, which you can use to control permissions. For more information, see Controlling Access to Pipelines and Resources in the AWS Data Pipeline Developer Guide.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::DataPipeline::Pipeline",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataPipeline::Pipeline").WithTerraformTypeName("awscc_datapipeline_pipeline")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"activate":          "Activate",
		"attributes":        "Attributes",
		"description":       "Description",
		"fields":            "Fields",
		"id":                "Id",
		"key":               "Key",
		"name":              "Name",
		"parameter_objects": "ParameterObjects",
		"parameter_values":  "ParameterValues",
		"pipeline_id":       "PipelineId",
		"pipeline_objects":  "PipelineObjects",
		"pipeline_tags":     "PipelineTags",
		"ref_value":         "RefValue",
		"string_value":      "StringValue",
		"value":             "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
