channel = "edge"

component "calico" {
  version = "0.1.0"
  namespace = "kube-system"
}

component "cert-manager" {
  version = "0.1.0"
  namespace = "cert-manager"

  letsencrypt = {
    email = "tech@example.com"
  }
}

component "external-secrets" {
  version = "0.1.0"
  namespace = "external-secrets"

  pollingIntervalMs = 2000

  vault = {
    enable = true
    address = "http://vault.default.svc.cluster.local:8200"
    defaultMountPoint = "kubernetes"
    defaultRole = "karavel"
  }
}

component "external-dns" {
  version = "0.1.0"
  namespace = "external-dns"

  domainFilter = "example.com"

  provider = "cloudflare"
  cloudflare = {
    email = "tech@example.com"
    secret = {
      backend = "vault"
      key = "secret/cloudflare/token"
    }
  }
}

component "velero" {
  version = "0.1.0"
  namespace = "velero"

  backups = {
    provider = "velero.io/aws"
    s3 = {
      bucket = "velero"
      endpoint = "minio:9000"
      region = "eu-west-1"
      encrypted = false
      insecure = true
      pathStyle = true

      accessKeySecret = {
        name = "minio-creds-secret"
        key = "accesskey"
      }

      secretKeySecret = {
        name = "minio-creds-secret"
        key = "secretkey"
      }
    }
  }

  snapshots = {
    enable = false
  }

  restic = {
    enable = true
  }
}

component "ingress-nginx" {
  version = "0.1.0"
  namespace = "ingress-nginx"
}

component "dex" {
  version = "0.1.0"
  namespace = "dex"

  # Params
  publicURL = "https://auth.example.com"

  connectors = [
    {
      type = "mockCallback"
      # Required field for connector id.
      id = "mock"
      # Required field for connector name.
      name = "Example"
    }
  ]

  secret = {
    backend = "secretsManager"
    key = "tools-cluster/dex-secret"
  }
}

component "argocd" {
  version = "0.1.0"
  namespace = "argocd"

  publicURL = "https://cd.example.com"

  git = {
    repo = "git@github.com:mikamai/karavel-local.git"
  }

  secret = {
    backend = "vault"
    key = "secret/argocd-secret"
  }

  credentialsSecret = {
    backend = "vault"
    key = "secret/argocd-pull-creds/sshPrivateKey"
    type = "ssh"
  }

  oidc = {
    config = {
      name = "SSO"
      issuer = "https://sso.example.com"
      client_id = "argocd"
      requested_scopes = [ "openid", "profile", "email", "groups" ]
    }
  }

  notifications = {
    enabled = true
    secret = {
      key = "secret/argocd-notifications"
    }
  }
}

component "grafana" {
  version = "0.1.0"
  namespace = "monitoring"

  publicURL = "https://grafana.example.com"

  dex = {
    issuer = "https://auth.example.com"
  }
}

component "prometheus" {
  version = "0.1.0"
  namespace = "monitoring"

  store = "filesystem"
}

component "loki" {
  version = "0.1.0"
  namespace = "monitoring"

  store = "filesystem"
}

component "tempo" {
  version = "0.1.0"
  namespace = "monitoring"

  store = "filesystem"
}

component "goldpinger" {
  version = "0.1.0"
  namespace = "monitoring"
}

component "olm" {
  version = "0.1.0"
  namespace = "olm"
}
