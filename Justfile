format:
  go fmt ./...
lint:
  docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v