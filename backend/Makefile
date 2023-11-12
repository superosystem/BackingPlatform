BINARY	:= engine

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #
## run: Run golang on development 
.PHONY: run
run:
	@echo 'Backing Platform Starting...'
	go run ./src/main.go

mockery-repository:
	mockery --dir=src/repository --name=$(name) --filename=$(filename).go --output=src/repository/mocks --outpkg=mocks

mockery-service:
	mockery --dir=src/service --name=$(name) --filename=$(filename).go --output=src/service/mocks --outpkg=mocks

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #
## test: Testing project 
.PHONY: test
test:
	@echo 'Backing Platform on Test...'
	go test ./tests/...

## audit: Tidy dependencies and format, vet and test all code
.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...

## vendor: Tidy and vendor dependencies
.PHONY: vendor
vendor:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	go mod verify
	@echo 'Vendoring dependencies...'
	go mod vendor

# ==================================================================================== #
# MIGRATION
# ==================================================================================== #
## migrate-up: Migrate database versioning
.PHONY: migrate-up
migrate-up:
	@echo 'Backing Platform migration up...'
	go run ./src/migrater/main.go -migrate=up

## migrate-down: Migration clean database versioning
.PHONY: migrate-down
migrate-down:
	@echo 'Backing Platform migration down...'
	go run ./src/migrater/main.go -migrate=down

# ==================================================================================== #
# PACKAGING
# ==================================================================================== #
## engine: Build project
.PHONY: engine
engine:
	@echo 'Build Backing Platform...'
	go build -o ${BINARY} src/*.go

.PHONY: win-engine
win-engine:
	@echo 'Build Backing Platform...'
	go build -o ${BINARY}.exe src/main.go

## engine: Build migrater for Migrations
.PHONY: migrater
migrater:
	@echo 'Build Backing Platform Migrater...'
	go build -o migrater src/migrater/main.go

.PHONY: win-migrater
win-migrater:
	@echo 'Build Backing Platform Migrater...'
	go build -o migrater.exe src/migrater/main.go

clean:
	@echo 'CLEAN BINARY...'
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# ==================================================================================== #
# CONTAINER
# ==================================================================================== #
## docker-build: Build container for development
.PHONY: docker-dev-up
docker-dev-up:
	@echo 'Build Backing Platform Container up...'
	docker compose -f docker-compose-dev.yml up

docker-dev-down:
	@echo 'Build Backing Platform Container down...'
	docker compose -f docker-compose-dev.yml down

## docker-build: Build project to image
.PHONY: docker-build
docker-build:
	@echo 'Build Backing Platform Image...'
	docker build -f ./docker/Dockerfile .

docker-up:
	@echo 'Build Backing Platform Container PROD...'
	docker-compose up --build -d

docker-down:
	@echo 'Backing Platform Container Stop...'
	docker-compose down