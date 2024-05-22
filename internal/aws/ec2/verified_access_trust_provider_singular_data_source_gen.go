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
	registry.AddDataSourceFactory("awscc_ec2_verified_access_trust_provider", verifiedAccessTrustProviderDataSource)
}

// verifiedAccessTrustProviderDataSource returns the Terraform awscc_ec2_verified_access_trust_provider data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::VerifiedAccessTrustProvider resource.
func verifiedAccessTrustProviderDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreationTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The creation time.",
		//	  "type": "string"
		//	}
		"creation_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The creation time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A description for the Amazon Web Services Verified Access trust provider.",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "A description for the Amazon Web Services Verified Access trust provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeviceOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The options for device identity based trust providers.",
		//	  "properties": {
		//	    "PublicSigningKeyUrl": {
		//	      "description": "URL Verified Access will use to verify authenticity of the device tokens.",
		//	      "type": "string"
		//	    },
		//	    "TenantId": {
		//	      "description": "The ID of the tenant application with the device-identity provider.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"device_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: PublicSigningKeyUrl
				"public_signing_key_url": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "URL Verified Access will use to verify authenticity of the device tokens.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TenantId
				"tenant_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The ID of the tenant application with the device-identity provider.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The options for device identity based trust providers.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: DeviceTrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of device-based trust provider. Possible values: jamf|crowdstrike",
		//	  "type": "string"
		//	}
		"device_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of device-based trust provider. Possible values: jamf|crowdstrike",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LastUpdatedTime
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The last updated time.",
		//	  "type": "string"
		//	}
		"last_updated_time": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The last updated time.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OidcOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The OpenID Connect details for an oidc -type, user-identity based trust provider.",
		//	  "properties": {
		//	    "AuthorizationEndpoint": {
		//	      "description": "The OIDC authorization endpoint.",
		//	      "type": "string"
		//	    },
		//	    "ClientId": {
		//	      "description": "The client identifier.",
		//	      "type": "string"
		//	    },
		//	    "ClientSecret": {
		//	      "description": "The client secret.",
		//	      "type": "string"
		//	    },
		//	    "Issuer": {
		//	      "description": "The OIDC issuer.",
		//	      "type": "string"
		//	    },
		//	    "Scope": {
		//	      "description": "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
		//	      "type": "string"
		//	    },
		//	    "TokenEndpoint": {
		//	      "description": "The OIDC token endpoint.",
		//	      "type": "string"
		//	    },
		//	    "UserInfoEndpoint": {
		//	      "description": "The OIDC user info endpoint.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"oidc_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AuthorizationEndpoint
				"authorization_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC authorization endpoint.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ClientId
				"client_id": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The client identifier.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: ClientSecret
				"client_secret": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The client secret.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Issuer
				"issuer": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC issuer.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: Scope
				"scope": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "OpenID Connect (OIDC) scopes are used by an application during authentication to authorize access to details of a user. Each scope returns a specific set of user attributes.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: TokenEndpoint
				"token_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC token endpoint.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: UserInfoEndpoint
				"user_info_endpoint": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The OIDC user info endpoint.",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The OpenID Connect details for an oidc -type, user-identity based trust provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyReferenceName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier to be used when working with policy rules.",
		//	  "type": "string"
		//	}
		"policy_reference_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier to be used when working with policy rules.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SseSpecification
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "The configuration options for customer provided KMS encryption.",
		//	  "properties": {
		//	    "CustomerManagedKeyEnabled": {
		//	      "description": "Whether to encrypt the policy with the provided key or disable encryption",
		//	      "type": "boolean"
		//	    },
		//	    "KmsKeyArn": {
		//	      "description": "KMS Key Arn used to encrypt the group policy",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sse_specification": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomerManagedKeyEnabled
				"customer_managed_key_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "Whether to encrypt the policy with the provided key or disable encryption",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: KmsKeyArn
				"kms_key_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "KMS Key Arn used to encrypt the group policy",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "The configuration options for customer provided KMS encryption.",
			Computed:    true,
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
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: TrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Type of trust provider. Possible values: user|device",
		//	  "type": "string"
		//	}
		"trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Type of trust provider. Possible values: user|device",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: UserTrustProviderType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of device-based trust provider. Possible values: oidc|iam-identity-center",
		//	  "type": "string"
		//	}
		"user_trust_provider_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of device-based trust provider. Possible values: oidc|iam-identity-center",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VerifiedAccessTrustProviderId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Amazon Web Services Verified Access trust provider.",
		//	  "type": "string"
		//	}
		"verified_access_trust_provider_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Amazon Web Services Verified Access trust provider.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::VerifiedAccessTrustProvider",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::VerifiedAccessTrustProvider").WithTerraformTypeName("awscc_ec2_verified_access_trust_provider")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"authorization_endpoint":            "AuthorizationEndpoint",
		"client_id":                         "ClientId",
		"client_secret":                     "ClientSecret",
		"creation_time":                     "CreationTime",
		"customer_managed_key_enabled":      "CustomerManagedKeyEnabled",
		"description":                       "Description",
		"device_options":                    "DeviceOptions",
		"device_trust_provider_type":        "DeviceTrustProviderType",
		"issuer":                            "Issuer",
		"key":                               "Key",
		"kms_key_arn":                       "KmsKeyArn",
		"last_updated_time":                 "LastUpdatedTime",
		"oidc_options":                      "OidcOptions",
		"policy_reference_name":             "PolicyReferenceName",
		"public_signing_key_url":            "PublicSigningKeyUrl",
		"scope":                             "Scope",
		"sse_specification":                 "SseSpecification",
		"tags":                              "Tags",
		"tenant_id":                         "TenantId",
		"token_endpoint":                    "TokenEndpoint",
		"trust_provider_type":               "TrustProviderType",
		"user_info_endpoint":                "UserInfoEndpoint",
		"user_trust_provider_type":          "UserTrustProviderType",
		"value":                             "Value",
		"verified_access_trust_provider_id": "VerifiedAccessTrustProviderId",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
