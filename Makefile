clean:
	rm -rf vendor/

test:
	go test -v -cover ./internal/...

vendor: clean
	go mod tidy && go mod vendor

.PHONY: clean test vendor
