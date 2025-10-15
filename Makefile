.PHONY: run

run:
	go run . $(ARGS)

build:
	go build -o ./bin/factgorial .