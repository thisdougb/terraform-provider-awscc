// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package billingconductor

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_billingconductor_custom_line_item", customLineItemDataSource)
}

// customLineItemDataSource returns the Terraform awscc_billingconductor_custom_line_item data source.
// This Terraform data source corresponds to the CloudFormation AWS::BillingConductor::CustomLineItem resource.
func customLineItemDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The account which this custom line item will be charged to",
		//	  "pattern": "[0-9]{12}",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The account which this custom line item will be charged to",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "ARN",
		//	  "pattern": "(arn:aws(-cn)?:billingconductor::[0-9]{12}:customlineitem/)?[a-zA-Z0-9]{10}",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AssociationSize
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Number of source values associated to this custom line item",
		//	  "type": "integer"
		//	}
		"association_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Number of source values associated to this custom line item",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BillingGroupArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Billing Group ARN",
		//	  "pattern": "arn:aws(-cn)?:billingconductor::[0-9]{12}:billinggroup/?[0-9]{12}",
		//	  "type": "string"
		//	}
		"billing_group_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Billing Group ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: BillingPeriodRange
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "ExclusiveEndBillingPeriod": {
		//	      "pattern": "\\d{4}-(0?[1-9]|1[012])",
		//	      "type": "string"
		//	    },
		//	    "InclusiveStartBillingPeriod": {
		//	      "pattern": "\\d{4}-(0?[1-9]|1[012])",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"billing_period_range": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ExclusiveEndBillingPeriod
				"exclusive_end_billing_period": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: InclusiveStartBillingPeriod
				"inclusive_start_billing_period": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Creation timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"creation_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Creation timestamp in UNIX epoch time format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: CurrencyCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "USD",
		//	    "CNY"
		//	  ],
		//	  "type": "string"
		//	}
		"currency_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: CustomLineItemChargeDetails
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Flat": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ChargeValue": {
		//	          "maximum": 1000000,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "ChargeValue"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "LineItemFilters": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Attribute": {
		//	            "enum": [
		//	              "LINE_ITEM_TYPE"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "MatchOption": {
		//	            "enum": [
		//	              "NOT_EQUAL"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Values": {
		//	            "insertionOrder": false,
		//	            "items": {
		//	              "enum": [
		//	                "SAVINGS_PLAN_NEGATION"
		//	              ],
		//	              "type": "string"
		//	            },
		//	            "type": "array",
		//	            "uniqueItems": true
		//	          }
		//	        },
		//	        "required": [
		//	          "Attribute",
		//	          "MatchOption",
		//	          "Values"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    },
		//	    "Percentage": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ChildAssociatedResources": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "pattern": "(arn:aws(-cn)?:billingconductor::[0-9]{12}:(customlineitem|billinggroup)/)?[a-zA-Z0-9]{10,12}",
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": true
		//	        },
		//	        "PercentageValue": {
		//	          "maximum": 10000,
		//	          "minimum": 0,
		//	          "type": "number"
		//	        }
		//	      },
		//	      "required": [
		//	        "PercentageValue"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "Type": {
		//	      "enum": [
		//	        "FEE",
		//	        "CREDIT"
		//	      ],
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "Type"
		//	  ],
		//	  "type": "object"
		//	}
		"custom_line_item_charge_details": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Flat
				"flat": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ChargeValue
						"charge_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: LineItemFilters
				"line_item_filters": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Attribute
							"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: MatchOption
							"match_option": schema.StringAttribute{ /*START ATTRIBUTE*/
								Computed: true,
							}, /*END ATTRIBUTE*/
							// Property: Values
							"values": schema.SetAttribute{ /*START ATTRIBUTE*/
								ElementType: types.StringType,
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Percentage
				"percentage": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ChildAssociatedResources
						"child_associated_resources": schema.SetAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: PercentageValue
						"percentage_value": schema.Float64Attribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: Type
				"type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: LastModifiedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Latest modified timestamp in UNIX epoch time format",
		//	  "type": "integer"
		//	}
		"last_modified_time": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "Latest modified timestamp in UNIX epoch time format",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[a-zA-Z0-9_\\+=\\.\\-@]+",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: ProductCode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 29,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"product_code": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
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
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::BillingConductor::CustomLineItem",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::BillingConductor::CustomLineItem").WithTerraformTypeName("awscc_billingconductor_custom_line_item")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":                      "AccountId",
		"arn":                             "Arn",
		"association_size":                "AssociationSize",
		"attribute":                       "Attribute",
		"billing_group_arn":               "BillingGroupArn",
		"billing_period_range":            "BillingPeriodRange",
		"charge_value":                    "ChargeValue",
		"child_associated_resources":      "ChildAssociatedResources",
		"creation_time":                   "CreationTime",
		"currency_code":                   "CurrencyCode",
		"custom_line_item_charge_details": "CustomLineItemChargeDetails",
		"description":                     "Description",
		"exclusive_end_billing_period":    "ExclusiveEndBillingPeriod",
		"flat":                            "Flat",
		"inclusive_start_billing_period":  "InclusiveStartBillingPeriod",
		"key":                             "Key",
		"last_modified_time":              "LastModifiedTime",
		"line_item_filters":               "LineItemFilters",
		"match_option":                    "MatchOption",
		"name":                            "Name",
		"percentage":                      "Percentage",
		"percentage_value":                "PercentageValue",
		"product_code":                    "ProductCode",
		"tags":                            "Tags",
		"type":                            "Type",
		"value":                           "Value",
		"values":                          "Values",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
