// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iam

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_iam_group_policy", groupPolicyDataSource)
}

// groupPolicyDataSource returns the Terraform awscc_iam_group_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::IAM::GroupPolicy resource.
func groupPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: GroupName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the group to associate the policy with.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.",
		//	  "type": "string"
		//	}
		"group_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the group to associate the policy with.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The policy document.\n You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.\n The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:\n  +  Any printable ASCII character ranging from the space character (``\\u0020``) through the end of the ASCII character range\n  +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\\u00FF``)\n  +  The special characters tab (``\\u0009``), line feed (``\\u000A``), and carriage return (``\\u000D``)",
		//	  "type": "object"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			CustomType:  jsontypes.NormalizedType{},
			Description: "The policy document.\n You must provide policies in JSON format in IAM. However, for CFN templates formatted in YAML, you can provide the policy in JSON or YAML format. CFN always converts a YAML policy to JSON format before submitting it to IAM.\n The [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex) used to validate this parameter is a string of characters consisting of the following:\n  +  Any printable ASCII character ranging from the space character (``\\u0020``) through the end of the ASCII character range\n  +  The printable characters in the Basic Latin and Latin-1 Supplement character set (through ``\\u00FF``)\n  +  The special characters tab (``\\u0009``), line feed (``\\u000A``), and carriage return (``\\u000D``)",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the policy document.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the policy document.\n This parameter allows (through its [regex pattern](https://docs.aws.amazon.com/http://wikipedia.org/wiki/regex)) a string of characters consisting of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: _+=,.@-",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::IAM::GroupPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::IAM::GroupPolicy").WithTerraformTypeName("awscc_iam_group_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"group_name":      "GroupName",
		"policy_document": "PolicyDocument",
		"policy_name":     "PolicyName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
