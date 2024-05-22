// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package msk

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_msk_cluster", clusterResource)
}

// clusterResource returns the Terraform awscc_msk_cluster resource.
// This Terraform resource corresponds to the CloudFormation AWS::MSK::Cluster resource.
func clusterResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: BrokerNodeGroupInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BrokerAZDistribution": {
		//	      "maxLength": 9,
		//	      "minLength": 6,
		//	      "type": "string"
		//	    },
		//	    "ClientSubnets": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "ConnectivityInfo": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "PublicAccess": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Type": {
		//	              "maxLength": 23,
		//	              "minLength": 7,
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "VpcConnectivity": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ClientAuthentication": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Sasl": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "Iam": {
		//	                      "additionalProperties": false,
		//	                      "properties": {
		//	                        "Enabled": {
		//	                          "type": "boolean"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "Enabled"
		//	                      ],
		//	                      "type": "object"
		//	                    },
		//	                    "Scram": {
		//	                      "additionalProperties": false,
		//	                      "properties": {
		//	                        "Enabled": {
		//	                          "type": "boolean"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "Enabled"
		//	                      ],
		//	                      "type": "object"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Tls": {
		//	                  "additionalProperties": false,
		//	                  "properties": {
		//	                    "Enabled": {
		//	                      "type": "boolean"
		//	                    }
		//	                  },
		//	                  "required": [
		//	                    "Enabled"
		//	                  ],
		//	                  "type": "object"
		//	                }
		//	              },
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "InstanceType": {
		//	      "maxLength": 32,
		//	      "minLength": 5,
		//	      "type": "string"
		//	    },
		//	    "SecurityGroups": {
		//	      "insertionOrder": false,
		//	      "items": {
		//	        "type": "string"
		//	      },
		//	      "type": "array",
		//	      "uniqueItems": false
		//	    },
		//	    "StorageInfo": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "EBSStorageInfo": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "ProvisionedThroughput": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Enabled": {
		//	                  "type": "boolean"
		//	                },
		//	                "VolumeThroughput": {
		//	                  "type": "integer"
		//	                }
		//	              },
		//	              "type": "object"
		//	            },
		//	            "VolumeSize": {
		//	              "maximum": 16384,
		//	              "minimum": 1,
		//	              "type": "integer"
		//	            }
		//	          },
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "ClientSubnets",
		//	    "InstanceType"
		//	  ],
		//	  "type": "object"
		//	}
		"broker_node_group_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BrokerAZDistribution
				"broker_az_distribution": schema.StringAttribute{ /*START ATTRIBUTE*/
					Optional: true,
					Computed: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(6, 9),
					}, /*END VALIDATORS*/
					PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
						stringplanmodifier.UseStateForUnknown(),
						stringplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ClientSubnets
				"client_subnets": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.RequiresReplace(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: ConnectivityInfo
				"connectivity_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: PublicAccess
						"public_access": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Type
								"type": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.String{ /*START VALIDATORS*/
										stringvalidator.LengthBetween(7, 23),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: VpcConnectivity
						"vpc_connectivity": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ClientAuthentication
								"client_authentication": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Sasl
										"sasl": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Iam
												"iam": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Enabled
														"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
															Required: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
														objectplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
												// Property: Scram
												"scram": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Enabled
														"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
															Required: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
													Optional: true,
													Computed: true,
													PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
														objectplanmodifier.UseStateForUnknown(),
													}, /*END PLAN MODIFIERS*/
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: Tls
										"tls": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Enabled
												"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
													Required: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
												objectplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: InstanceType
				"instance_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(5, 32),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
				// Property: SecurityGroups
				"security_groups": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Optional:    true,
					Computed:    true,
					PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
						generic.Multiset(),
						listplanmodifier.UseStateForUnknown(),
						listplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: StorageInfo
				"storage_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: EBSStorageInfo
						"ebs_storage_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: ProvisionedThroughput
								"provisioned_throughput": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Enabled
										"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
												boolplanmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
										// Property: VolumeThroughput
										"volume_throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
											Optional: true,
											Computed: true,
											PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
												int64planmodifier.UseStateForUnknown(),
											}, /*END PLAN MODIFIERS*/
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
										objectplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: VolumeSize
								"volume_size": schema.Int64Attribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									Validators: []validator.Int64{ /*START VALIDATORS*/
										int64validator.Between(1, 16384),
									}, /*END VALIDATORS*/
									PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
										int64planmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: ClientAuthentication
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Sasl": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Iam": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Enabled": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Scram": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Enabled": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Tls": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CertificateAuthorityArnList": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "type": "string"
		//	          },
		//	          "type": "array",
		//	          "uniqueItems": false
		//	        },
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    },
		//	    "Unauthenticated": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "Enabled": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "required": [
		//	        "Enabled"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"client_authentication": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Sasl
				"sasl": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Iam
						"iam": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Scram
						"scram": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Tls
				"tls": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CertificateAuthorityArnList
						"certificate_authority_arn_list": schema.ListAttribute{ /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
							PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
								generic.Multiset(),
								listplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: Unauthenticated
				"unauthenticated": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: Enabled
						"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ClusterName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 64,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"cluster_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 64),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfigurationInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Arn": {
		//	      "type": "string"
		//	    },
		//	    "Revision": {
		//	      "type": "integer"
		//	    }
		//	  },
		//	  "required": [
		//	    "Revision",
		//	    "Arn"
		//	  ],
		//	  "type": "object"
		//	}
		"configuration_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Arn
				"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
				// Property: Revision
				"revision": schema.Int64Attribute{ /*START ATTRIBUTE*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: CurrentVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The current version of the MSK cluster",
		//	  "type": "string"
		//	}
		"current_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The current version of the MSK cluster",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EncryptionInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "EncryptionAtRest": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "DataVolumeKMSKeyId": {
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "DataVolumeKMSKeyId"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "EncryptionInTransit": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "ClientBroker": {
		//	          "enum": [
		//	            "TLS",
		//	            "TLS_PLAINTEXT",
		//	            "PLAINTEXT"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "InCluster": {
		//	          "type": "boolean"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"encryption_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: EncryptionAtRest
				"encryption_at_rest": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: DataVolumeKMSKeyId
						"data_volume_kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
							Required: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
						objectplanmodifier.RequiresReplaceIfConfigured(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
				// Property: EncryptionInTransit
				"encryption_in_transit": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: ClientBroker
						"client_broker": schema.StringAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							Validators: []validator.String{ /*START VALIDATORS*/
								stringvalidator.OneOf(
									"TLS",
									"TLS_PLAINTEXT",
									"PLAINTEXT",
								),
							}, /*END VALIDATORS*/
							PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
								stringplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: InCluster
						"in_cluster": schema.BoolAttribute{ /*START ATTRIBUTE*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Bool{ /*START PLAN MODIFIERS*/
								boolplanmodifier.UseStateForUnknown(),
								boolplanmodifier.RequiresReplaceIfConfigured(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Optional: true,
					Computed: true,
					PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
						objectplanmodifier.UseStateForUnknown(),
					}, /*END PLAN MODIFIERS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: EnhancedMonitoring
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "DEFAULT",
		//	    "PER_BROKER",
		//	    "PER_TOPIC_PER_BROKER",
		//	    "PER_TOPIC_PER_PARTITION"
		//	  ],
		//	  "maxLength": 23,
		//	  "minLength": 7,
		//	  "type": "string"
		//	}
		"enhanced_monitoring": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(7, 23),
				stringvalidator.OneOf(
					"DEFAULT",
					"PER_BROKER",
					"PER_TOPIC_PER_BROKER",
					"PER_TOPIC_PER_PARTITION",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: KafkaVersion
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"kafka_version": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: LoggingInfo
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "BrokerLogs": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "CloudWatchLogs": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Enabled": {
		//	              "type": "boolean"
		//	            },
		//	            "LogGroup": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "Firehose": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "DeliveryStream": {
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "S3": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "Bucket": {
		//	              "type": "string"
		//	            },
		//	            "Enabled": {
		//	              "type": "boolean"
		//	            },
		//	            "Prefix": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "Enabled"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "BrokerLogs"
		//	  ],
		//	  "type": "object"
		//	}
		"logging_info": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: BrokerLogs
				"broker_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: CloudWatchLogs
						"cloudwatch_logs": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: LogGroup
								"log_group": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: Firehose
						"firehose": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: DeliveryStream
								"delivery_stream": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: S3
						"s3": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: Bucket
								"bucket": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
								// Property: Enabled
								"enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
								// Property: Prefix
								"prefix": schema.StringAttribute{ /*START ATTRIBUTE*/
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
										stringplanmodifier.UseStateForUnknown(),
									}, /*END PLAN MODIFIERS*/
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: NumberOfBrokerNodes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "integer"
		//	}
		"number_of_broker_nodes": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Required: true,
		}, /*END ATTRIBUTE*/
		// Property: OpenMonitoring
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "Prometheus": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "JmxExporter": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "EnabledInBroker": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "EnabledInBroker"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "NodeExporter": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "EnabledInBroker": {
		//	              "type": "boolean"
		//	            }
		//	          },
		//	          "required": [
		//	            "EnabledInBroker"
		//	          ],
		//	          "type": "object"
		//	        }
		//	      },
		//	      "type": "object"
		//	    }
		//	  },
		//	  "required": [
		//	    "Prometheus"
		//	  ],
		//	  "type": "object"
		//	}
		"open_monitoring": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: Prometheus
				"prometheus": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: JmxExporter
						"jmx_exporter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: EnabledInBroker
								"enabled_in_broker": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
						// Property: NodeExporter
						"node_exporter": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: EnabledInBroker
								"enabled_in_broker": schema.BoolAttribute{ /*START ATTRIBUTE*/
									Required: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Optional: true,
							Computed: true,
							PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
								objectplanmodifier.UseStateForUnknown(),
							}, /*END PLAN MODIFIERS*/
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Required: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Optional: true,
			Computed: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: StorageMode
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "LOCAL",
		//	    "TIERED"
		//	  ],
		//	  "maxLength": 6,
		//	  "minLength": 5,
		//	  "type": "string"
		//	}
		"storage_mode": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 6),
				stringvalidator.OneOf(
					"LOCAL",
					"TIERED",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "A key-value pair to associate with a resource.",
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "A key-value pair to associate with a resource.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.Map{ /*START PLAN MODIFIERS*/
				mapplanmodifier.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::MSK::Cluster",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::MSK::Cluster").WithTerraformTypeName("awscc_msk_cluster")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                            "Arn",
		"broker_az_distribution":         "BrokerAZDistribution",
		"broker_logs":                    "BrokerLogs",
		"broker_node_group_info":         "BrokerNodeGroupInfo",
		"bucket":                         "Bucket",
		"certificate_authority_arn_list": "CertificateAuthorityArnList",
		"client_authentication":          "ClientAuthentication",
		"client_broker":                  "ClientBroker",
		"client_subnets":                 "ClientSubnets",
		"cloudwatch_logs":                "CloudWatchLogs",
		"cluster_name":                   "ClusterName",
		"configuration_info":             "ConfigurationInfo",
		"connectivity_info":              "ConnectivityInfo",
		"current_version":                "CurrentVersion",
		"data_volume_kms_key_id":         "DataVolumeKMSKeyId",
		"delivery_stream":                "DeliveryStream",
		"ebs_storage_info":               "EBSStorageInfo",
		"enabled":                        "Enabled",
		"enabled_in_broker":              "EnabledInBroker",
		"encryption_at_rest":             "EncryptionAtRest",
		"encryption_in_transit":          "EncryptionInTransit",
		"encryption_info":                "EncryptionInfo",
		"enhanced_monitoring":            "EnhancedMonitoring",
		"firehose":                       "Firehose",
		"iam":                            "Iam",
		"in_cluster":                     "InCluster",
		"instance_type":                  "InstanceType",
		"jmx_exporter":                   "JmxExporter",
		"kafka_version":                  "KafkaVersion",
		"log_group":                      "LogGroup",
		"logging_info":                   "LoggingInfo",
		"node_exporter":                  "NodeExporter",
		"number_of_broker_nodes":         "NumberOfBrokerNodes",
		"open_monitoring":                "OpenMonitoring",
		"prefix":                         "Prefix",
		"prometheus":                     "Prometheus",
		"provisioned_throughput":         "ProvisionedThroughput",
		"public_access":                  "PublicAccess",
		"revision":                       "Revision",
		"s3":                             "S3",
		"sasl":                           "Sasl",
		"scram":                          "Scram",
		"security_groups":                "SecurityGroups",
		"storage_info":                   "StorageInfo",
		"storage_mode":                   "StorageMode",
		"tags":                           "Tags",
		"tls":                            "Tls",
		"type":                           "Type",
		"unauthenticated":                "Unauthenticated",
		"volume_size":                    "VolumeSize",
		"volume_throughput":              "VolumeThroughput",
		"vpc_connectivity":               "VpcConnectivity",
	})

	opts = opts.WithCreateTimeoutInMinutes(120).WithDeleteTimeoutInMinutes(30)

	opts = opts.WithUpdateTimeoutInMinutes(720)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
