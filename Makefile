.PHONY: lint test build run run-test-1 run-test-2

target_url_1 = https://pipedream.com/apps/swapi
target_url_2 = https://www.theautomatedtester.co.uk

lint: # runs linter
	golangci-lint run 

test: # runs all the tests along with calculating coverage and race detection
	go test -cover -race ./...

build: # builds the app binary
	go build cmd/main.go

run: # runs the app with the passed in url param
	go run cmd/main.go -url=$(url)

run-target-1: # runs the app with the predefined target url 1
	go run cmd/main.go -url=$(target_url_1)

run-target-2: # runs the app with the predefined target url 2
	go run cmd/main.go -url=$(target_url_2)