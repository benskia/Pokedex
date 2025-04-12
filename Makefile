.PHONY: run test build clean install

run:
	go run ./cmd/Pokedex/

test:
	go test ./...

build:
	go build -o ./bin/Pokedex ./cmd/Pokedex/main.go

clean:
	rm -rf ./bin/

install:
	go install ./cmd/Pokedex/
