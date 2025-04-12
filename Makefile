.PHONY: run test build clean install

run:
	go run .

test:
	go test ./...

build:
	go build -o ./bin/Pokedex ./main.go

clean:
	rm -rf ./bin/

install:
	go install .
