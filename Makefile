export GO111MODULE=on
export PATH:=/usr/local/go/bin:$(PATH)
exec_path := /usr/local/bin/
exec_name := gophertelekomcloud


default: test
test: test-unit

fmt:
	@echo Running go fmt
	@go fmt

lint:
	@echo Running go lint
	@golangci-lint run

vet:
	@echo "go vet ."
	@go vet ./...

test-unit:
	@go test ./openstack/... -parallel 4

test-acc:
	@echo "Starting acceptance tests..."
	@go test ./... -race -covermode=atomic -coverprofile=coverage.txt -timeout 20m -v
