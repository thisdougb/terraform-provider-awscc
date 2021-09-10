// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package location

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
	registry.AddDataSourceTypeFactory("awscc_location_map", mapDataSourceType)
}

// mapDataSourceType returns the Terraform awscc_location_map data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Location::Map resource type.
func mapDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"configuration": {
			// Property: Configuration
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Style": {
			//       "maxLength": 100,
			//       "minLength": 1,
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "Style"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"style": {
						// Property: Style
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"create_time": {
			// Property: CreateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
		"data_source": {
			// Property: DataSource
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"map_arn": {
			// Property: MapArn
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 1600,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"map_name": {
			// Property: MapName
			// CloudFormation resource type schema:
			// {
			//   "maxLength": 100,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"pricing_plan": {
			// Property: PricingPlan
			// CloudFormation resource type schema:
			// {
			//   "enum": [
			//     "RequestBasedUsage",
			//     "MobileAssetTracking",
			//     "MobileAssetManagement"
			//   ],
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"update_time": {
			// Property: UpdateTime
			// CloudFormation resource type schema:
			// {
			//   "description": "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The datetime value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ss.sssZ)",
			Type:        types.StringType,
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::Location::Map",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Location::Map").WithTerraformTypeName("awscc_location_map")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"configuration": "Configuration",
		"create_time":   "CreateTime",
		"data_source":   "DataSource",
		"description":   "Description",
		"map_arn":       "MapArn",
		"map_name":      "MapName",
		"pricing_plan":  "PricingPlan",
		"style":         "Style",
		"update_time":   "UpdateTime",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_location_map", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
