# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
# 
# SPDX-License-Identifier: MIT
provider "aws" {
  region = "eu-central-1"
  version = "~> 2.29"
}

terraform {
  backend "s3" {
    region = "eu-central-1"
    key = "example-service.tfstate"
    encrypt = false
  }
}

locals {
  project_name = "example"
  lambda_path = "${path.module}/function.zip"
}

module "name" {
  source = "git::https://github.com/cloudposse/terraform-null-label.git?ref=tags/0.13.0"
  namespace = "meetup"
  stage = var.stage
  name = local.project_name
}

module "lambda" {
  source = "./modules/lambda"
  base_name = module.name.id
  env = {
    EXAMPLE_ENVIRONMENT = var.stage
    EXAMPLE_LOCAL_MODE = false
  }
  lambda_path = local.lambda_path
}

module "api" {
  source = "./modules/api"
  lambda_function_name = module.lambda.lambda_function_name
  lambda_invoke_arn = module.lambda.lambda_invoke_arn
  stage = var.stage
  base_name = module.name.id
}

resource "aws_dynamodb_table" "item_table" {
  name = "${module.name.id}-items"
  billing_mode = "PAY_PER_REQUEST"
  lifecycle {
    # Prevent TF from changing capacities while autoscaling is active
    ignore_changes = [
      read_capacity,
      write_capacity]
    # Prevent TF from destroying and recreating the table without saving the data
    # prevent_destroy = true
  }
  hash_key = "ID"

  attribute {
    name = "ID"
    type = "S"
  }

  attribute {
    name = "Name"
    type = "S"
  }

  global_secondary_index {
    hash_key = "Name"
    name = "NameIndex"
    projection_type = "ALL"
  }
}
