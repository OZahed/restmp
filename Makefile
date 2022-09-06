BUILD_BRANCH=$$(git rev-parse --abbrev-ref HEAD)
BUILD_INFO_PKG=github.com/OZahed/restmp/internal/configs
BUILD_TIME =$$(date '+FT%T')
BUILD_TAG=$$(git describe --abbrev=0 )


# Did not use ldflags because in build time if we forgot using ldflags we get no errors
export APP_NAME=restmp
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export LDFLAGS="-s -w -X $(BUILD_INFO_PKG).appname=$(APP_NAME) -X $(BUILD_INFO_PKG).build=$(BUILD_TIME) -X $(BUILD_INFO_PKG).commit=$(git rev-parse HEAD | cut -c 1-8) -X $(BUILD_INFO_PKG).branch=$(BUILD_BRANCH) -X $(BUILD_INFO_PKG).tag=$(BUILD_TAG)"

build: swagger
	CGO_ENABLED=1 go build -ldflags $(LDFLAGS)  .


build-static:
	CGO_ENABLED=1 go build -v -o $(APP) -a -installsuffix cgo -ldflags $(LDFLAGS) .

build-docker:
	docker build $(APP_NAME):$(BUILD_TAG) .

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
