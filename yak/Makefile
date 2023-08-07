prepare:
	curl -L -s https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 -o $(GOPATH)/bin/dep
	chmod +x $(GOPATH)/bin/dep
	go get -u golang.org/x/lint/golint

deps:
	dep ensure

lint:
	go vet ./...
	golint ./cmd/... ./pkg/...

test:
	GIN_MODE=test go test ./... -coverprofile=coverage.txt -covermode=atomic

run:
	go run yak-webshop.go

build:
	go build .

docker:
	GOOS=linux go build . && docker build . -t=yak-webshop && rm -rf yak-webshop