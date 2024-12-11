.PHONY: all

all: shell tidy build

shell:
	@echo 'SHELL='$(SHELL)

tidy:
	go mod tidy

build:
	go vet ./cmd/gua64

	CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64 ./cmd/gua64

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64_linux_amd64 ./cmd/gua64

	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64_linux_arm64 ./cmd/gua64

	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64_darwin_amd64 ./cmd/gua64

	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64_darwin_arm64 ./cmd/gua64

	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w" -o ./releases/gua64_windows_amd64.exe ./cmd/gua64