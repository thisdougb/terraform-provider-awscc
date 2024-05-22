// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package lambda

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_lambda_permission", permissionDataSource)
}

// permissionDataSource returns the Terraform awscc_lambda_permission data source.
// This Terraform data source corresponds to the CloudFormation AWS::Lambda::Permission resource.
func permissionDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Action
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^(lambda:[*]|lambda:[a-zA-Z]+|[*])$",
		//	  "type": "string"
		//	}
		"action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The action that the principal can use on the function. For example, ``lambda:InvokeFunction`` or ``lambda:GetFunction``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: EventSourceToken
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For Alexa Smart Home functions, a token that the invoker must supply.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^[a-zA-Z0-9._\\-]+$",
		//	  "type": "string"
		//	}
		"event_source_token": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For Alexa Smart Home functions, a token that the invoker must supply.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name or ARN of the Lambda function, version, or alias.\n  **Name formats**\n +   *Function name* ? ``my-function`` (name-only), ``my-function:v1`` (with alias).\n  +   *Function ARN* ? ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.\n  +   *Partial ARN* ? ``123456789012:function:my-function``.\n  \n You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
		//	  "maxLength": 140,
		//	  "minLength": 1,
		//	  "pattern": "^(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\\d{1}:)?(\\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\\$LATEST|[a-zA-Z0-9-_]+))?$",
		//	  "type": "string"
		//	}
		"function_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name or ARN of the Lambda function, version, or alias.\n  **Name formats**\n +   *Function name* ? ``my-function`` (name-only), ``my-function:v1`` (with alias).\n  +   *Function ARN* ? ``arn:aws:lambda:us-west-2:123456789012:function:my-function``.\n  +   *Partial ARN* ? ``123456789012:function:my-function``.\n  \n You can append a version number or alias to any of the formats. The length constraint applies only to the full ARN. If you specify only the function name, it is limited to 64 characters in length.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: FunctionUrlAuthType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).",
		//	  "enum": [
		//	    "AWS_IAM",
		//	    "NONE"
		//	  ],
		//	  "type": "string"
		//	}
		"function_url_auth_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of authentication that your function URL uses. Set to ``AWS_IAM`` if you want to restrict access to authenticated users only. Set to ``NONE`` if you want to bypass IAM authentication to create a public endpoint. For more information, see [Security and auth model for Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^.*$",
		//	  "type": "string"
		//	}
		"permission_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Principal
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS-service or AWS-account that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.",
		//	  "maxLength": 256,
		//	  "minLength": 1,
		//	  "pattern": "^.*$",
		//	  "type": "string"
		//	}
		"principal": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS-service or AWS-account that invokes the function. If you specify a service, use ``SourceArn`` or ``SourceAccount`` to limit who can invoke the function through that service.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PrincipalOrgID
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.",
		//	  "maxLength": 34,
		//	  "minLength": 12,
		//	  "pattern": "^o-[a-z0-9]{10,32}$",
		//	  "type": "string"
		//	}
		"principal_org_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier for your organization in AOlong. Use this to grant permissions to all the AWS-accounts under this organization.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceAccount
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.",
		//	  "maxLength": 12,
		//	  "minLength": 12,
		//	  "pattern": "^\\d{12}$",
		//	  "type": "string"
		//	}
		"source_account": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For AWS-service, the ID of the AWS-account that owns the resource. Use this together with ``SourceArn`` to ensure that the specified account owns the resource. It is possible for an Amazon S3 bucket to be deleted by its owner and recreated by another account.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.\n Note that Lambda configures the comparison using the ``StringLike`` operator.",
		//	  "maxLength": 1024,
		//	  "minLength": 12,
		//	  "pattern": "^arn:(aws[a-zA-Z0-9-]*):([a-zA-Z0-9\\-])+:([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\\d{1})?:(\\d{12})?:(.*)$",
		//	  "type": "string"
		//	}
		"source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "For AWS-services, the ARN of the AWS resource that invokes the function. For example, an Amazon S3 bucket or Amazon SNS topic.\n Note that Lambda configures the comparison using the ``StringLike`` operator.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Lambda::Permission",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Lambda::Permission").WithTerraformTypeName("awscc_lambda_permission")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                 "Action",
		"event_source_token":     "EventSourceToken",
		"function_name":          "FunctionName",
		"function_url_auth_type": "FunctionUrlAuthType",
		"permission_id":          "Id",
		"principal":              "Principal",
		"principal_org_id":       "PrincipalOrgID",
		"source_account":         "SourceAccount",
		"source_arn":             "SourceArn",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
