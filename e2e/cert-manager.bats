#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "cert-manager" {
  info
  setup

  kubectl create namespace cert-manager || true

  install() {
    apply e2e/out/vendor/cert-manager
  }

  test() {
    kubectl get pods -l karavel.io/component-name=cert-manager -o json -n cert-manager |jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
