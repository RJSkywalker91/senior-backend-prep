# ------------------------------
# Dev workflow Makefile
# ------------------------------

# ---- Config (override with: make build IMAGE=myapp TAG=dev PORT=3000) ----
IMAGE ?= hello
TAG ?= dev
PORT ?= 8080
COMPOSE ?= docker compose -f build/docker/docker-compose.yaml

# Version metadata for -ldflags (Dockerfile uses these build args)
VERSION ?= 0.1.0
COMMIT  := $(shell git rev-parse --short HEAD 2>/dev/null || echo none)
BUILT_AT := $(shell date -u +%Y-%m-%dT%H:%M:%SZ 2>/dev/null || echo unknown)

# Default target
.DEFAULT_GOAL := help

# ---- Docker image lifecycle ----
.PHONY: build
build: ## Build image with version metadata
	docker build -f build/docker/Dockerfile \
	  --build-arg VERSION=$(VERSION) \
	  --build-arg COMMIT=$(COMMIT) \
	  --build-arg BUILT_AT=$(BUILT_AT) \
	  -t $(IMAGE):$(TAG) .

.PHONY: rebuild
rebuild: ## Rebuild image without cache
	docker build -f build/docker/Dockerfile --no-cache \
	  --build-arg VERSION=$(VERSION) \
	  --build-arg COMMIT=$(COMMIT) \
	  --build-arg BUILT_AT=$(BUILT_AT) \
	  -t $(IMAGE):$(TAG) .

.PHONY: run
run: ## Run image directly (maps $(PORT) and sets APP_MESSAGE)
	docker run --rm -p $(PORT):8080 \
	  -e APP_MESSAGE="Howdy from Make!" \
	  $(IMAGE):$(TAG)

.PHONY: size
size: ## Show image size
	docker image ls $(IMAGE) | head -n 2

.PHONY: rmi
rmi: ## Remove the built image
	-docker rmi -f $(IMAGE):$(TAG)

# ---- Compose workflow ----
.PHONY: up
up: ## Start via docker compose (builds if needed)
	$(COMPOSE) up --build

.PHONY: up-dev
up-dev: ## Start compose with the 'dev' profile (if you added one)
	$(COMPOSE) --profile dev up --build

.PHONY: down
down: ## Stop compose stack
	$(COMPOSE) down

.PHONY: logs
logs: ## Tail logs from compose service

.PHONY: ps
ps: ## List compose services
	$(COMPOSE) ps

.PHONY: restart
restart: ## Restart compose services
	$(COMPOSE) restart

# ---- App developer tasks ----
.PHONY: run-local
run-local: ## Run the app locally without Docker (uses Go)
	APP_MESSAGE="Howdy from Local!" go run ./cmd

.PHONY: build-local
build-local: ## Build the local binary to ./bin/app
	mkdir -p bin
	GOOS=linux CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o bin/app ./cmd

.PHONY: test
test: ## Run Go tests
	go test ./...

.PHONY: fmt
fmt: ## Format source
	gofmt -s -w .

.PHONY: vet
vet: ## Vet code
	go vet ./...

.PHONY: tidy
tidy: ## Go mod tidy
	go mod tidy -v

# ---- Utilities ----
.PHONY: curl
curl: ## Hit the service root and /healthz (requires curl)
	@echo "GET /";    curl -sS http://127.0.0.1:$(PORT)/ || true; echo
	@echo "GET /healthz"; curl -sS http://127.0.0.1:$(PORT)/healthz || true; echo

.PHONY: clean
clean: ## Remove dangling containers/images for this image tag
	-@docker rm -f $$(docker ps -aq --filter "ancestor=$(IMAGE):$(TAG)") 2>/dev/null || true
	-@docker rmi -f $(IMAGE):$(TAG) 2>/dev/null || true
	-@docker system prune -f --volumes 2>/dev/null || true

# ---- Help ----
.PHONY: help
help: ## Show this help
	@printf "\nDev workflow targets:\n\n"
	@awk '/^[a-zA-Z0-9\-\_]+:.*##/ { \
	  split($$0, a, ":"); \
	  gsub(/^ +| +$$/, "", a[1]); \
	  match($$0, /##.*/); \
	  printf "  \033[36m%-12s\033[0m %s\n", a[1], substr($$0, RSTART+3) \
	}' $(MAKEFILE_LIST)
	@printf "\nExamples:\n"
	@printf "  make build           # build Docker image\n"
	@printf "  make run             # run container on port $(PORT)\n"
	@printf "  make up / make down  # compose up/down\n"
	@printf "  make test            # run Go tests\n\n"