GOOSE_DRIVER=postgres
GOOSE_DIR=db/migrations
GOOSE_DBSTRING="host=localhost user=root password=password dbname=databasae port=5434 sslmode=disable"

setup: install db_schema dbml_publish
install:
	go install github.com/pressly/goose/v3/cmd/goose@v3.17.0
	npm install -g dbdocs @dbml/cli
	go install github.com/swaggo/swag/cmd/swag@v1.16.2
	go get github.com/joho/godotenv@latest
	go get -u github.com/rs/zerolog/log
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/postgres
	go get -u github.com/pressly/goose/v3/cmd/goose

db_schema:
	dbml2sql --postgres -o docs/db/schema.sql docs/db/db.dbml
dbml_publish:
	dbdocs login
	dbdocs build docs/db.dbml


goose_operations: gm gs gu guo gd gdo
gm: # goose_migration
	 goose -dir $(GOOSE_DIR) create $(name) sql
gs: # goose_status
	 goose -dir $(GOOSE_DIR) postgres $(GOOSE_DBSTRING) status
gu: # goose_up
	goose -dir $(GOOSE_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up
guo: # goose_up_one:
	goose -dir $(GOOSE_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) up-by-one
gd: # goose_down
	goose -dir $(GOOSE_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) reset
gdo: # goose_down_one
	goose -dir $(GOOSE_DIR) $(GOOSE_DRIVER) $(GOOSE_DBSTRING) down

db_seed:
	go run db/seeders/main.go


testers: mock test
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/techmadeeazy/agency-banking-core/db/sqlc Store
test:
	go test -v -cover -short ./...


server:
	go run main.go

swag:
	swag init

.PHONY: setup goose_operations testers server seed_db swag

