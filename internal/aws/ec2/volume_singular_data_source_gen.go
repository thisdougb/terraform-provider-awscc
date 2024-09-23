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
	registry.AddDataSourceFactory("awscc_ec2_volume", volumeDataSource)
}

// volumeDataSource returns the Terraform awscc_ec2_volume data source.
// This Terraform data source corresponds to the CloudFormation AWS::EC2::Volume resource.
func volumeDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AutoEnableIO
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.",
		//	  "type": "boolean"
		//	}
		"auto_enable_io": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the volume is auto-enabled for I/O operations. By default, Amazon EBS disables I/O to the volume from attached EC2 instances when it determines that a volume's data is potentially inconsistent. If the consistency of the volume is not a concern, and you prefer that the volume be made available immediately if it's impaired, you can configure the volume to automatically enable I/O.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: AvailabilityZone
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``.",
		//	  "type": "string"
		//	}
		"availability_zone": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ID of the Availability Zone in which to create the volume. For example, ``us-east-1a``.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Encrypted
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether the volume should be encrypted. The effect of setting the encryption state to ``true`` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*.\n Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances).",
		//	  "type": "boolean"
		//	}
		"encrypted": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether the volume should be encrypted. The effect of setting the encryption state to ``true`` depends on the volume origin (new or from a snapshot), starting encryption state, ownership, and whether encryption by default is enabled. For more information, see [Encryption by default](https://docs.aws.amazon.com/ebs/latest/userguide/work-with-ebs-encr.html#encryption-by-default) in the *Amazon EBS User Guide*.\n Encrypted Amazon EBS volumes must be attached to instances that support Amazon EBS encryption. For more information, see [Supported instance types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-encryption-requirements.html#ebs-encryption_supported_instances).",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Iops
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The number of I/O operations per second (IOPS). For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.\n The following are the supported values for each volume type:\n  +   ``gp3``: 3,000 - 16,000 IOPS\n  +   ``io1``: 100 - 64,000 IOPS\n  +   ``io2``: 100 - 256,000 IOPS\n  \n For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html). On other instances, you can achieve performance up to 32,000 IOPS.\n This parameter is required for ``io1`` and ``io2`` volumes. The default for ``gp3`` volumes is 3,000 IOPS. This parameter is not supported for ``gp2``, ``st1``, ``sc1``, or ``standard`` volumes.",
		//	  "type": "integer"
		//	}
		"iops": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The number of I/O operations per second (IOPS). For ``gp3``, ``io1``, and ``io2`` volumes, this represents the number of IOPS that are provisioned for the volume. For ``gp2`` volumes, this represents the baseline performance of the volume and the rate at which the volume accumulates I/O credits for bursting.\n The following are the supported values for each volume type:\n  +   ``gp3``: 3,000 - 16,000 IOPS\n  +   ``io1``: 100 - 64,000 IOPS\n  +   ``io2``: 100 - 256,000 IOPS\n  \n For ``io2`` volumes, you can achieve up to 256,000 IOPS on [instances built on the Nitro System](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html). On other instances, you can achieve performance up to 32,000 IOPS.\n This parameter is required for ``io1`` and ``io2`` volumes. The default for ``gp3`` volumes is 3,000 IOPS. This parameter is not supported for ``gp2``, ``st1``, ``sc1``, or ``standard`` volumes.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: KmsKeyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The identifier of the kms-key-long to use for Amazon EBS encryption. If ``KmsKeyId`` is specified, the encrypted state must be ``true``.\n If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.\n Alternatively, if you want to specify a different key, you can specify one of the following:\n  +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.\n  +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.",
		//	  "type": "string"
		//	}
		"kms_key_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The identifier of the kms-key-long to use for Amazon EBS encryption. If ``KmsKeyId`` is specified, the encrypted state must be ``true``.\n If you omit this property and your account is enabled for encryption by default, or *Encrypted* is set to ``true``, then the volume is encrypted using the default key specified for your account. If your account does not have a default key, then the volume is encrypted using the aws-managed-key.\n Alternatively, if you want to specify a different key, you can specify one of the following:\n  +  Key ID. For example, 1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Key alias. Specify the alias for the key, prefixed with ``alias/``. For example, for a key with the alias ``my_cmk``, use ``alias/my_cmk``. Or to specify the aws-managed-key, use ``alias/aws/ebs``.\n  +  Key ARN. For example, arn:aws:kms:us-east-1:012345678910:key/1234abcd-12ab-34cd-56ef-1234567890ab.\n  +  Alias ARN. For example, arn:aws:kms:us-east-1:012345678910:alias/ExampleAlias.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: MultiAttachEnabled
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "Indicates whether Amazon EBS Multi-Attach is enabled.\n  CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.",
		//	  "type": "boolean"
		//	}
		"multi_attach_enabled": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "Indicates whether Amazon EBS Multi-Attach is enabled.\n  CFNlong does not currently support updating a single-attach volume to be multi-attach enabled, updating a multi-attach enabled volume to be single-attach, or updating the size or number of I/O operations per second (IOPS) of a multi-attach enabled volume.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: OutpostArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Outpost.",
		//	  "type": "string"
		//	}
		"outpost_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Outpost.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Size
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.\n The following are the supported volumes sizes for each volume type:\n  +   ``gp2`` and ``gp3``: 1 - 16,384 GiB\n  +   ``io1``: 4 - 16,384 GiB\n  +   ``io2``: 4 - 65,536 GiB\n  +   ``st1`` and ``sc1``: 125 - 16,384 GiB\n  +   ``standard``: 1 - 1024 GiB",
		//	  "type": "integer"
		//	}
		"size": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The size of the volume, in GiBs. You must specify either a snapshot ID or a volume size. If you specify a snapshot, the default is the snapshot size. You can specify a volume size that is equal to or larger than the snapshot size.\n The following are the supported volumes sizes for each volume type:\n  +   ``gp2`` and ``gp3``: 1 - 16,384 GiB\n  +   ``io1``: 4 - 16,384 GiB\n  +   ``io2``: 4 - 65,536 GiB\n  +   ``st1`` and ``sc1``: 125 - 16,384 GiB\n  +   ``standard``: 1 - 1024 GiB",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: SnapshotId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.",
		//	  "type": "string"
		//	}
		"snapshot_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The snapshot from which to create the volume. You must specify either a snapshot ID or a volume size.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags to apply to the volume during creation.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Specifies a tag. For more information, see [Resource tags](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The tag key.",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The tag value.",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": false
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag key.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The tag value.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags to apply to the volume during creation.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Throughput
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The throughput to provision for a volume, with a maximum of 1,000 MiB/s.\n This parameter is valid only for ``gp3`` volumes. The default value is 125.\n Valid Range: Minimum value of 125. Maximum value of 1000.",
		//	  "type": "integer"
		//	}
		"throughput": schema.Int64Attribute{ /*START ATTRIBUTE*/
			Description: "The throughput to provision for a volume, with a maximum of 1,000 MiB/s.\n This parameter is valid only for ``gp3`` volumes. The default value is 125.\n Valid Range: Minimum value of 125. Maximum value of 1000.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VolumeId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "",
		//	  "type": "string"
		//	}
		"volume_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: VolumeType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The volume type. This parameter can be one of the following values:\n  +  General Purpose SSD: ``gp2`` | ``gp3`` \n  +  Provisioned IOPS SSD: ``io1`` | ``io2`` \n  +  Throughput Optimized HDD: ``st1`` \n  +  Cold HDD: ``sc1`` \n  +  Magnetic: ``standard`` \n  \n For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html).\n Default: ``gp2``",
		//	  "type": "string"
		//	}
		"volume_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The volume type. This parameter can be one of the following values:\n  +  General Purpose SSD: ``gp2`` | ``gp3`` \n  +  Provisioned IOPS SSD: ``io1`` | ``io2`` \n  +  Throughput Optimized HDD: ``st1`` \n  +  Cold HDD: ``sc1`` \n  +  Magnetic: ``standard`` \n  \n For more information, see [Amazon EBS volume types](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volume-types.html).\n Default: ``gp2``",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EC2::Volume",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EC2::Volume").WithTerraformTypeName("awscc_ec2_volume")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"auto_enable_io":       "AutoEnableIO",
		"availability_zone":    "AvailabilityZone",
		"encrypted":            "Encrypted",
		"iops":                 "Iops",
		"key":                  "Key",
		"kms_key_id":           "KmsKeyId",
		"multi_attach_enabled": "MultiAttachEnabled",
		"outpost_arn":          "OutpostArn",
		"size":                 "Size",
		"snapshot_id":          "SnapshotId",
		"tags":                 "Tags",
		"throughput":           "Throughput",
		"value":                "Value",
		"volume_id":            "VolumeId",
		"volume_type":          "VolumeType",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
