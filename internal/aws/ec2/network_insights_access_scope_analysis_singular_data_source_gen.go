// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ec2_network_insights_access_scope_analysis", networkInsightsAccessScopeAnalysisDataSource)
}

// networkInsightsAccessScopeAnalysisDataSource returns the Terraform awscc_ec2_network_insights_access_scope_analysis data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::NetworkInsightsAccessScopeAnalysis resource.
func networkInsightsAccessScopeAnalysisDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AnalyzedEniCount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"analyzed_eni_count": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: EndDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"end_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: FindingsFound
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "true",
		//	    "false",
		//	    "unknown"
		//	  ],
		//	  "type": "string"
		//	}
		"findings_found": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsAccessScopeAnalysisArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_access_scope_analysis_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsAccessScopeAnalysisId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_access_scope_analysis_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkInsightsAccessScopeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"network_insights_access_scope_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StartDate
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"start_date": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Status
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "running",
		//	    "failed",
		//	    "succeeded"
		//	  ],
		//	  "type": "string"
		//	}
		"status": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: StatusMessage
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"status_message": schema.StringAttribute{ /*START ATTRIBUTE*/
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
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
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
		Description: "Data Source schema for AWS::EC2::NetworkInsightsAccessScopeAnalysis",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::NetworkInsightsAccessScopeAnalysis").WithTerraformTypeName("awscc_ec2_network_insights_access_scope_analysis")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"analyzed_eni_count": "AnalyzedEniCount",
		"end_date":           "EndDate",
		"findings_found":     "FindingsFound",
		"key":                "Key",
		"network_insights_access_scope_analysis_arn": "NetworkInsightsAccessScopeAnalysisArn",
		"network_insights_access_scope_analysis_id":  "NetworkInsightsAccessScopeAnalysisId",
		"network_insights_access_scope_id":           "NetworkInsightsAccessScopeId",
		"start_date":                                 "StartDate",
		"status":                                     "Status",
		"status_message":                             "StatusMessage",
		"tags":                                       "Tags",
		"value":                                      "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
