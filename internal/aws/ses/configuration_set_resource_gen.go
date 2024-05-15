// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ses

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ses_configuration_set", configurationSetResource)
}

// configurationSetResource returns the Terraform awscc_ses_configuration_set resource.
// This Terraform resource corresponds to the CloudFormation AWS::SES::ConfigurationSet resource.
func configurationSetResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DeliveryOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.",
		//	  "properties": {
		//	    "SendingPoolName": {
		//	      "description": "The name of the dedicated IP pool to associate with the configuration set.",
		//	      "type": "string"
		//	    },
		//	    "TlsPolicy": {
		//	      "description": "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.",
		//	      "pattern": "REQUIRE|OPTIONAL",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"delivery_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SendingPoolName
				"sending_pool_name": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The name of the dedicated IP pool to associate with the configuration set.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: TlsPolicy
				"tls_policy": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies whether messages that use the configuration set are required to use Transport Layer Security (TLS). If the value is Require , messages are only delivered if a TLS connection can be established. If the value is Optional , messages can be delivered in plain text if a TLS connection can't be established.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.RegexMatches(regexp.MustCompile("REQUIRE|OPTIONAL"), ""),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that defines the dedicated IP pool that is used to send emails that you send using the configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the configuration set.",
		//	  "pattern": "^[a-zA-Z0-9_-]{1,64}$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the configuration set.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_-]{1,64}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ReputationOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.",
		//	  "properties": {
		//	    "ReputationMetricsEnabled": {
		//	      "description": "If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.",
		//	      "pattern": "true|false",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"reputation_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: ReputationMetricsEnabled
				"reputation_metrics_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Description: "If true , tracking of reputation metrics is enabled for the configuration set. If false , tracking of reputation metrics is disabled for the configuration set.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that defines whether or not Amazon SES collects reputation metrics for the emails that you send that use the configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SendingOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that defines whether or not Amazon SES can send email that you send using the configuration set.",
		//	  "properties": {
		//	    "SendingEnabled": {
		//	      "pattern": "true|false",
		//	      "type": "boolean"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"sending_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SendingEnabled
				"sending_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
						boolplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that defines whether or not Amazon SES can send email that you send using the configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SuppressionOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that contains information about the suppression list preferences for your account.",
		//	  "properties": {
		//	    "SuppressedReasons": {
		//	      "description": "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "description": "The reason that the address was added to the suppression list for your account",
		//	        "pattern": "BOUNCE|COMPLAINT",
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"suppression_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SuppressedReasons
				"suppressed_reasons": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Description: "A list that contains the reasons that email addresses are automatically added to the suppression list for your account.",
					Optional:    true,
					Computed:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.ValueStringsAre(
							stringvalidator.RegexMatches(regexp.MustCompile("BOUNCE|COMPLAINT"), ""),
						),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that contains information about the suppression list preferences for your account.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: TrackingOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that defines the open and click tracking options for emails that you send using the configuration set.",
		//	  "properties": {
		//	    "CustomRedirectDomain": {
		//	      "description": "The domain to use for tracking open and click events.",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tracking_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: CustomRedirectDomain
				"custom_redirect_domain": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "The domain to use for tracking open and click events.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that defines the open and click tracking options for emails that you send using the configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: VdmOptions
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "An object that contains Virtual Deliverability Manager (VDM) settings for this configuration set.",
		//	  "properties": {
		//	    "DashboardOptions": {
		//	      "additionalProperties": false,
		//	      "description": "Preferences regarding the Dashboard feature.",
		//	      "properties": {
		//	        "EngagementMetrics": {
		//	          "description": "Whether emails sent with this configuration set have engagement tracking enabled.",
		//	          "pattern": "ENABLED|DISABLED",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "EngagementMetrics"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "GuardianOptions": {
		//	      "additionalProperties": false,
		//	      "description": "Preferences regarding the Guardian feature.",
		//	      "properties": {
		//	        "OptimizedSharedDelivery": {
		//	          "description": "Whether emails sent with this configuration set have optimized delivery algorithm enabled.",
		//	          "pattern": "ENABLED|DISABLED",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "OptimizedSharedDelivery"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"vdm_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DashboardOptions
				"dashboard_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EngagementMetrics
						"engagement_metrics": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Whether emails sent with this configuration set have engagement tracking enabled.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("ENABLED|DISABLED"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Preferences regarding the Dashboard feature.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: GuardianOptions
				"guardian_options": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: OptimizedSharedDelivery
						"optimized_shared_delivery": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Whether emails sent with this configuration set have optimized delivery algorithm enabled.",
							Required:    true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.RegexMatches(regexp.MustCompile("ENABLED|DISABLED"), ""),
							}, /*END VALIDATORS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Description: "Preferences regarding the Guardian feature.",
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "An object that contains Virtual Deliverability Manager (VDM) settings for this configuration set.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
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
		Description: "Resource schema for AWS::SES::ConfigurationSet.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::ConfigurationSet").WithTerraformTypeName("awscc_ses_configuration_set")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"custom_redirect_domain":     "CustomRedirectDomain",
		"dashboard_options":          "DashboardOptions",
		"delivery_options":           "DeliveryOptions",
		"engagement_metrics":         "EngagementMetrics",
		"guardian_options":           "GuardianOptions",
		"name":                       "Name",
		"optimized_shared_delivery":  "OptimizedSharedDelivery",
		"reputation_metrics_enabled": "ReputationMetricsEnabled",
		"reputation_options":         "ReputationOptions",
		"sending_enabled":            "SendingEnabled",
		"sending_options":            "SendingOptions",
		"sending_pool_name":          "SendingPoolName",
		"suppressed_reasons":         "SuppressedReasons",
		"suppression_options":        "SuppressionOptions",
		"tls_policy":                 "TlsPolicy",
		"tracking_options":           "TrackingOptions",
		"vdm_options":                "VdmOptions",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
