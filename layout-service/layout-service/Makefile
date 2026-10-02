include .env
export

export PROJECT_ROOT=${shell pwd}
POSTGRES_ROOT=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable

env-up:
	docker compose up -d layout-service-postgres

env-down:
	docker compose down layout-service-postgres

env-cleanup: env-down
	rm -rf ${PROJECT_ROOT}/out/pgdata

migrate-up:
	@sleep 2
	migrate -path ./migrations -database "${POSTGRES_ROOT}" up

migrate-create:
	migrate create -ext sql -dir ./migrations -seq ${name}

migrate-down:
	migrate -path ./migrations -database "${POSTGRES_ROOT}" down 1

migrate-version:
	migrate -path ./migrations -database "${POSTGRES_ROOT}" version ${version}

migrate-goto:
	migrate -path ./migrations -database "${POSTGRES_ROOT}" goto ${step}


app-start: env-up migrate-up 
	go run main.go

app-restart: env-cleanup app-start
