resource "aws_ecr_repository" "url-app-repo"{
    name                 = var.repository_name
    image_tag_mutability = "IMMUTABLE"
    encryption_configuration {
        encryption_type = "KMS"
    }
    image_scanning_configuration {
        scan_on_push = true
    }
    tags = merge (
        var.tags,
        {
            Environment = var.environment
        }
    )
}

resource "aws_ecr_lifecycle_policy" "app_cleanup_policy"{
    repository = aws_ecr_repository.url-app-repo.name

    policy = jsonencode({
        rules = [
            {
                rulePriority = 1,
                description = "Keep the latest five images",
                selection = {
                    tagStatus = "any"
                    countType = "imageCountMoreThan"
                    countNumber = 5
                },
                action = {
                    type = "expire"
                }
            },
            {
                rulePriority = 2,
                description = "Delete images older than 14 days",
                selection = {
                    tagStatus = "any"
                    countType = "sinceImagePushed"
                    countUnit = "days"
                    countNumber = 14
                },
                action = {
                    type = "expire"
                }
            }
        ]
    })
}