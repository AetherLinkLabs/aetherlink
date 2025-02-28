BuildDate=`date +%FT%T%z`
Commit=`git describe --always --dirty`

.PHONY: storage
storage:
	go build -ldflags="-s -w" -o ./dist/$@ ./cmd/storage
