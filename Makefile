run: main.go
	@gofmt -w . && go run .
	
f:
	@gofmt -w .
	
test:
	go test ./...