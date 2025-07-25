.PHONY: install
install:
	@echo "Installing..."
	go install github.com/bombsimon/wsl/v4/cmd...@v4.4.1
	go install mvdan.cc/gofumpt@v0.6.0
	go install github.com/daixiang0/gci@v0.13.4
	go install github.com/daixiang0/gci@latest

.PHONY: fmt
fmt:
	@echo "Formatting..."
	go mod tidy
	go fmt ./...
	gci write -s standard -s default -s "prefix(github.com/lagzenthakuri/secyourflow)" .
	gofumpt -l -w .
	wsl -fix ./... || true
	cd ui && bun format

.PHONY: lint
lint:
	golangci-lint version
	golangci-lint run  ./...

.PHONY: test
test:
	@echo "Testing..."
	go test -v ./...
	cd ui && bun test

.PHONY: test-coverage
test-coverage:
	@echo "Testing with coverage..."
	go test -coverpkg=./... -coverprofile=coverage.out -count 1 ./...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out

.PHONY: build-ui
build-ui:
	@echo "Building..."
	cd ui && bun install
	cd ui && bun build-only

.PHONY: build
build: build-ui
	@echo "Building..."
	go build -o secyourflow .


.PHONY: build-linux
build-linux: build-ui
	@echo "Building..."
	GOOS=linux GOARCH=amd64 go build -o secyourflow .

.PHONY: docker
docker: build-linux
	@echo "Building Docker image..."
	docker build -f docker/Dockerfile -t secyourflow .

.PHONY: dev
dev:
	@echo "Running..."
	rm -rf secyourflow_data
	go run . admin create admin@secyourflow-soar.com 1234567890
	go run . fake-data
	go run . serve --app-url http://localhost:8090 --flags dev

.PHONY: dev-10000
dev-10000:
	@echo "Running..."
	rm -rf secyourflow_data
	go run . admin create admin@secyourflow-soar.com 1234567890
	go run . fake-data --users 100 --tickets 10000
	go run . serve --app-url http://localhost:8090 --flags dev

.PHONY: default-data
default-data:
	rm -rf secyourflow_data
	go run . default-data

.PHONY: serve-ui
serve-ui:
	cd ui && bun dev --port 3000
