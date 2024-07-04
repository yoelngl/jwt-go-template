.PHONY: migrate-up migrate-down migrate-status

#Load .env variables
include .env
export $(shell sed 's/=.*//' .env)

MIGRATE=migrations
DB_URL=mysql://$(DB_USER):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)

migrate-up:
	migrate -path $(MIGRATE) -database "$(DB_URL)" -verbose up
	
migrate-down:
	migrate -path $(MIGRATE) -database "$(DB_URL)" -verbose down

migrate-status:
	migrate -path $(MIGRATE) -database "$(DB_URL)" -verbose version
