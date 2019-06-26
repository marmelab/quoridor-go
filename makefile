.PHONY: test

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

install: ## Install docker environnement
	docker-compose build
	docker-compose run api go get github.com/gorilla/mux

start: ## Start the server
	docker-compose up

stop: ## Stop the server
	docker-compose down

test: ## Tests the API
	docker-compose run api go test -v ./test/...
