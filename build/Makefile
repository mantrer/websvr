APP?=websvr
PORT?=58080
RELEASE?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
GOARCH?=amd64

GOFILES?=../cmd/$(APP)/*.go

clean:
	@echo "  >  Deleting binary..."
	rm -f ${APP}

build: clean
	@echo "  >  Building binary..."
#	go  build -o $(APP) $(GOFILES)
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build --tags netgo \
		-ldflags "-s -w -X 127.0.0.1/version.Release=${RELEASE} \
		-X 127.0.0.1/version.Commit=${COMMIT} -X 127.0.0.1/version.BuildTime=${BUILD_TIME}" \
		-o ${APP} ${GOFILES}
container: build
	docker build -t $(APP):$(RELEASE) .
	
run: container
	docker-compose up -d

test:
	go test -v -race ${GOFILES}
