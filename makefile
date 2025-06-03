run:
	go run cmd/weatherapp/main.go

run-race:
	go run -race cmd/weatherapp/main.go

run-performance-test:
	go run cmd/weatherapp/main.go -profile="testing"

run-race-performance-test:
	go run -race cmd/weatherapp/main.go -profile="testing"

test:
	go test -v ./...

test-race:
	go test --race -v ./...

lint:
	golangci-lint run