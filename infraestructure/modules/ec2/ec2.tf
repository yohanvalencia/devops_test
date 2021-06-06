#This part will help us get the latest AMI for the current region
#in this case I'll use a Amazon Machine
data "aws_ami" "app_ami" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
}

#Create a security group with the proper ingress and egress rules.
#Ports allowed 80 (HTTP), 8080 (HTTP server.go), 443 (HTTPS) and 22 (SSH). This rules are
#binded to 
resource "aws_security_group" "allow_web" {
  name        = "allow_web_traffic"
  description = "Allow TLS inbound traffic"

  ingress {
    description = "HTTPS with self signed certificate"
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTP server.go"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "HTTP redirect to HTTPS"
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    description = "SSH"
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = var.commonTag
  }
}

#Configure the AWS EC2 instance. It is important to attach the 
#security group policy to the ec2 in order to allow inbound traffic
#in ports 80, 443, and 22.
resource "aws_instance" "ec2" {
  ami                    = data.aws_ami.app_ami.id
  vpc_security_group_ids = [aws_security_group.allow_web.id]
  availability_zone      = var.region_and_az
  instance_type          = var.instanceType
  key_name               = var.keyName
  user_data = <<-EOF
              #! /bin/bash
              sudo yum update -y
              sudo amazon-linux-extras install docker -y
              sudo service docker start
              sudo usermod -a -G docker ec2-user
              sudo systemctl enable docker.service
              sudo reboot
              EOF
  
  tags = {
    Name = var.commonTag
  }

}
