# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
# 
# SPDX-License-Identifier: MIT
####################
# Lambda
resource "aws_lambda_function" "lambda" {
  function_name = "${var.base_name}_lambda"

  filename = var.lambda_path

  handler = "main"
  runtime = "go1.x"

  role = aws_iam_role.lambda_exec.arn

  source_code_hash = filebase64sha256(var.lambda_path)

  timeout = 20

  depends_on = [
    aws_iam_role_policy_attachment.lambda_logs,
    aws_iam_role_policy_attachment.dynamodb
  ]

  environment {
    variables = var.env
  }
}



resource "aws_iam_role" "lambda_exec" {
  name = "${var.base_name}_lambda_role"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

####################
# Logging
resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.lambda_logging.arn
}

resource "aws_iam_policy" "lambda_logging" {
  name        = "${var.base_name}_lambda_logging"
  path        = "/"
  description = "IAM policy for logging from a lambda"

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "logs:CreateLogGroup",
        "logs:CreateLogStream",
        "logs:PutLogEvents"
      ],
      "Resource": "arn:aws:logs:*:*:*",
      "Effect": "Allow"
    }
  ]
}
EOF
}

####################
# Dynamo DB access
resource "aws_iam_role_policy_attachment" "dynamodb" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.dynamodb.arn
}

resource "aws_iam_policy" "dynamodb" {
  name        = "${var.base_name}_dynamodb"
  path        = "/"
  description = "IAM policy for dynamo DB access"

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Action": [
                "dynamodb:*"
            ],
            "Effect": "Allow",
            "Resource": "*"
        }
    ]
}
EOF
}
