#Defining a specific version of terraform's aws provider to keep 
#consistency over time.
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "3.34.0"
    }
  }
}

#Create a module based on the requirements of simplicity and the
#the fact that can be used in a repeatable way. This can also be
#deployed directly from the github project url.

module "ec2" {
  source           = "./modules/ec2"
  private_key_path = "../holded.pem"
  ansible_hosts    = "../application/holded"
}

output "public_dns" {
  description = "Outputs the DNS of the instance"
  value       = module.ec2.public_dns
}