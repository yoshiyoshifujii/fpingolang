# Repository Guidelines

## Project Structure & Module Organization

- Root contains `go.mod` (module definition) and `README.md`.
- Go source files should live in the repository root or under package folders such as `pkg/` or `internal/` as the project grows.
- Tests should be colocated with code as `*_test.go` files.

## Build, Test, and Development Commands

- `go build ./...` — compile all packages in the module.
- `go test ./...` — run all tests across packages.
- `go test -v ./...` — verbose test output for debugging.
- `go fmt ./...` — format all Go files with `gofmt`.

## Coding Style & Naming Conventions

- Indentation: tabs (standard Go formatting via `gofmt`).
- Exported identifiers: `CamelCase` (e.g., `NewParser`).
- Unexported identifiers: `camelCase` (e.g., `parseLine`).
- File names: lowercase with underscores when needed (e.g., `parser_utils.go`).
- Keep packages small and focused; avoid circular imports.

## Testing Guidelines

- Use Go’s built-in `testing` package.
- Test file pattern: `*_test.go`; test functions: `TestXxx(t *testing.T)`.
- Prefer table-driven tests for multiple cases.
- Run `go test ./...` before opening a PR; no coverage target is defined yet.

## Commit & Pull Request Guidelines

- No commit history is available yet; follow Conventional Commits where possible (e.g., `feat: add parser`).
- PRs should include a short description of the change, testing notes, and any relevant issues.
- If behavior changes, include before/after details or examples.

## Configuration Notes

- Go version is defined in `go.mod`; keep it in sync with local tooling.
