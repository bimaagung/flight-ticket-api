DB=postgres://developer:12345678@localhost:5432/flightticketapi?sslmode=disable

run:
	go run cmd/main.go

migrate-create:
	migrate create -ext sql -dir db/migrations

migrate-up:
	migrate -database "${DB}" -path db/migrations up

migrate-down:
	migrate -database "${DB}" -path db/migrations down