// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package medialive

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_medialive_cluster", clusterDataSource)
}

// clusterDataSource returns the Terraform awscc_medialive_cluster data source.
// This Terraform data source corresponds to the CloudFormation AWS::MediaLive::Cluster resource.
func clusterDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the Cluster.",
		//	  "pattern": "^arn:.+:medialive:.+:cluster:.+$",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the Cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ChannelIds
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The MediaLive Channels that are currently running on Nodes in this Cluster.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "MediaLive Channel Ids",
		//	    "type": "string"
		//	  },
		//	  "type": "array"
		//	}
		"channel_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The MediaLive Channels that are currently running on Nodes in this Cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ClusterType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The hardware type for the cluster.",
		//	  "enum": [
		//	    "ON_PREMISES",
		//	    "OUTPOSTS_RACK",
		//	    "OUTPOSTS_SERVER",
		//	    "EC2"
		//	  ],
		//	  "type": "string"
		//	}
		"cluster_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The hardware type for the cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Id
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique ID of the Cluster.",
		//	  "type": "string"
		//	}
		"cluster_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique ID of the Cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: InstanceRoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The IAM role your nodes will use.",
		//	  "pattern": "^arn:.+:iam:.+:role/.+$",
		//	  "type": "string"
		//	}
		"instance_role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The IAM role your nodes will use.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The user-specified name of the Cluster to be created.",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The user-specified name of the Cluster to be created.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: NetworkSettings
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "On premises settings which will have the interface network mappings and default Output logical interface",
		//	  "properties": {
		//	    "DefaultRoute": {
		//	      "description": "Default value if the customer does not define it in channel Output API",
		//	      "type": "string"
		//	    },
		//	    "InterfaceMappings": {
		//	      "description": "Network mappings for the cluster",
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "additionalProperties": false,
		//	        "description": "Network mappings for the cluster",
		//	        "properties": {
		//	          "LogicalInterfaceName": {
		//	            "description": "logical interface name, unique in the list",
		//	            "type": "string"
		//	          },
		//	          "NetworkId": {
		//	            "description": "Network Id to be associated with the logical interface name, can be duplicated in list",
		//	            "type": "string"
		//	          }
		//	        },
		//	        "type": "object"
		//	      },
		//	      "type": "array"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"network_settings": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: DefaultRoute
				"default_route": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Default value if the customer does not define it in channel Output API",
					Computed:    true,
				}, /*END ATTRIBUTE*/
				// Property: InterfaceMappings
				"interface_mappings": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
					NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: LogicalInterfaceName
							"logical_interface_name": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "logical interface name, unique in the list",
								Computed:    true,
							}, /*END ATTRIBUTE*/
							// Property: NetworkId
							"network_id": schema.StringAttribute{ /*START ATTRIBUTE*/
								Description: "Network Id to be associated with the logical interface name, can be duplicated in list",
								Computed:    true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
					}, /*END NESTED OBJECT*/
					Description: "Network mappings for the cluster",
					Computed:    true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "On premises settings which will have the interface network mappings and default Output logical interface",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: State
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current state of the Cluster.",
		//	  "enum": [
		//	    "CREATING",
		//	    "CREATE_FAILED",
		//	    "ACTIVE",
		//	    "DELETING",
		//	    "DELETED"
		//	  ],
		//	  "type": "string"
		//	}
		"state": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current state of the Cluster.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A collection of key-value pairs.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "type": "string"
		//	      }
		//	    },
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
			Description: "A collection of key-value pairs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::MediaLive::Cluster",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MediaLive::Cluster").WithTerraformTypeName("awscc_medialive_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"channel_ids":            "ChannelIds",
		"cluster_id":             "Id",
		"cluster_type":           "ClusterType",
		"default_route":          "DefaultRoute",
		"instance_role_arn":      "InstanceRoleArn",
		"interface_mappings":     "InterfaceMappings",
		"key":                    "Key",
		"logical_interface_name": "LogicalInterfaceName",
		"name":                   "Name",
		"network_id":             "NetworkId",
		"network_settings":       "NetworkSettings",
		"state":                  "State",
		"tags":                   "Tags",
		"value":                  "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
