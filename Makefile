DB_DSN := "postgres://postgres:123@localhost:5433/main-tasks?sslmode=disable"
MIGRATE := ~/go/bin/migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

run:
	go run cmd/server/main.go 