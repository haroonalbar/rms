buid:
	@go build . -o ./bin/rms

run: build
	@./bin/rms

test:
	@go test ./... 
