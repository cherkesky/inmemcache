GO_PROGRAM=inmemcache

run:
	go run main.go cache.go

build:
	go build -o $(GO_PROGRAM) main.go cache.go

build_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM)

build_mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM)

build_windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM).exe

buildandrun_linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM) && ./$(GO_PROGRAM)

buildandrun_mac:
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM) && ./$(GO_PROGRAM)

buildandrun_windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o $(GO_PROGRAM).exe && ./$(GO_PROGRAM).exe

clean:
	rm -f $(GO_PROGRAM)

test:
	go test ./... -v
