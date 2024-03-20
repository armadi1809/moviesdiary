.PHONY: css
css:
	@tailwindcss -i ./css/input.css -o ./public/output.css --watch 

.PHONY: templ
templ:
	@templ generate --watch --proxy=http://localhost:3000

.PHONY: build
build:
	@templ generate
	@go build -o bin/moviesdiary main.go

.PHONY: run 
run:
	air

