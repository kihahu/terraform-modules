module "sqs" {
  source  = "terraform-aws-modules/sqs/aws"
  version = "~> 2.0"

  name = var.name

  tags = {
    Service     = "user"
    Environment = "dev"
  }
}