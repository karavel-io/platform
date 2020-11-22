variable "cluster_name" {
  type        = string
  description = "EKS cluster name. Must be unique per AWS account"
}

variable "cluster_version" {
  type        = string
  description = "Kubernetes version. Values are specific for each cloud provider"
}

variable "vpc" {
  type        = string
  description = "VPC where the cluster will be located"
}

variable "subnets" {
  type        = list(string)
  description = "Subnets where the cluster will be located"
}

variable "control_plane_cidrs" {
  type        = list(string)
  description = "CIDRs allowed to access the control plane endpoint"
}

variable "worker_pools" {
  description = "List of worker pool configurations"
  type = list(object({
    name          = string
    version       = string
    instance_type = string
    min_size      = number
    max_size      = number
    labels        = list(string)
    taints        = list(string)
    tags          = map(string)
  }))
  default = []
}

variable "tags" {
  type        = map(string)
  description = "Tags applied to the created cloud resources"
  default     = {}
}
