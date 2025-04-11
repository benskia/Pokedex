.PHONY: run test build clean install

run:
	go run .

build:
	go build -o ./bin/Pokedex ./main.go

clean:
	rm -rf ./bin/

install:
	go install .
