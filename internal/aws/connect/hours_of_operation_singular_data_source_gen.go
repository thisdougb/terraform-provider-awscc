// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package connect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_connect_hours_of_operation", hoursOfOperationDataSource)
}

// hoursOfOperationDataSource returns the Terraform awscc_connect_hours_of_operation data source.
// This Terraform data source corresponds to the CloudFormation AWS::Connect::HoursOfOperation resource.
func hoursOfOperationDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Config
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Configuration information for the hours of operation: day, start time, and end time.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Contains information about the hours of operation.",
		//	    "properties": {
		//	      "Day": {
		//	        "description": "The day that the hours of operation applies to.",
		//	        "enum": [
		//	          "SUNDAY",
		//	          "MONDAY",
		//	          "TUESDAY",
		//	          "WEDNESDAY",
		//	          "THURSDAY",
		//	          "FRIDAY",
		//	          "SATURDAY"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "EndTime": {
		//	        "additionalProperties": false,
		//	        "description": "The end time that your contact center closes.",
		//	        "properties": {
		//	          "Hours": {
		//	            "description": "The hours.",
		//	            "maximum": 23,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          },
		//	          "Minutes": {
		//	            "description": "The minutes.",
		//	            "maximum": 59,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Hours",
		//	          "Minutes"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "StartTime": {
		//	        "additionalProperties": false,
		//	        "description": "The start time that your contact center opens.",
		//	        "properties": {
		//	          "Hours": {
		//	            "description": "The hours.",
		//	            "maximum": 23,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          },
		//	          "Minutes": {
		//	            "description": "The minutes.",
		//	            "maximum": 59,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Hours",
		//	          "Minutes"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "required": [
		//	      "Day",
		//	      "StartTime",
		//	      "EndTime"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 100,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"config": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Day
					"day": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The day that the hours of operation applies to.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: EndTime
					"end_time": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Hours
							"hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The hours.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Minutes
							"minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The minutes.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The end time that your contact center closes.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: StartTime
					"start_time": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Hours
							"hours": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The hours.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: Minutes
							"minutes": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Description: "The minutes.",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Description: "The start time that your contact center opens.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "Configuration information for the hours of operation: day, start time, and end time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the hours of operation.",
		//	  "maxLength": 250,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the hours of operation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: HoursOfOperationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) for the hours of operation.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/operating-hours/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"hours_of_operation_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) for the hours of operation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the Amazon Connect instance.",
		//	  "pattern": "^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$",
		//	  "type": "string"
		//	}
		"instance_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the Amazon Connect instance.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the hours of operation.",
		//	  "maxLength": 127,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the hours of operation.",
			Computed:    true,
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
		//	        "description": "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
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
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is maximum of 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -. ",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "One or more tags.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TimeZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time zone of the hours of operation.",
		//	  "type": "string"
		//	}
		"time_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time zone of the hours of operation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Connect::HoursOfOperation",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Connect::HoursOfOperation").WithTerraformTypeName("awscc_connect_hours_of_operation")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"config":                 "Config",
		"day":                    "Day",
		"description":            "Description",
		"end_time":               "EndTime",
		"hours":                  "Hours",
		"hours_of_operation_arn": "HoursOfOperationArn",
		"instance_arn":           "InstanceArn",
		"key":                    "Key",
		"minutes":                "Minutes",
		"name":                   "Name",
		"start_time":             "StartTime",
		"tags":                   "Tags",
		"time_zone":              "TimeZone",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
