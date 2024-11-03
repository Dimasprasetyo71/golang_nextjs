BINARY=Golang Backend NextJS

build:
	go build -o $(BINARY) ./cmd

run: build
	./$(BINARY)

test:
	go test ./...

clean:
	rm -f $(BINARY)

all: build run
