module "ecr"{
    source = "../modules/ecr"
    repository_name = "url-shortner"
    environment = var.environment
    tags = {
        Project = "go-url-shortner"
    }
}

output "ecr_repo_url"{
    value = module.ecr.repositrory_url
}