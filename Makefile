clean:
	rm -rf ./bin

bin:
	mkdir -p ./bin

cli: clean bin
	go build -o ./bin/karavel ./cli/cmd/karavel
