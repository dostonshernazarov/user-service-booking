proto-gen:
	./scripts/gen-proto.sh

create-migrate:
	migrate create -ext sql -dir ./migrations -seq users-tables

migrate-up:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/usersdb?sslmode=disable' up

migrate-down:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/usersdb?sslmode=disable' down

migrate-force:
	migrate -path ./migrations -database 'postgres://postgres:qwerty@localhost:5432/usersdb?sslmode=disable' force 1
