locals {
  cluster_name = "karavel-test"
}
provider "aws" {
  region = "eu-west-1"
}

data "aws_availability_zones" "available" {}

module "vpc" {
  source             = "terraform-aws-modules/vpc/aws"
  version            = "2.64.0"
  name               = "karavel_test"
  cidr               = "10.0.0.0/16"
  azs                = data.aws_availability_zones.available.names
  private_subnets    = ["10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"]
  public_subnets     = ["10.0.101.0/24", "10.0.102.0/24", "10.0.103.0/24"]
  enable_nat_gateway = true
  single_nat_gateway = true
  public_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" : "owned",
    "kubernetes.io/role/elb": 1,
  }
  private_subnet_tags = {
    "kubernetes.io/cluster/${local.cluster_name}" : "owned",
    "kubernetes.io/role/internal-elb" : "1"
  }
  tags = {
    "created-by" = "matteo.joliveau"
  }
}

module "eks" {
  source              = "../"
  cluster_name        = local.cluster_name
  cluster_version     = "1.18"
  vpc                 = module.vpc.vpc_id
  subnets             = module.vpc.private_subnets
  control_plane_cidrs = ["0.0.0.0/24"]
  worker_pools = [
    {
      name          = "test"
      version       = "1.18"
      instance_type = "t3a.medium"
      min_size      = 3
      max_size      = 6
      labels        = ["mikamai.com/reserved-for=test"]
      taints        = []
      tags = {
        "created-by" = "matteo.joliveau"
      }
    }
  ]
  tags = {
    "created-by" = "matteo.joliveau"
  }
}

output "cluster_endpoint" {
  description = "The endpoint for the Kubernetes API server"
  value       = module.eks.cluster_endpoint
}

output "kubeconfig" {
  description = "Kubeconfig to connect to the cluster"
  value       = module.eks.kubeconfig
}

