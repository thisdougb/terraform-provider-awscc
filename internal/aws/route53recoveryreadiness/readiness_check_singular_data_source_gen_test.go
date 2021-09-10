// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package route53recoveryreadiness_test

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-awscc/internal/acctest"
)

func TestAccAWSRoute53RecoveryReadinessReadinessCheckDataSource_basic(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryReadiness::ReadinessCheck", "awscc_route53recoveryreadiness_readiness_check", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.EmptyDataSourceConfig(),
			ExpectError: regexp.MustCompile("Missing required argument"),
		},
	})
}

func TestAccAWSRoute53RecoveryReadinessReadinessCheckDataSource_NonExistent(t *testing.T) {
	td := acctest.NewTestData(t, "AWS::Route53RecoveryReadiness::ReadinessCheck", "awscc_route53recoveryreadiness_readiness_check", "test")

	td.DataSourceTest(t, []resource.TestStep{
		{
			Config:      td.DataSourceWithNonExistentIDConfig(),
			ExpectError: regexp.MustCompile("Not Found"),
		},
	})
}
