generate:
	go generate ./...

test: generate
test:
	go test -v ./...
