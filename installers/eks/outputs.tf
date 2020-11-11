output "cluster_endpoint" {
  description = "The endpoint for the Kubernetes API server"
  value       = data.aws_eks_cluster.cluster.endpoint
}

output "kubeconfig" {
  description = "Generated Kubeconfig to connect to the cluster"
  sensitive   = true
  value       = <<EOT
apiVersion: v1
kind: Config
clusters:
- cluster:
    server: ${data.aws_eks_cluster.cluster.endpoint}
    certificate-authority-data: ${data.aws_eks_cluster.cluster.certificate_authority.0.data}
  name: eks-${var.cluster_name}
contexts:
- context:
    cluster: eks-${var.cluster_name}
    user: eks-${var.cluster_name}-user
  name: eks-${var.cluster_name}
current-context: eks-${var.cluster_name}
preferences: {}
users:
- name: eks-${var.cluster_name}-user
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      command: aws
      args:
        - "eks"
        - "get-token"
        - "--cluster-name"
        - "${var.cluster_name}"
EOT
}

// Uncomment when the EKS module stops using `aws-iam-authenticator` instead of the aws-cli
//output "kubeconfig" {
//  description = "Generated Kubeconfig to connect to the cluster"
//  value       = module.cluster.kubeconfig
//}
