generate:
	go generate ./...

tidy:
	go mod tidy

run-database-migration:
	go run . database-migration

run-restful-api-server:
	go run . restful-api-server

run: generate tidy run-database-migration
	go run github.com/cosmtrek/air restful-api-server