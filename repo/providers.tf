terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "~> 3.14"
    }
    cloudflare = {
      source = "cloudflare/cloudflare"
      version = "~> 2.13"
    }
  }
}

provider "aws" {
  region = "eu-west-1"
}

provider "aws" {
  alias = "us-east-1"
  region = "us-east-1"
}

provider "cloudflare" {}
