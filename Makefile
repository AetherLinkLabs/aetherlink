BuildDate=`date +%FT%T%z`
Commit=`git describe --always --dirty`

.PHONY: aether
aether:
	go build -ldflags="-s -w" -o ./dist/$@ ./cmd/