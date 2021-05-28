TEST_COVERAGE_DIR=$(shell pwd)/coverage
.PHONY: e2e
e2e:
	mkdir -p ${TEST_COVERAGE_DIR}
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

