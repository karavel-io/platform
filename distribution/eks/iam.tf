data "aws_iam_policy_document" "external_secrets_assume_role" {
  statement {
    actions = ["sts:AssumeRoleWithWebIdentity"]
    principals {
      type = "Federated"
      identifiers = [module.cluster.oidc_provider_arn]
    }
    condition {
      test = "StringEquals"
      values = ["sts.amazonaws.com"]
      variable = "${module.cluster.cluster_oidc_issuer_url}:aud"
    }
  }
}

resource "aws_iam_role" "external_secrets" {
  name = "${var.cluster_name}-external-secrets"
  assume_role_policy = data.aws_iam_policy_document.external_secrets_assume_role.json
}

data "aws_iam_policy_document" "secrets_manager_read_only" {
  statement {
    actions = [
      "secretsmanager:GetResourcePolicy",
      "secretsmanager:GetSecretValue",
      "secretsmanager:DescribeSecret",
      "secretsmanager:ListSecretVersionIds"
    ]
    resources = ["*"]
  }

  statement {
    actions = ["ssm:GetParameter"]
    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "secrets_manager_read_only" {
  role = aws_iam_role.external_secrets.id

  policy = data.aws_iam_policy_document.secrets_manager_read_only.json
}
