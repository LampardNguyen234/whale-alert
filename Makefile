build: build-app
run:
	make build-app && ./wa --config config.json
build-app:
	go build -o wa
clean:
	rm -f ns && go clean --modcache && go mod tidy
