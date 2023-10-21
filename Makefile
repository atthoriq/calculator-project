.PHONY: build
build: 
	go build -o ./build/app .

run: build
	./build/app

test:
	go clean -testcache
	go test ./... -cover

mockpkg:
	mockgen --source=handler.go --destination=mock/calculator_mock.go