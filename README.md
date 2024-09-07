# do : Command Runner

## Features
- Can define short commands
- Can define commands in the current directory, parent directory, DO_HOME, or user home directory
- Can specify the working directory for each command
- Lightweight execution
- Can display command descriptions

## Setting example
```yaml
commands:
  run:
    exec:
      - go run cmd/do/main.go
    working_dir: .
    description: run do
  lint:
    exec:
      - staticcheck ./...
      - go vet ./...
    working_dir: .
    description: style check
  build:
    exec:
      - go build -o dist/do cmd/do/main.go
    working_dir: .
    description: build ca
  lm:
    exec:
      - ls -l | more
    description: list and more
```
