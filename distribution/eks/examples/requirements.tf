terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.14"
    }

    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 1.13"
    }
  }
}
