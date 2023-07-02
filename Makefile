test:
	go test . -v

build:
	go build ./cmd/todo/. -o ./bin/todo

run: build
	./bin/todo