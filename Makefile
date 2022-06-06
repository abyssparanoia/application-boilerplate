# note: call scripts from /scripts

init:
	go get -u github.com/volatiletech/sqlboiler/v4

format:
	$(call format)

build:
	go build -o ./bin/rapid-cli ./cmd/rapid

test:
	go test `go list ./... | grep -v internal/dbmodels`

generate:
	go generate ./internal/...

sqlboiler:
	sqlboiler mysql --config ./db/default/sqlboiler.toml --pkgname defaultdb --wipe --no-hooks --struct-tag-casing camel --output ./internal/dbmodels/defaultdb --templates ${GOPATH}/src/github.com/volatiletech/sqlboiler/templates,${GOPATH}/src/github.com/volatiletech/sqlboiler/templates_test
	$(call format)

swaggergen:
	swagger generate server -f ./swagger/reference/api.v1.yaml -s restapi -m restapi/restmodels -t ./internal
	rm -rf ./internal/cmd
	$(call format)

define format
	go fmt ./... && goimports -w ./ && go mod tidy
endef