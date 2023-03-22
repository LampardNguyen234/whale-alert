build: build-app
run:
	make build-app && ./wa --config config.json
build-app:
	make swagger && go build -o wa
swagger:
	go install github.com/swaggo/swag/cmd/swag@v1.8.10 && swag init
clean:
	rm -f wa && go clean --modcache && go mod tidy
