.PHONY: build
build: 
	go build -o ./build/app .

run: build
	./build/app

test:
	go clean -testcache
	go test ./... -cover

mockpkg:
	mockgen --source=calculator/new_calculator.go --destination=mock/calculator_mock.go