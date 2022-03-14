.PHONY: migrate migratewindows run

migrate:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path=/migrations \
		-database postgres://postgres:postgres007@localhost:5436/digitalrepublic?sslmode=disable up

migratewindows:
	@docker container run --rm --name migrate -v ${PWD}/migrations:/migrations --network host migrate/migrate \
    -path ./migrations \
		-database postgres://postgres:postgres007@localhost:5436/digitalrepublic?sslmode=disable up

run:
	@go run cmd/main.go