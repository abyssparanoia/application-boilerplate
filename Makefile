# note: call scripts from /scripts

install:
	$(eval BIN:=$(abspath bin))
	mkdir -p ./bin
	go mod download golang.org/x/tools
	GOBIN="$(BIN)" go install github.com/volatiletech/sqlboiler/v4@v4.11.0
	GOBIN="$(BIN)" go install github.com/pressly/goose/v3/cmd/goose@v3.5.3
	GOBIN="$(BIN)" go install github.com/golang/mock/mockgen@v1.6.0
	GOBIN="$(BIN)" go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest

format:
	$(call format)

build:
	go build -o ./bin/rapid-cli ./cmd/rapid

test:
	go test `go list ./... | grep -v internal/dbmodels`

generate:
	go generate ./internal/...

sqlboiler:
	./bin/sqlboiler mysql \
		--config ./db/default/sqlboiler.toml \
		--pkgname defaultdb \
		--wipe \
		--no-hooks \
		--no-tests \
		--struct-tag-casing camel \
		--output ./internal/dbmodels/defaultdb \
		--templates ${GOPATH}/pkg/mod/github.com/volatiletech/sqlboiler/v4@v4.11.0/templates/main
	$(call format)

swaggergen:
	./bin/swagger generate server -f ./swagger/reference/api.v1.yaml -s restapi -m restapi/restmodels -t ./internal
	rm -rf ./internal/cmd
	$(call format)

GOOSE_FILE_NAME=something
goose-create:
	./bin/goose create ${GOOSE_FILE_NAME} sql

goose-up:
	./bin/goose --dir db/default/migrations mysql "$(DB_USER):$(DB_PASSWORD)@$(DB_HOST)/$(DB_DATABASE)?parseTime=true" up

define format
	go fmt ./... && goimports -w ./ && go mod tidy
endef