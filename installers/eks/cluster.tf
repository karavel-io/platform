locals {
  core_pool = {
    name          = "CriticalAddons"
    version       = var.cluster_version
    ami_id        = data.aws_ami.core_pool.image_id
    instance_type = "m5a.large"
    min_size      = 1
    max_size      = 1
    labels = [
      "CriticalAddonsOnly",
    "mikamai.com/reserved-for=critical-addons"]
    taints = [
    "CriticalAddonsOnly:NoSchedule"]
    tags = {
      "mikamai.com/internal" : "true"
    }
  }
  wp = concat(var.worker_pools, [local.core_pool])
  worker_pools = [
    for worker in local.wp :
    {
      name : worker.name,
      ami_id : lookup(worker, "ami_id", null) != null ? worker.ami_id : element(data.aws_ami.worker.*.image_id, index(var.worker_pools.*.name, worker.name)),
      min_size : worker.min_size,
      max_size : worker.max_size,
      version = worker.version,
      instance_type : worker.instance_type,
      tags : [for tag_key, tag_value in worker.tags : {
        key : tag_key,
        value : tag_value,
        propagate_at_launch : true
      }],
      volume_size : 8,
      bootstrap_extra_args : ""
      kubelet_extra_args : <<EOT
%{if length(worker.labels) > 0}--node-labels %{for l in worker.labels}${l},%{endfor}%{endif}
%{if length(worker.taints) > 0}--register-with-taints %{for t in worker.taints}${t},%{endfor}%{endif}
EOT
    }
  ]
}

resource "aws_security_group" "workers" {
  name_prefix = "${var.cluster_name}-"
  vpc_id      = var.vpc
  tags        = var.tags
}

data "aws_ami" "core_pool" {
  filter {
    name = "name"
    values = [
    "amazon-eks-node-${var.cluster_version}-v*"]
  }
  most_recent = true
  owners = [
  "602401143452"]
  # Amazon
}

data "aws_ami" "worker" {
  count = length(var.worker_pools)
  filter {
    name = "name"
    values = [
    "amazon-eks-node-${element(var.worker_pools, count.index).version != null ? element(var.worker_pools, count.index).version : var.cluster_version}-v*"]
  }
  most_recent = true
  owners = [
  "602401143452"]
  # Amazon
}

module "cluster" {
  source  = "terraform-aws-modules/eks/aws"
  version = "13.2.0"

  cluster_create_timeout                         = "30m"
  cluster_delete_timeout                         = "30m"
  cluster_endpoint_private_access                = true
  cluster_endpoint_private_access_cidrs          = var.control_plane_cidrs
  cluster_create_endpoint_private_access_sg_rule = true
  cluster_endpoint_public_access                 = true
  cluster_log_retention_in_days                  = 90
  cluster_enabled_log_types = [
    "api",
    "audit",
    "authenticator",
    "controllerManager",
    "scheduler"
  ]
  cluster_name    = var.cluster_name
  cluster_version = var.cluster_version
  create_eks      = true
  enable_irsa     = true
  iam_path        = "/${var.cluster_name}/"
  kubeconfig_name = var.cluster_name
  subnets         = var.subnets
  tags            = var.tags
  vpc_id          = var.vpc
  worker_additional_security_group_ids           = [aws_security_group.workers.id]
  worker_groups = [
    for worker_pool in local.worker_pools :
    {
      name                 = worker_pool.name
      ami_id               = worker_pool.ami_id
      asg_max_size         = worker_pool.max_size
      asg_min_size         = worker_pool.min_size
      instance_type        = worker_pool.instance_type
      root_volume_size     = worker_pool.volume_size
      additional_security_group_ids = [aws_security_group.workers.id]
      public_ip            = false
      subnets              = var.subnets
      kubelet_extra_args   = replace(replace(chomp(worker_pool.kubelet_extra_args), ", ", " "), "\n", " ")
      tags                 = worker_pool.tags
      bootstrap_extra_args = worker_pool.bootstrap_extra_args
    }
  ]
  worker_sg_ingress_from_port = 22
  write_kubeconfig            = false
}
