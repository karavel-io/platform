terraform {
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = "~> 1.13"
    }

    kustomization = {
      source = "kbst/kustomization"
      version = "~> 0.2.2"
    }
  }
}
