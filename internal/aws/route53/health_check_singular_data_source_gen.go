// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53

import (
	"context"

	hclog "github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tflog "github.com/hashicorp/terraform-plugin-log"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	providertypes "github.com/hashicorp/terraform-provider-awscc/internal/types"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_route53_health_check", healthCheckDataSourceType)
}

// healthCheckDataSourceType returns the Terraform awscc_route53_health_check data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Route53::HealthCheck resource type.
func healthCheckDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"health_check_config": {
			// Property: HealthCheckConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "A complex type that contains information about the health check.",
			//   "properties": {
			//     "AlarmIdentifier": {
			//       "additionalProperties": false,
			//       "description": "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
			//       "properties": {
			//         "Name": {
			//           "description": "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
			//           "maxLength": 256,
			//           "minLength": 1,
			//           "type": "string"
			//         },
			//         "Region": {
			//           "description": "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "Name",
			//         "Region"
			//       ],
			//       "type": "object"
			//     },
			//     "ChildHealthChecks": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 256,
			//       "type": "array"
			//     },
			//     "EnableSNI": {
			//       "type": "boolean"
			//     },
			//     "FailureThreshold": {
			//       "maximum": 10,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "FullyQualifiedDomainName": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "HealthThreshold": {
			//       "maximum": 256,
			//       "minimum": 0,
			//       "type": "integer"
			//     },
			//     "IPAddress": {
			//       "maxLength": 45,
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "InsufficientDataHealthStatus": {
			//       "enum": [
			//         "Healthy",
			//         "LastKnownStatus",
			//         "Unhealthy"
			//       ],
			//       "type": "string"
			//     },
			//     "Inverted": {
			//       "type": "boolean"
			//     },
			//     "MeasureLatency": {
			//       "type": "boolean"
			//     },
			//     "Port": {
			//       "maximum": 65535,
			//       "minimum": 1,
			//       "type": "integer"
			//     },
			//     "Regions": {
			//       "insertionOrder": false,
			//       "items": {
			//         "type": "string"
			//       },
			//       "maxItems": 64,
			//       "type": "array"
			//     },
			//     "RequestInterval": {
			//       "maximum": 30,
			//       "minimum": 10,
			//       "type": "integer"
			//     },
			//     "ResourcePath": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "RoutingControlArn": {
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     },
			//     "SearchString": {
			//       "maxLength": 255,
			//       "type": "string"
			//     },
			//     "Type": {
			//       "enum": [
			//         "CALCULATED",
			//         "CLOUDWATCH_METRIC",
			//         "HTTP",
			//         "HTTP_STR_MATCH",
			//         "HTTPS",
			//         "HTTPS_STR_MATCH",
			//         "TCP",
			//         "RECOVERY_CONTROL"
			//       ],
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Type"
			//   ],
			//   "type": "object"
			// }
			Description: "A complex type that contains information about the health check.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"alarm_identifier": {
						// Property: AlarmIdentifier
						Description: "A complex type that identifies the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether the specified health check is healthy.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"name": {
									// Property: Name
									Description: "The name of the CloudWatch alarm that you want Amazon Route 53 health checkers to use to determine whether this health check is healthy.",
									Type:        types.StringType,
									Computed:    true,
								},
								"region": {
									// Property: Region
									Description: "For the CloudWatch alarm that you want Route 53 health checkers to use to determine whether this health check is healthy, the region that the alarm was created in.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"child_health_checks": {
						// Property: ChildHealthChecks
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"enable_sni": {
						// Property: EnableSNI
						Type:     types.BoolType,
						Computed: true,
					},
					"failure_threshold": {
						// Property: FailureThreshold
						Type:     types.NumberType,
						Computed: true,
					},
					"fully_qualified_domain_name": {
						// Property: FullyQualifiedDomainName
						Type:     types.StringType,
						Computed: true,
					},
					"health_threshold": {
						// Property: HealthThreshold
						Type:     types.NumberType,
						Computed: true,
					},
					"ip_address": {
						// Property: IPAddress
						Type:     types.StringType,
						Computed: true,
					},
					"insufficient_data_health_status": {
						// Property: InsufficientDataHealthStatus
						Type:     types.StringType,
						Computed: true,
					},
					"inverted": {
						// Property: Inverted
						Type:     types.BoolType,
						Computed: true,
					},
					"measure_latency": {
						// Property: MeasureLatency
						Type:     types.BoolType,
						Computed: true,
					},
					"port": {
						// Property: Port
						Type:     types.NumberType,
						Computed: true,
					},
					"regions": {
						// Property: Regions
						Type:     types.ListType{ElemType: types.StringType},
						Computed: true,
					},
					"request_interval": {
						// Property: RequestInterval
						Type:     types.NumberType,
						Computed: true,
					},
					"resource_path": {
						// Property: ResourcePath
						Type:     types.StringType,
						Computed: true,
					},
					"routing_control_arn": {
						// Property: RoutingControlArn
						Type:     types.StringType,
						Computed: true,
					},
					"search_string": {
						// Property: SearchString
						Type:     types.StringType,
						Computed: true,
					},
					"type": {
						// Property: Type
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"health_check_id": {
			// Property: HealthCheckId
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"health_check_tags": {
			// Property: HealthCheckTags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The key name of the tag.",
			//         "maxLength": 128,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The value for the tag.",
			//         "maxLength": 256,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource.",
			Attributes: providertypes.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The key name of the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The value for the tag.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				providertypes.SetNestedAttributesOptions{},
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
		Description: "Data Source schema for AWS::Route53::HealthCheck",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Route53::HealthCheck").WithTerraformTypeName("awscc_route53_health_check")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"alarm_identifier":                "AlarmIdentifier",
		"child_health_checks":             "ChildHealthChecks",
		"enable_sni":                      "EnableSNI",
		"failure_threshold":               "FailureThreshold",
		"fully_qualified_domain_name":     "FullyQualifiedDomainName",
		"health_check_config":             "HealthCheckConfig",
		"health_check_id":                 "HealthCheckId",
		"health_check_tags":               "HealthCheckTags",
		"health_threshold":                "HealthThreshold",
		"insufficient_data_health_status": "InsufficientDataHealthStatus",
		"inverted":                        "Inverted",
		"ip_address":                      "IPAddress",
		"key":                             "Key",
		"measure_latency":                 "MeasureLatency",
		"name":                            "Name",
		"port":                            "Port",
		"region":                          "Region",
		"regions":                         "Regions",
		"request_interval":                "RequestInterval",
		"resource_path":                   "ResourcePath",
		"routing_control_arn":             "RoutingControlArn",
		"search_string":                   "SearchString",
		"type":                            "Type",
		"value":                           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_route53_health_check", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
