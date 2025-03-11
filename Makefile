
test:
	go clean -testcache
	go test -v ./...

run:
	go run ./cmd/watchcat.go

.PHONY: build
build:
	rm -rf ./build/
	mkdir -p ./build/
	go build \
		-o ./build/watchcat \
		-gcflags -m=2 \
		./cmd/ 

install:
	ln "$(realpath ./build/watchcat)" -s ~/.local/bin/watchcat

hub_update:
	@hub_ctrl ${HUB_MODE} ln "$(realpath ./build/watchcat)"

