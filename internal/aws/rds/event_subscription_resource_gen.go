// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package rds

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_rds_event_subscription", eventSubscriptionResource)
}

// eventSubscriptionResource returns the Terraform awscc_rds_event_subscription resource.
// This Terraform resource corresponds to the CloudFormation AWS::RDS::EventSubscription resource.
func eventSubscriptionResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Enabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": true,
		//	  "description": "Specifies whether to activate the subscription. If the event notification subscription isn't activated, the subscription is created but not active.",
		//	  "type": "boolean"
		//	}
		"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies whether to activate the subscription. If the event notification subscription isn't activated, the subscription is created but not active.",
			Optional:    true,
			Computed:    true,
			Default:     booldefault.StaticBool(true),
			PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
				boolplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EventCategories
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A list of event categories for a particular source type (``SourceType``) that you want to subscribe to. You can see a list of the categories for a given source type in the \"Amazon RDS event categories and event messages\" section of the [Amazon RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html) or the [Amazon Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html). You can also see this list by using the ``DescribeEventCategories`` operation.",
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"event_categories": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A list of event categories for a particular source type (``SourceType``) that you want to subscribe to. You can see a list of the categories for a given source type in the \"Amazon RDS event categories and event messages\" section of the [Amazon RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.Messages.html) or the [Amazon Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Events.Messages.html). You can also see this list by using the ``DescribeEventCategories`` operation.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SnsTopicArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the SNS topic created for event notification. SNS automatically creates the ARN when you create a topic and subscribe to it.\n  RDS doesn't support FIFO (first in, first out) topics. For more information, see [Message ordering and deduplication (FIFO topics)](https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html) in the *Amazon Simple Notification Service Developer Guide*.",
		//	  "type": "string"
		//	}
		"sns_topic_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the SNS topic created for event notification. SNS automatically creates the ARN when you create a topic and subscribe to it.\n  RDS doesn't support FIFO (first in, first out) topics. For more information, see [Message ordering and deduplication (FIFO topics)](https://docs.aws.amazon.com/sns/latest/dg/sns-fifo-topics.html) in the *Amazon Simple Notification Service Developer Guide*.",
			Required:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The list of identifiers of the event sources for which events are returned. If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens. It can't end with a hyphen or contain two consecutive hyphens.\n Constraints:\n  +  If ``SourceIds`` are supplied, ``SourceType`` must also be provided.\n  +  If the source type is a DB instance, a ``DBInstanceIdentifier`` value must be supplied.\n  +  If the source type is a DB cluster, a ``DBClusterIdentifier`` value must be supplied.\n  +  If the source type is a DB parameter group, a ``DBParameterGroupName`` value must be supplied.\n  +  If the source type is a DB security group, a ``DBSecurityGroupName`` value must be supplied.\n  +  If the source type is a DB snapshot, a ``DBSnapshotIdentifier`` value must be supplied.\n  +  If the source type is a DB cluster snapshot, a ``DBClusterSnapshotIdentifier`` value must be supplied.\n  +  If the source type is an RDS Proxy, a ``DBProxyName`` value must be supplied.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"source_ids": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The list of identifiers of the event sources for which events are returned. If not specified, then all sources are included in the response. An identifier must begin with a letter and must contain only ASCII letters, digits, and hyphens. It can't end with a hyphen or contain two consecutive hyphens.\n Constraints:\n  +  If ``SourceIds`` are supplied, ``SourceType`` must also be provided.\n  +  If the source type is a DB instance, a ``DBInstanceIdentifier`` value must be supplied.\n  +  If the source type is a DB cluster, a ``DBClusterIdentifier`` value must be supplied.\n  +  If the source type is a DB parameter group, a ``DBParameterGroupName`` value must be supplied.\n  +  If the source type is a DB security group, a ``DBSecurityGroupName`` value must be supplied.\n  +  If the source type is a DB snapshot, a ``DBSnapshotIdentifier`` value must be supplied.\n  +  If the source type is a DB cluster snapshot, a ``DBClusterSnapshotIdentifier`` value must be supplied.\n  +  If the source type is an RDS Proxy, a ``DBProxyName`` value must be supplied.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SourceType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of source that is generating the events. For example, if you want to be notified of events generated by a DB instance, you set this parameter to ``db-instance``. For RDS Proxy events, specify ``db-proxy``. If this value isn't specified, all events are returned.\n Valid Values:``db-instance | db-cluster | db-parameter-group | db-security-group | db-snapshot | db-cluster-snapshot | db-proxy | zero-etl | custom-engine-version | blue-green-deployment``",
		//	  "type": "string"
		//	}
		"source_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of source that is generating the events. For example, if you want to be notified of events generated by a DB instance, you set this parameter to ``db-instance``. For RDS Proxy events, specify ``db-proxy``. If this value isn't specified, all events are returned.\n Valid Values:``db-instance | db-cluster | db-parameter-group | db-security-group | db-snapshot | db-cluster-snapshot | db-proxy | zero-etl | custom-engine-version | blue-green-deployment``",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: SubscriptionName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the subscription.\n Constraints: The name must be less than 255 characters.",
		//	  "maxLength": 255,
		//	  "type": "string"
		//	}
		"subscription_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the subscription.\n Constraints: The name must be less than 255 characters.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An optional array of key-value pairs to apply to this subscription.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Metadata assigned to an Amazon RDS resource consisting of a key-value pair.\n For more information, see [Tagging Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.html) in the *Amazon RDS User Guide* or [Tagging Amazon Aurora and Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html) in the *Amazon Aurora User Guide*.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A key is the required name of the tag. The string value can be from 1 to 128 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "A value is the optional value of the tag. The string value can be from 1 to 256 Unicode characters in length and can't be prefixed with ``aws:`` or ``rds:``. The string can only contain only the set of Unicode letters, digits, white-space, '_', '.', ':', '/', '=', '+', '-', '@' (Java regex: \"^([\\\\p{L}\\\\p{Z}\\\\p{N}_.:/=+\\\\-@]*)$\").",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(0, 256),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An optional array of key-value pairs to apply to this subscription.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
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
		Description: "The ``AWS::RDS::EventSubscription`` resource allows you to receive notifications for Amazon Relational Database Service events through the Amazon Simple Notification Service (Amazon SNS). For more information, see [Using Amazon RDS Event Notification](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Events.html) in the *Amazon RDS User Guide*.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::RDS::EventSubscription").WithTerraformTypeName("awscc_rds_event_subscription")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"enabled":           "Enabled",
		"event_categories":  "EventCategories",
		"key":               "Key",
		"sns_topic_arn":     "SnsTopicArn",
		"source_ids":        "SourceIds",
		"source_type":       "SourceType",
		"subscription_name": "SubscriptionName",
		"tags":              "Tags",
		"value":             "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
