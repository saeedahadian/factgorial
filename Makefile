.PHONY: run test

run:
	go run . $(ARGS)

build:
	go build -o ./bin/factgorial .

test:
	go test ./...