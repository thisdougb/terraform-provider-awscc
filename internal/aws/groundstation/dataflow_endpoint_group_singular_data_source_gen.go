// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package groundstation

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_groundstation_dataflow_endpoint_group", dataflowEndpointGroupDataSourceType)
}

// dataflowEndpointGroupDataSourceType returns the Terraform awscc_groundstation_dataflow_endpoint_group data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GroundStation::DataflowEndpointGroup resource type.
func dataflowEndpointGroupDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"endpoint_details": {
			// Property: EndpointDetails
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Endpoint": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Address": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Name": {
			//                 "type": "string"
			//               },
			//               "Port": {
			//                 "type": "integer"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Mtu": {
			//             "type": "integer"
			//           },
			//           "Name": {
			//             "pattern": "",
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "SecurityDetails": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "RoleArn": {
			//             "type": "string"
			//           },
			//           "SecurityGroupIds": {
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           },
			//           "SubnetIds": {
			//             "items": {
			//               "type": "string"
			//             },
			//             "type": "array"
			//           }
			//         },
			//         "type": "object"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"endpoint": {
						// Property: Endpoint
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"address": {
									// Property: Address
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"name": {
												// Property: Name
												Type:     types.StringType,
												Computed: true,
											},
											"port": {
												// Property: Port
												Type:     types.NumberType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"mtu": {
									// Property: Mtu
									Type:     types.NumberType,
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"security_details": {
						// Property: SecurityDetails
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"role_arn": {
									// Property: RoleArn
									Type:     types.StringType,
									Computed: true,
								},
								"security_group_ids": {
									// Property: SecurityGroupIds
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"subnet_ids": {
									// Property: SubnetIds
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
							},
						),
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "Value": {
			//         "pattern": "",
			//         "type": "string"
			//       }
			//     },
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Computed: true,
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Computed: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::GroundStation::DataflowEndpointGroup",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GroundStation::DataflowEndpointGroup").WithTerraformTypeName("awscc_groundstation_dataflow_endpoint_group")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"address":            "Address",
		"arn":                "Arn",
		"endpoint":           "Endpoint",
		"endpoint_details":   "EndpointDetails",
		"id":                 "Id",
		"key":                "Key",
		"mtu":                "Mtu",
		"name":               "Name",
		"port":               "Port",
		"role_arn":           "RoleArn",
		"security_details":   "SecurityDetails",
		"security_group_ids": "SecurityGroupIds",
		"subnet_ids":         "SubnetIds",
		"tags":               "Tags",
		"value":              "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_groundstation_dataflow_endpoint_group", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
