VERSION = 0.0.1
BUILD_COMMIT = $(shell git rev-parse HEAD)
BUILD_TIME = $(shell date +%Y-%m-%dT%H:%M:%S%z)
CURRENT_USER = $(USER)
FLAGS = -X main.version=$(VERSION) -X main.commit=$(BUILD_COMMIT) -X main.buildTime=$(BUILD_TIME) -X main.buildUser=$(CURRENT_USER)

# --- cassandra ---
.PHONY: run-c2 run-c3 shell-c2 shell-c3 down
run-c2:
	docker-compose up c2
run-c3:
	docker-compose up c3
shell-c2:
	docker-compose exec c2 /bin/bash
shell-c3:
	docker-compose exec c3 /bin/bash
down:
	docker-compose down
# --- cassandra ---

.PHONY: fmt install dep-update dep-install
fmt:
	gofmt -d -l -w ./cmd ./pkg

install:
	go install -ldflags "$(FLAGS)" ./cmd/gocqlsh
	go install -ldflags "$(FLAGS)" ./cmd/goyourcassandra

dep-update:
	dep ensure -v -update

dep-install:
	dep ensure