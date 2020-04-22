package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestTerraformSQS(t *testing.T) {
	// Run the test in parallel
	t.Parallel()
	// Generate a unique name for the resource to avoid name colisson
	sqsName := fmt.Sprintf("test-sqs-%s", random.UniqueId())
	// You can randmonize aws regions
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	/* Pass options to terraform e.g directory for terraform code
	Input variables that the module / resource expects
	Environment variables the module / resource uses
	*/
	terraformOptions := &terraform.Options{
		TerraformDir: "../sqs",
		Vars: map[string]interface{}{
			"name": sqsName,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
		NoColor: false,
	}
	// Terraform destroy always runs
	defer terraform.Destroy(t, terraformOptions)
	// Terraform apply to create resources
	terraform.InitAndApply(t, terraformOptions)
	// Check if the outputs from resource creation matches expected output
	output := terraform.Output(t, terraformOptions, "this_sqs_queue_name")
	assert.Equal(t, sqsName, output)
}

func TestTerraformSNS(t *testing.T) {
	// Run the test in parallel
	t.Parallel()
	// Generate a unique name for the resource to avoid name colisson
	snsName := fmt.Sprintf("test-sns-%s", random.UniqueId())
	// You can randmonize aws regions
	awsRegion := aws.GetRandomStableRegion(t, nil, nil)
	/* Pass options to terraform e.g directory for terraform code
	Input variables that the module / resource expects
	Environment variables the module / resource uses
	*/
	terraformOptions := &terraform.Options{
		TerraformDir: "../sns",
		Vars: map[string]interface{}{
			"name": snsName,
		},
		EnvVars: map[string]string{
			"AWS_DEFAULT_REGION": awsRegion,
		},
		NoColor: false,
	}
	// Terraform destroy always runs
	defer terraform.Destroy(t, terraformOptions)
	// Terraform apply to create resources
	terraform.InitAndApply(t, terraformOptions)
	// Check if the outputs from resource creation matches expected output
	output := terraform.Output(t, terraformOptions, "this_sns_topic_arn")
	snsTopicArn := fmt.Sprintf("arn:aws:sns:us-east-1:000000000000:%s", snsName)
	assert.Equal(t, snsTopicArn, output)
}
