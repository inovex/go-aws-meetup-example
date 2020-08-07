# SPDX-FileCopyrightText: 2020 inovex GmbH <https://www.inovex.de>
# 
# SPDX-License-Identifier: MIT
output "api_gw_url" {
  value       = aws_api_gateway_deployment.example.invoke_url
  description = "API Gateway URL"
}
