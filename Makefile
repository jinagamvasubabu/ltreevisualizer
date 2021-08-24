GO111MODULE=on

build:
	export GO111MODULE on; \
	go build ./...

build-test:
    export GO111MODULE on; \
    go build ./... && go test ./...

lint: build
	golint -set_exit_status .

test-short: lint
	go test ./... -v -covermode=count -coverprofile=coverage.out

test-coverage: test
	go tool cover -html=coverage.out