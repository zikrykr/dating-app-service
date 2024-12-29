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
	
mockgen:
	@go install github.com/golang/mock/mockgen
	@mockgen -destination=mock/mock_auth_repository.go -package=mock -source=internal/auth/port/repository.go IAuthRepo
	@mockgen -destination=mock/mock_premium_repository.go -package=mock -source=internal/premium/port/repository.go IPremiumRepo
	@mockgen -destination=mock/mock_recommendations_repository.go -package=mock -source=internal/recommendations/port/repository.go IRecommendationRepo
	@mockgen -destination=mock/mock_swipe_repository.go -package=mock -source=internal/swipe/port/repository.go ISwipeRepository
	@mockgen -destination=mock/mock_auth_service.go -package=mock -source=internal/auth/port/service.go
	@mockgen -destination=mock/mock_recommendation_service.go -package=mock -source=internal/recommendations/port/service.go IRecommendationService
	@mockgen -destination=mock/mock_premium_service.go -package=mock -source=internal/premium/port/service.go IPremiumService
	@mockgen -destination=mock/mock_swipe_service.go -package=mock -source=internal/swipe/port/service.go ISwipeService
	@mockgen -destination=mock/mock_auth_handler.go -package=mock -source=internal/auth/port/handler.go
	@mockgen -destination=mock/mock_premium_handler.go -package=mock -source=internal/premium/port/handler.go IPremiumHandler
	@mockgen -destination=mock/mock_recommendation_handler.go -package=mock -source=internal/recommendations/port/handler.go IRecommendationHandler
	@mockgen -destination=mock/mock_swipe_handler.go -package=mock -source=internal/swipe/port/handler.go ISwipeHandler