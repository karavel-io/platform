#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "Ingress NGINX" {
  info
  setup

  kubectl create namespace ingress-nginx || true

  install() {
    apply e2e/out/vendor/ingress-nginx
  }

  test() {
    kubectl get pods -l karavel.io/component-name=ingress-nginx -o json -n ingress-nginx | jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
