variable "region_and_az" {
    default="eu-west-1a"
}

variable "instanceType" {
    default="t2.micro"
}

variable "keyName" {
    default="holded"
}

variable "commonTag" {
    default="holded"
}

variable "ssh_user" {
    default="ec2-user"
}

variable "private_key_path" {
    default="holded.pem"
}

variable "ansible_hosts" {
    default="."
}