# co : Command Runner

## Features
- Can define short commands
- Can define commands in the current directory, parent directory, CO_HOME, or user home directory
- Can specify the working directory for each command
- Lightweight execution
- Can display command descriptions

## Setting example
```yaml
commands:
  run:
    exec:
      - go run cmd/co/main.go
    working_dir: .
    description: run co
  lint:
    exec:
      - staticcheck ./...
      - go vet ./...
    working_dir: .
    description: style check
  build:
    exec:
      - go build -o dist/co cmd/co/main.go
    working_dir: .
    description: build co
  lm:
    exec:
      - ls -l | more
    description: list and more
```
