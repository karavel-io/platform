.PHONY: build
build: clean bin
	go build -o ./bin/karavel ./cmd/karavel

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: bin
bin:
	mkdir -p ./bin

.PHONY: test
test:
	go test ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: vet
vet:
	go vet ./...
