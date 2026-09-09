$ErrorActionPreference = "Stop"

swag init -g cmd/api/main.go -o docs
npx --yes swagger2openapi docs/swagger.yaml -o docs/openapi.yaml --yaml
