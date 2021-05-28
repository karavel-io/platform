#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "Goldpinger" {
  info
  setup

  kubectl create namespace monitoring || true

  install() {
    apply e2e/out/vendor/goldpinger
  }

  test() {
    kubectl get pods -l karavel.io/component-name=goldpinger -o json -n monitoring |jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
