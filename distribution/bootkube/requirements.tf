terraform {
  required_version = ">= 0.13.0, < 0.14.0"
  required_providers {
    ct = {
      source  = "poseidon/ct"
      version = "~> 3.14"
    }

    random   = "~> 2.2"
    template = "~> 2.2"
    tls      = "~> 2.2"
  }
}
