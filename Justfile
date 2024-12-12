profOutput := ".profile.prof"

default:
  just --list

format:
  go fmt ./...
lint:
  docker run -t --rm -v $(pwd):/app:ro -w /app golangci/golangci-lint:v1.61.0 golangci-lint run -v
sec:
  docker run --rm -it -w /aoc-2024/ -v $(pwd):/aoc-2024:ro securego/gosec /aoc-2024/
build:
  go build
profile day: build
  ./aoc-2024 -day {{day}} -cpuprofile {{ profOutput }}
viewProfile:
  go tool pprof -http 127.0.0.1:8080 aoc-2024 {{ profOutput }}
