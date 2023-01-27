generate:
	go generate ./...

tidy:
	go mod tidy

migration:
	go run . database-migration

http:
	go run . restful-api-server

run: generate tidy
	go run github.com/cosmtrek/air restful-api-server