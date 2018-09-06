#####################################################
# VARIABLES
#####################################################
variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "aws_region" {
    description = "The aws region where all the resources are to be created"
    default     = "us-east-1"
}
variable "VpcBlock" {
    description = "The CIDR range for the VPC. This should be a valid private (RFC 1918) CIDR range."
    default     = "192.168.0.0/16"
}
variable "Subnet01Block" {
    description = "CidrBlock for subnet 01 within the VPC"
    default     = "192.168.64.0/18"
}
variable "Subnet02Block" {
    description = "CidrBlock for subnet 02 within the VPC"
    default     = "192.168.128.0/18"
}
variable "Subnet03Block" {
    description = "CidrBlock for subnet 03 within the VPC"
    default     = "192.168.192.0/18"
}

#Data

data "aws_availability_zones" "available" {}

####################################################
# RESOURCES
####################################################

resource "aws_vpc" "vpc" {
  cidr_block           = "${var.VpcBlock}"
  enable_dns_hostnames = "true"
  enable_dns_support   = "true"
  tags {
        Name = "eks-vpc"
    }
}

resource "aws_internet_gateway" "vpc_ig" {
  vpc_id = "${aws_vpc.vpc.id}"
   tags {
        Name = "eks-vpc-ig"
    }
}

resource "aws_route_table" "public" {
  vpc_id           = "${aws_vpc.vpc.id}"
   tags {
        Name = "eks-vpc-rt"
    }
}

resource "aws_route" "public_internet_gateway" {
  route_table_id         = "${aws_route_table.public.id}"
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = "${aws_internet_gateway.vpc_ig.id}"
}

resource "aws_subnet" "subnet01" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "${var.Subnet01Block}"
  availability_zone = "${data.aws_availability_zones.available.names[0]}"
   tags {
        Name = "eks-subnet-1"
    }
}

resource "aws_subnet" "subnet02" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "${var.Subnet02Block}"
  availability_zone = "${data.aws_availability_zones.available.names[1]}"
   tags {
        Name = "eks-subnet-2"
    }
}

resource "aws_subnet" "subnet03" {
  vpc_id            = "${aws_vpc.vpc.id}"
  cidr_block        = "${var.Subnet03Block}"
  availability_zone = "${data.aws_availability_zones.available.names[2]}"
   tags {
        Name = "eks-subnet-3"
    }
}

resource "aws_security_group" "eks_sg" {
  name   = "eks-control-plane-sg"
  vpc_id = "${aws_vpc.vpc.id}"
}

output "vpc_id" {
  value = "${aws_vpc.vpc.id}"
}

output "eks_subnet1" {
  value = ["${aws_subnet.subnet01.id}"]
}

output "eks_subnet2" {
  value = ["${aws_subnet.subnet02.id}"]
}

output "eks_subnet3" {
  value = ["${aws_subnet.subnet03.id}"]
}

output "controlplane_sg_id" {
  value = ["${aws_security_group.eks_sg.id}"]
}