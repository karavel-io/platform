#!/usr/bin/env bats

# shellcheck disable=SC2154

load ./support/helper

@test "Calico" {
  info
  setup

  install() {
    apply e2e/out/vendor/calico
  }

  test() {
    kubectl get pods -l karavel.io/component-name=calico -o json -n kube-system |jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
  }

  loop_it install 60 10
  loop_it test 60 10
  status=${loop_it_result}
  [ "$status" -eq 0 ]
}
