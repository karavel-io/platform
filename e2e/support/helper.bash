#!/usr/bin/env bats
# shellcheck disable=SC2154,SC2034,SC2086

# Greatly inspired from https://github.com/sighupio/fury-distribution/blob/master/katalog/tests/helper.bash

apply (){
  kustomize build $1 >&2
  kustomize build $1 | kubectl apply -f - 2>&3
}

delete (){
  kustomize build $1 >&2
  kustomize build $1 | kubectl delete -f - 2>&3
}

info(){
  echo -e "${BATS_TEST_NUMBER}: ${BATS_TEST_DESCRIPTION}" >&3
}

setup() {
  if [[ "$BATS_TEST_NUMBER" -eq 1 ]]; then
    kubectl config current-context | grep kind # check that we are running against KinD
    mkdir -p e2e/out
    cp e2e/fixtures/karavel.hcl e2e/out
    karavel render -f e2e/out/karavel.hcl
    find e2e/out/vendor | grep customresourcedefinition | xargs -I {} kubectl apply -f {} 2>&3
  fi
}

loop_it(){
  retry_counter=0
  max_retry=${2:-100}
  wait_time=${3:-2}
  run ${1}
  ko=${status}
  loop_it_result=${ko}
  while [[ ko -ne 0 ]]
  do
    if [ $retry_counter -ge $max_retry ]; then echo "Timeout waiting a condition"; return 1; fi
    sleep ${wait_time} && echo "# waiting..." $retry_counter >&3
    run ${1}
    ko=${status}
    loop_it_result=${ko}
    retry_counter=$((retry_counter + 1))
  done
  return 0
}
