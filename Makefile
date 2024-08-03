.PHONY: lint build run run-target-1 run-target-2

target_url_1 = https://pipedream.com/apps/swapi
target_url_2 = https://www.theautomatedtester.co.uk

lint: # runs linter
	golangci-lint run 

build: # builds the app binary
	go build -o web-crawler cmd/main.go

run: # runs the app with the passed in url param
	go run cmd/main.go -url=$(url)

run-target-1: # runs the app with the predefined target url 1
	go run cmd/main.go -url=$(target_url_1)

run-target-2: # runs the app with the predefined target url 2
	go run cmd/main.go -url=$(target_url_2)