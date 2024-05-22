// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package detective

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_detective_organization_admin", organizationAdminDataSource)
}

// organizationAdminDataSource returns the Terraform awscc_detective_organization_admin data source.
// This Terraform data source corresponds to the CloudFormation AWS::Detective::OrganizationAdmin resource.
func organizationAdminDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The account ID of the account that should be registered as your Organization's delegated administrator for Detective",
		//	  "pattern": "[0-9]{12}",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The account ID of the account that should be registered as your Organization's delegated administrator for Detective",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: GraphArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Detective graph ARN",
		//	  "type": "string"
		//	}
		"graph_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Detective graph ARN",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Detective::OrganizationAdmin",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Detective::OrganizationAdmin").WithTerraformTypeName("awscc_detective_organization_admin")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id": "AccountId",
		"graph_arn":  "GraphArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
