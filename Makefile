E2E_TEST_NAME=*
.PHONY: e2e
e2e:
	bats -t e2e/${E2E_TEST_NAME}.bats

.PHONY: start-kind
kind-start:
	kind create cluster --name karavel-platform-e2e --config e2e/fixtures/kind.yml

.PHONY: stop-kind
kind-stop:
	kind delete cluster --name karavel-platform-e2e


.PHONY: addlicense
addlicense:
	addlicense -c "The Karavel Project" -l apache .

.PHONY: clean
clean:
	rm -rf testbin coverage dist

