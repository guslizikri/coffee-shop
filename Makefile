APP=coffeeshopbe
BUILD="./build/$(APP)"
DB_DRIVER=postgres
DB_SOURCE="postgresql://admin:admin@localhost/coffeeshop?sslmode=disable&search_path=public"
MIGRATIONS_DIR=./migrations
# https://github.com/golang-migrate/migrate/tree/master/cmd/migrate


install:
	go get -u ./... && go mod tidy

build:
	CGO_ENABLED=0 GOOS=linux go build -o ${BUILD}
# go build -o "./build/coffeeshopbe.exe" ./cmd/main.go
test:
	go test -cover -v ./...

migrate-init:
	migrate create -dir ${MIGRATIONS_DIR} -ext sql $(name)

migrate-up:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose up

migrate-down:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} -verbose down

migrate-fix:
	migrate -path ${MIGRATIONS_DIR} -database ${DB_SOURCE} force 0