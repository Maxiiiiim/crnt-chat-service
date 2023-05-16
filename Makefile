DST_DIR_RELATIVE=./pkg

SWAGGER_FILES=$(shell find $(DST_DIR_RELATIVE)/ \
                      -name '*.swagger.json*')

copy-swagger:
	cp $(SWAGGER_FILES) ./swagger

all: bin/chat

bin/chat:
	go build -o bin/chat cmd/*.go

run: bin/chat
	bin/chat

.PHONY: generate
generate:
	buf mod update
	buf generate
	make copy-swagger

init-local-db:
	docker-compose -f scripts/docker-compose.yml up -d
	goose -dir migrations postgres "postgres://chats-user:chats-password@localhost:5433/chats" up
