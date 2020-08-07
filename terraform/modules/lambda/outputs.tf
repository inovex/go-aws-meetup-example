# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
# 
# SPDX-License-Identifier: MIT
output "lambda_invoke_arn" {
  value = aws_lambda_function.lambda.invoke_arn
}

output "lambda_function_name" {
  value = aws_lambda_function.lambda.function_name
}

output "lambda_exec_name" {
  value = aws_iam_role.lambda_exec.name
}
