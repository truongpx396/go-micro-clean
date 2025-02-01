PROTO_DIR := proto
AUTH_PROTO_FILE := auth.proto
USER_PROTO_FILE := user.proto
GEN_DIR := .

DB_URL := postgres://dev-user:dev-password@localhost:25432/dev_database?sslmode=disable
MIGRATIONS_DIR := ./migrations

.PHONY: all proto auth_proto user_proto migrate

all: proto

proto: auth_proto user_proto

auth_proto: $(GEN_DIR)
	cd $(PROTO_DIR)/auth && protoc -I . --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(AUTH_PROTO_FILE)

user_proto: $(GEN_DIR)
	cd $(PROTO_DIR)/user && protoc -I . --go_out=$(GEN_DIR) --go_opt=paths=source_relative --go-grpc_out=$(GEN_DIR) --go-grpc_opt=paths=source_relative,require_unimplemented_servers=false $(USER_PROTO_FILE)

$(GEN_DIR):
	mkdir -p $(GEN_DIR)

migrate:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up