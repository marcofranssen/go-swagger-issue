reproduce: generate
	@git diff

docs/swagger.json:
	@mkdir -p docs
	@touch docs/swagger.json

generate: docs/swagger.json
	@echo Generate code
	@go generate ./...
	
install: install-tools

install-tools:
	@echo Installing tools from tools.go
	@cat tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %