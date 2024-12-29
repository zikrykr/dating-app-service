MIGRATE := migrate
MIGRATIONS_DIR := config/db/migration
CONTAINER_NAME := dating-app-service

run: 
	@docker-compose down
	@docker-compose up --build

migrate-up:
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) up

migrate-down:
	@$(MIGRATE) -path $(MIGRATIONS_DIR) -database $(DB_URL) down $(STEP)

create-migration:
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME)
	
