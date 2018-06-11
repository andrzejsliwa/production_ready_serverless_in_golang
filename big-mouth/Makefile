.PHONY = clean build

all: build

clean: ## Clean build
	rm -rf bin/*
	packr clean

build: clean ## Build handler binary
	dep ensure
	packr
	env GOOS=linux go build -o bin/get-index get-index/*.go

deploy: build ## Deploy function
	sls deploy

help: ## This help-view
	@printf "\033[36mHelp: \033[0m\n"
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36mmake %-20s\033[0m%s\n", $$1, $$2}'
