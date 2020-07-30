output "api_gw_url" {
  value       = aws_api_gateway_deployment.example.invoke_url
  description = "API Gateway URL"
}
