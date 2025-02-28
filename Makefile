BuildDate=`date +%FT%T%z`
Commit=`git describe --always --dirty`

.PHONY: storage
storage:
	go build -ldflags="-s -w" -o ./dist/$@ ./cmd/storage

.PHONY: gateway
gateway:
	go build -ldflags="-s -w" -o ./dist/$@ ./cmd/gateway

.PHONY: gateway-linux
gateway-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./dist/gateway-linux ./cmd/gateway