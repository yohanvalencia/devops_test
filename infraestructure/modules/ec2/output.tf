output "public_dns" {
  description = "Outputs the public DNS of the instance"
  value       = aws_instance.ec2.public_dns
}