PROTO_DIR := proto
AUTH_PROTO_FILE := auth.proto
USER_PROTO_FILE := user.proto
GEN_DIR := .

include .env

DB_URL := postgres://$(POSTGRES_USERNAME):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_PORT)/$(POSTGRES_DATABASE)?sslmode=disable
MIGRATIONS_DIR := ./migrations

.PHONY: all proto auth_proto user_proto migrate_up migrate_down create_migration

all: proto

proto: auth_proto user_proto

auth_proto: $(GEN_DIR)
	cd $(PROTO_DIR)/auth && protoc -I . --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(AUTH_PROTO_FILE)

user_proto: $(GEN_DIR)
	cd $(PROTO_DIR)/user && protoc -I . --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(USER_PROTO_FILE)

$(GEN_DIR):
	mkdir -p $(GEN_DIR)

migrate_up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose up 1

migrate_down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose down 1

migrate_force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force 20250201000002

create_migration:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) -seq $(name)

