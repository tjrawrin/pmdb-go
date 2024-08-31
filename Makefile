ENV ?= prod

build:
	@bun run build
	@templ generate
	@if [ "$$ENV" = "dev" ]; then \
		go build -o ./tmp/pmdb-go ./cmd/pmdb-go/main.go; \
	else \
		go build -o pmdb-go ./cmd/pmdb-go/main.go; \
	fi

dev:
	@air -c .air.toml

clean:
	@rm -rf ./tmp
	@rm -rf ./web/dist
	@rm -f ./pmdb-go


PHONY: build dev clean
