## Setup

### Prerequisites
- golang
- docker

### Install localstack
`pip install localstack`

### Start localstack
`localstack start`

### Running The Terratest Tests
`cd tests`

### Run SQS tests
`go test -run TestTerraformSQS`

### Run SNS tests
`go test -run TestTerraformSNS`
