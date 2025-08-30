DB_DSN := "postgres://auth_user:auth_pass@localhost:5433/auth?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

migrate:
	$(MIGRATE) up

migrate-down:
	$(MIGRATE) down

lint:
	golangci-lint run --color=always

run:
	go run cmd/server/main.go
