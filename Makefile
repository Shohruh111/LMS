
run:
	go run cmd/main.go

swag-init:
	swag init -g api/api.go -o api/docs

migration-up:
	migrate -path ./migration/postgres -database 'postgres://postgres:nfrpSoH235HFHEJhRcnaXCuu9BoGX7K4@localhost:5432/login?sslmode=disable' up

migration-down:
	migrate -path ./migration/postgres -database 'postgres://postgres:1234@localhost:5432/login?sslmode=disable' down


