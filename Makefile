test:
	go test . -v

build:
	go build -o ./bin/todo

run: build
	./bin/todo