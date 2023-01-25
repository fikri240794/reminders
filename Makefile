clean:
	rm -rf transports/restful_api/servers/wire_gen.go docs

generate:
	go generate ./...

tidy:
	go mod tidy

run-database-migration:
	go run . database-migration

run-restful-api-server:
	go run . restful-api-server

run: clean generate tidy run-database-migration run-restful-api-server