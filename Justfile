profOutput := ".profile.prof"

default:
  just --list

format:
  go fmt ./...
lint:
  docker run -t --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v
build:
  go build
profile day: build
  ./aoc-2024 -day {{day}} -cpuprofile {{ profOutput }}
viewProfile:
  go tool pprof -http 127.0.0.1:8080 aoc-2024 {{ profOutput }}
