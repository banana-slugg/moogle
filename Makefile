css:
	@npx tailwindcss -i ./templates/main.css -o ./public/output.css --minify

build: css
	@go build -o ./bin/server ./cmd/server/main.go

run: build
	@./bin/server