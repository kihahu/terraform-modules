module "sns" {
    source                                   = "terraform-aws-modules/sns/aws"
    name                                     = var.name
}