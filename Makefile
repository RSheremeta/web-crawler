.PHONY: lint test build run

lint: # runs linter
	golangci-lint run 

test: # runs all the tests along with calculating coverage and race detection
	go test -cover -race ./...

build: # builds the app binary
	go build .

run: # runs the app
	go run .