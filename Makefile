.PHONY: deps
deps:
	@npm install

node_modules/.bin/tailwindcss: package.json package-lock.json
	@npm install

.PHONY: css
css: node_modules/.bin/tailwindcss
	@./node_modules/.bin/tailwindcss -i ./css/input.css -o ./public/output.css --watch 

.PHONY: templ
templ:
	@templ generate --watch --proxy=http://localhost:3000

.PHONY: build
build:
	@templ generate
	@go build -o bin/moviesdiary main.go

.PHONY: run 
run:
	@if command -v air >/dev/null 2>&1; then \
		air; \
	else \
		echo "air not found; running 'go run .' instead"; \
		go run .; \
	fi

