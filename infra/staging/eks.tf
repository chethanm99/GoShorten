terraform {
  required_version = ">= 1.2.0"

  backend "s3" {
    bucket         = "url-shortner-terraform-state-chethan99-devops"  # <- your bucket
    key            = "staging/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "terraform-locks"
    encrypt        = true
  }
}

provider "aws" {
  region = var.aws_region
}

module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "20.1.10"  

  cluster_name    = "go-url-shortner-staging"
  cluster_version = "1.29"

  vpc_create = true
  vpc_cidr   = "10.0.0.0/16"
  
  subnets = [
    "10.0.1.0/24",
    "10.0.2.0/24"
  ]

  node_groups = {
    ng_default = {
      desired_capacity = 1
      max_capacity     = 2
      min_capacity     = 1

      instance_types = ["t3.small"]
      key_name       = var.ssh_key_name  
    }
  }

  manage_aws_auth = true
  enable_irsa     = true
  tags = {
    Environment = var.environment
    Project     = "go-url-shortner"
  }
}