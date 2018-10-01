package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestWebServer(t *testing.T) {
	terraformOptions := &terraform.Options {
		// The path to where your Terraform code is located
		TerraformDir: "../tf-templates",
	}
	// At the end of the test, run `terraform destroy`
	defer terraform.Destroy(t, terraformOptions)
	// Run `terraform init` and `terraform apply`
	terraform.InitAndApply(t, terraformOptions)
	// Run `terraform output` to get the value of an output variable
	url := terraform.Output(t, terraformOptions, "url")
	// Verify that we get back a 200 OK with the expected text. It
	// takes ~1 min for the Instance to boot, so retry a few times.
	status := 200
	text := "Hello, World"
	retries := 15
	sleep := 5 * time.Second
	http_helper.HttpGetWithRetry(t, url, status, text, retries, sleep)
}
