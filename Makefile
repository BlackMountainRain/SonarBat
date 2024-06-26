GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest
	go install github.com/google/wire/cmd/wire@latest

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-http_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       --openapi_out=fq_schema_naming=true,default_response=false:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

run: $(filter-out $@,$(MAKECMDGOALS))
	@:

auth:
	go run -ldflags "-X main.Version=$(VERSION) -X main.Name=auth" ./cmd/auth/

resource:
	go run -ldflags "-X main.Version=$(VERSION) -X main.Name=resource" ./cmd/resource/

task:
	go run -ldflags "-X main.Version=$(VERSION) -X main.Name=task" ./cmd/task/

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: validate
# generate validate proto
validate:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   --validate_out=paths=source_relative,lang=go:. \
		   $(API_PROTO_FILES)


.PHONY: errors
# generate errors
errors:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:. \
           --go-errors_out=paths=source_relative:. \
           $(API_PROTO_FILES)

.PHONY: all
# generate all
all:
	make api;
	make config;
	make generate;
	make validate;
	make errors;
	go mod tidy;

.PHONY: test
# run all tests
test:
	go test -v ./...

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

# start the dev environment
dev:
	docker compose -f ./build/dev/docker-compose.yml --project-directory . up -d

# start the dev environment
dev-f:
	docker compose -f ./build/dev/docker-compose.yml --project-directory . up -d --force-recreate

dev-del:
	docker rmi sonarbat-backend:latest

# tear down the dev environment
down:
	docker compose -f ./build/dev/docker-compose.yml --project-directory . down

# dive into the dev container by service name
in:
	@docker compose -f ./build/dev/docker-compose.yml --project-directory . exec $(filter-out $@,$(MAKECMDGOALS))

%:
	@: