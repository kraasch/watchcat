
test_e2e:
	@make build > /dev/null 2>&1 # run a build before the e2e tests.
	go clean -testcache
	go test -v ./e2e/...

test:
	@ echo '##############'
	@ echo '# UNIT TESTS #'
	@ echo '##############'
	make test_unit
	@ echo '##############'
	@ echo '# E2E TESTS  #'
	@ echo '##############'
	make test_e2e

test_unit:
	go clean -testcache
	go test -v ./pkg/...

clean:
	rm -r ./e2e/cli01/

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
