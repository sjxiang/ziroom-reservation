
run: 
	@go run ./cmd/http-server/main.go

seed:
	@go run scripts/seed.go

test:
	@go test -v ./...
