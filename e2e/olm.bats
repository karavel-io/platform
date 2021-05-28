#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "Operator Lifecycle Manager" {
  info
  setup

  kubectl create namespace olm || true
  kubectl create namespace operators || true

  install() {
    apply e2e/out/vendor/olm
  }

  test() {
    kubectl get pods -l karavel.io/component-name=olm -o json -n olm |jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
