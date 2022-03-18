.PHONY: migrate run test

migrate:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres007@localhost:5436/digitalrepublic?sslmode=disable up

run:
	@go run cmd/main.go

test:
	@go test -v ./...