run: build
	@bin\snippet

build:
	@go build -o bin\snippet .

css:
	npx tailwindcss -i views\css\app.css -o public\styles.css --watch