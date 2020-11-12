variable "bucket_name" {
  type = string
  description = "S3 bucket name"
}

variable "public_domain" {
  type = string
  default = "The public domain associated with the repo"
}

variable "cloudflare_zone_id" {
  type = string
  default = "The Cloudflare zone to create records in"
}
