# ca : Command Alias

## Features
- [ ] Command aliases can be defined and executed in short commands
- [ ] Command aliases in upper directory definitions can also be executed.
- [ ] Working directory can be set for command execution
- [ ] Runs lightly
- [ ] Displays predefined commands

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
