BUILD_BRANCH=$$(git rev-parse --abbrev-ref HEAD)
BUILD_INFO_PKG=github.com/OZahed/restmp/configs
BUILD_TIME =$$(date '+FT%T')

APP_NAME=restmp # override it

# Did not use ldflags because in build time if we forgot using ldflags we get no errors
export APP_NAME=$(APP_NAME)
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export LDFLAGS="-X $(BUILD_INFO_PKG).BuildTime=$(BUILD_TIME) -X $(BUILD_INFO_PKG).GitHash=$(git rev-parse HEAD | cut -c 1-8) -X $(BUILD_INFO_PKG).GitBranch=$(BUILD_BRANCH)"

build: swagger
	CGO_ENABLED=1 go build -ldflags $(LDFLAGS)  .


build-static:
	CGO_ENABLED=1 go build -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

build-docker:
	docker build restmp my

install:
	CGO_ENABLED=1 go install -ldflags $(LDFLAGS)

swagger: 
	which swagger  || go get -u github.com/go-swagger/go-swagger/cmd/swagger
	GO111MODULE=on swagger generate spec -o ./docs/swagger.yaml --scan-models

generate: 
	which easyjson || go get -u github.com/mailru/easyjson/...
	which mockgen  || go get -u github.com/golang/mock/mockgen
	which stringer || go get -u golang.org/x/tools/cmd/stringer
	go generate $(ROOT)/...
