#!/usr/bin/env bats

# shellcheck disable=SC2154,SC2012

load ./support/helper

#@test "ArgoCD" {
#  info
#  setup
#
#  apply e2e/out/vendor/argocd
#  test() {
#    kubectl get pods -l karavel.io/component-name=argocd -o json -n argocd | jq '.items[].status.containerStatuses[].ready' | uniq | grep -q true
#  }
#
#  loop_it test 60 10
#  status=${loop_it_result}
#  [ "$status" -eq 0 ]
#}
#
#@test "ArgoCD Applications" {
#  info
#  setup
#
#  apply e2e/out/applications
#  test() {
#    kubectl get applications -o json -n argocd | jq '.items[].status.health.status' | uniq | grep -q Healthy
#  }
#
#  loop_it test 60 10
#  status=${loop_it_result}
#  [ "$status" -eq 0 ]
#}
#
#@test "ArgoCD Projects" {
#  info
#  setup
#
#  apply e2e/out/applications
#  test() {
#    expected=$(ls e2e/out/projects | gre -v kustomization | wc -l)
#    projs=$(kubectl get appprojects -o json -n argocd | jq '.items[].kind' | wc -l)
#    [ "$expected" -eq "$projs" ]
#  }
#
#  loop_it test 60 10
#  status=${loop_it_result}
#  [ "$status" -eq 0 ]
#}
