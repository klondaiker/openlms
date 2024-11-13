install:
	go mod tidy
	goose up
build:
	go build
run:
	./openlms