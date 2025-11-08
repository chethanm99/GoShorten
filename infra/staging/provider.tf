terraform {
    required_version = ">= 1.2.0"

    backend "s3" {
        bucket = "url-shortner-terraform-state-chethan99-devops"
        key    = "staging/terraform.tfstate"
        region = "us-east-1"
        dynamodb_table = "url-shortner-terraform-locks"
        encrypt = true
    }
}

provider "aws"{
    region = "us-east-1"
}