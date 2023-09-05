build:
	@go mod tidy
	@go build -v -o ./cmd/NGC_echo ./cmd/NGC_echo.go

run:
	@make build
	@./cmd/NGC_echo

migrate:
	@migrate -database "postgres://haryo.assyafah@localhost:5432/ngc_echo?sslmode=disable" -path migrations down
	@migrate -database "postgres://haryo.assyafah@localhost:5432/ngc_echo?sslmode=disable" -path migrations up
