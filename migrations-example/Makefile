include .env
export

MIGRATION_DIR=internal/db/migrations
CONNECTION_STRING=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSL)

importdb:
	docker exec -it postgres_db psql -U root -d master-golang < ./backupdb-master-golang.sql

exportdb:
	docker exec -i postgres_db pg_dump -U root -d master-golang > ./backupdb-master-golang.sql

server:
	go run .

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_DIR) -seq $(NAME)

migrate-up:
	migrate -path $(MIGRATION_DIR) -database "$(CONNECTION_STRING)" up 

#rollback to the last migration
migrate-down:
	migrate -path $(MIGRATION_DIR) -database "$(CONNECTION_STRING)" down 1

#rollback when caution
migrate-force:
	migrate -path $(MIGRATION_DIR) -database "$(CONNECTION_STRING)" force 1

#drop everything
migrate-drop:
	migrate -path $(MIGRATION_DIR) -database "$(CONNECTION_STRING)" drop

#apply specific migration version
migrate-goto:
	migrate -path $(MIGRATION_DIR) -database "$(CONNECTION_STRING)" goto $(VERSION) 



.PHONY: importdb exportdb server migrate-create migrate-up migrate-down migrate-force

