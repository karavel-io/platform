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

    local = {
      source = "hashicorp/local"
    }

    null = {
      source = "hashicorp/null"
    }

    template = {
      source = "hashicorp/template"
    }

    random = {
      source = "hashicorp/random"
    }
  }
}
