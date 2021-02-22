NETWORKS="$(shell docker network ls)"
VOLUMES="$(shell docker volume ls)"
SUCCESS=[ done "\xE2\x9C\x94" ]

user ?= iamaul
service ?= api

BINARY=engine

test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} main.go

unittest:
	go test -short  ./...

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

all: gateway-network
	@echo [ starting api... ]
	docker-compose up --build traefik api

gateway-network:
ifeq (,$(findstring gateway,$(NETWORKS)))
	@echo [ creating traefik network... ]
	docker network create gateway
	@echo $(SUCCESS)
endif

exec:
	@echo [ executing $(cmd) in $(service) ]
	docker-compose exec -u $(user) $(service) $(cmd)
	@echo $(SUCCESS)

build:
	docker-compose build

run:
	docker-compose up -d --no-build

stop:
	@echo [ stopping all containers... ]
	docker-compose down
	@echo $(SUCCESS)

prune:
	docker system prune -a

.PHONY: test
.PHONY: engine
.PHONY: unittest
.PHONY: clean
.PHONY: lint-prepare
.PHONY: lint
.PHONY: gateway-network
.PHONY: exec
.PHONY: build
.PHONY: run
.PHONY: stop
.PHONY: prune
