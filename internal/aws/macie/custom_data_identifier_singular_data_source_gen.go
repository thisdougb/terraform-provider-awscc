// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package macie

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
	registry.AddDataSourceTypeFactory("awscc_macie_custom_data_identifier", customDataIdentifierDataSourceType)
}

// customDataIdentifierDataSourceType returns the Terraform awscc_macie_custom_data_identifier data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::Macie::CustomDataIdentifier resource type.
func customDataIdentifierDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Custom data identifier ARN.",
			//   "type": "string"
			// }
			Description: "Custom data identifier ARN.",
			Type:        types.StringType,
			Computed:    true,
		},
		"created_at": {
			// Property: CreatedAt
			// CloudFormation resource type schema:
			// {
			//   "description": "Date-time at which the custom data identifier was created.",
			//   "format": "date-time",
			//   "type": "string"
			// }
			Description: "Date-time at which the custom data identifier was created.",
			Type:        types.StringType,
			Computed:    true,
		},
		"deleted": {
			// Property: Deleted
			// CloudFormation resource type schema:
			// {
			//   "description": "Whether the customer data identifier has been deleted.",
			//   "type": "boolean"
			// }
			Description: "Whether the customer data identifier has been deleted.",
			Type:        types.BoolType,
			Computed:    true,
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "Description of custom data identifier.",
			//   "type": "string"
			// }
			Description: "Description of custom data identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"id": {
			// Property: Id
			// CloudFormation resource type schema:
			// {
			//   "description": "Custom data identifier ID.",
			//   "type": "string"
			// }
			Description: "Custom data identifier ID.",
			Type:        types.StringType,
			Computed:    true,
		},
		"ignore_words": {
			// Property: IgnoreWords
			// CloudFormation resource type schema:
			// {
			//   "description": "Words to be ignored.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Words to be ignored.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"keywords": {
			// Property: Keywords
			// CloudFormation resource type schema:
			// {
			//   "description": "Keywords to be matched against.",
			//   "items": {
			//     "type": "string"
			//   },
			//   "type": "array"
			// }
			Description: "Keywords to be matched against.",
			Type:        types.ListType{ElemType: types.StringType},
			Computed:    true,
		},
		"maximum_match_distance": {
			// Property: MaximumMatchDistance
			// CloudFormation resource type schema:
			// {
			//   "description": "Maximum match distance.",
			//   "type": "integer"
			// }
			Description: "Maximum match distance.",
			Type:        types.NumberType,
			Computed:    true,
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of custom data identifier.",
			//   "type": "string"
			// }
			Description: "Name of custom data identifier.",
			Type:        types.StringType,
			Computed:    true,
		},
		"regex": {
			// Property: Regex
			// CloudFormation resource type schema:
			// {
			//   "description": "Regular expression for custom data identifier.",
			//   "type": "string"
			// }
			Description: "Regular expression for custom data identifier.",
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
		Description: "Data Source schema for AWS::Macie::CustomDataIdentifier",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Macie::CustomDataIdentifier").WithTerraformTypeName("awscc_macie_custom_data_identifier")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                    "Arn",
		"created_at":             "CreatedAt",
		"deleted":                "Deleted",
		"description":            "Description",
		"id":                     "Id",
		"ignore_words":           "IgnoreWords",
		"keywords":               "Keywords",
		"maximum_match_distance": "MaximumMatchDistance",
		"name":                   "Name",
		"regex":                  "Regex",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	tflog.Debug(ctx, "Generated schema", "tfTypeName", "awscc_macie_custom_data_identifier", "schema", hclog.Fmt("%v", schema))

	return singularDataSourceType, nil
}
