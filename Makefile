.PHONY: format,wire,tidy,run

wire:
	cd internal/httpserver && wire
	cd cmd/app && wire
	@ echo "wire generated"

format:
	gofmt -w .
	@echo "go fmt success"

tidy:
	go mod tidy
	@echo "go tidy success"

run:
	go run cmd/main.go
	@echo "go run success"