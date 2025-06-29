.PHONY: format,wire,tidy,swag,run

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

# 生成 Swagger 文档
swag:
	swag init \
		-g cmd/main.go \
		--output cmd/docs/ \
		--parseInternal \
		--parseDependency

# 启动服务（依赖 swag 目标）
run: swag
	go run cmd/main.go
	@echo "go run success"


# 清除swagger
clean:
	rm -rf cmd/docs/
	@echo "clean swagger success"

# git add
add: tidy
	git add .