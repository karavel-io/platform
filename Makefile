ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
TEST_COVERAGE_DIR=$(shell pwd)/coverage
test: ## Run tests.
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.8.3/hack/setup-envtest.sh
	mkdir -p ${TEST_COVERAGE_DIR}
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test ./... -ginkgo.v -test.v -coverprofile ${TEST_COVERAGE_DIR}/cover.out

E2E_KUBE_VERSION=1.21.1
.PHONY: e2e
e2e:
	go test ./... -ginkgo.v -test.v -coverprofile ${TEST_COVERAGE_DIR}/cover.out

.PHONY: start-kind
start-kind:
	kind create cluster --name karavel-platform-e2e

.PHONY: stop-kind
stop-kind:
	kind delete cluster --name karavel-platform-e2e


.PHONY: addlicense
addlicense:
	addlicense -c "The Karavel Project" -l apache .

.PHONY: clean
clean:
	rm -rf testbin coverage dist

