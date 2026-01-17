build: clean
	go build -o build/plugins/vault-plugin-secrets-solana cmd/vault-plugin-secrets-solana/main.go

clean:
	rm -rf build/ vendor/

test:
	go test -v -cover ./internal/...

vendor: clean
	go mod tidy && go mod vendor

.PHONY: clean test vendor
