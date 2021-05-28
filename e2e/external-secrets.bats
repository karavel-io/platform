#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "External Secrets" {
  info
  setup

  kubectl create namespace external-secrets || true

  install() {
    apply e2e/out/vendor/external-secrets
  }

  test() {
    kubectl get pods -l karavel.io/component-name=external-secrets -o json -n external-secrets |jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
