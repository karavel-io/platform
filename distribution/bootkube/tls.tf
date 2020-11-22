# === #
# CAs
# === #

resource "tls_private_key" "kubernetes_ca_key" {
  algorithm = "RSA"
  rsa_bits = 4096
}

resource "tls_private_key" "etcd_ca_key" {
  algorithm = "RSA"
  rsa_bits = 4096
}

resource "tls_private_key" "front_proxy_ca_key" {
  algorithm = "RSA"
  rsa_bits = 4096
}

# ================= #
# Etcd certificates
# ================= #

