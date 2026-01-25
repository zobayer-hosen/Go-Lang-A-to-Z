# 🐹 Go-Lang-A-to-Z — Complete Learning Roadmap

A hands-on Go journey from basics to advanced topics: master syntax, tooling, concurrency, and production-grade services.

---

## 📋 Table of Contents

- [Overview](#overview)
- [Why Go](#why-go)
- [Prerequisites](#prerequisites)
- [Installation & Setup](#installation--setup)
- [Quick Start Guide](#quick-start-guide)
- [Complete Go Roadmap](#complete-go-roadmap)
- [Repository Structure](#repository-structure)
- [Practical Examples](#practical-examples)
- [Hands-on Practice](#hands-on-practice)
- [Best Practices](#best-practices)
- [Resources](#resources)
- [Contributing](#contributing)

---

## Overview

This repo is your step-by-step Go workshop: short files per day, concise examples, and practical exercises to build confidence for real-world backend and systems work.

## Why Go

- Simple, readable syntax with batteries-included tooling.
- Fast compilation and strong concurrency primitives for scalable services.
- Great for CLIs, APIs, cloud-native workloads, and distributed systems.
- Excellent standard library plus first-class support for testing.

## Prerequisites

- Basic programming experience (variables, loops, functions).
- Git and a terminal shell installed.
- Go 1.21+ (see setup below).

## Installation & Setup

1. Install Go: https://go.dev/dl/
2. Verify: `go version`
3. Initialize modules (already present here): `go mod tidy`
4. Editor: VS Code with the Go extension for linting, formatting, and test integration.

## Quick Start Guide

1. Clone the repo and open it in VS Code.
2. Run a basics example: `go run 002_DAY002/variable_content.go`
3. Try loops and defer: `go run 005_DAY005/loop/loops2.go`
4. Explore slices and structs: `go run DAY-8/slice.go`
5. Add your own file under a day folder and run it with `go run path/to/file.go`.

## Complete Go Roadmap

- **Foundations**: variables, constants, scopes, control flow, functions. Build comfort with Go tooling (`go run`, `go fmt`, `go test`).
- **Data & Types**: arrays, slices, maps, structs, pointers, methods, interfaces; type conversions and zero values.
- **Error Handling**: idiomatic `error` values, `panic`/`recover`, defer patterns, and guard clauses.
- **Concurrency**: goroutines, channels, `select`, sync primitives (`WaitGroup`, `Mutex`), and cancellation with `context`.
- **I/O & Networking**: files, JSON, HTTP clients/servers; design small REST/gRPC services.
- **Testing & Benchmarking**: table-driven tests, subtests, benchmarks, code coverage.
- **Deployment & Tooling**: modules, versioning, `go mod tidy`, lint/format, containerizing small services.

## Repository Structure

- 001_DAY001/ — day-one notes and warm-ups.
- 002_DAY002/ — variables, scope, and type basics.
- 003_DAY003/ — type conversion exercises.
- 004_DAY004/ — constants and functions.
- 005_DAY005/ — functions, anonymous functions, defer, and loops (see loop/ subfolder).
- 006_DAY006/ — scopes, `goto`, and error-with-goto examples.
- 007_DAY007/ — pointers and dereferencing.
- DAY-8/ — arrays, slices, and structs.
- go.mod — module metadata for dependency management.

## Practical Examples

- Variables and shadowing: [002_DAY002/gobal_local_variable.go](002_DAY002/gobal_local_variable.go)
- Type conversion walkthrough: [003_DAY003/typeConversion.go](003_DAY003/typeConversion.go)
- Functions and named returns: [005_DAY005/nacked_return.go](005_DAY005/nacked_return.go)
- Defer patterns: [005_DAY005/defer/defer01.go](005_DAY005/defer/defer01.go)
- Loops practice: [005_DAY005/loop/loops2.go](005_DAY005/loop/loops2.go)
- Slices and arrays: [DAY-8/slice.go](DAY-8/slice.go), [DAY-8/array.go](DAY-8/array.go)
- Struct basics: [DAY-8/STRUCT.go](DAY-8/STRUCT.go)

## Hands-on Practice

- Rewrite a C-style `for` loop using Go's range form over slices and maps.
- Add a new example that uses a pointer receiver method and run it with `go run`.
- Extend the defer example to include error handling and `recover`.
- Write a table-driven test for any function (use `go test ./...`).

## Best Practices

- Favor clarity: short functions, explicit errors, and clear naming.
- Keep code formatted: `go fmt ./...` before commits.
- Handle errors early; return concrete types from packages and keep interfaces small.
- Use `context` for cancellation in long-running or I/O-heavy code.

## Resources

- Go docs: https://go.dev/doc/
- Effective Go: https://go.dev/doc/effective_go
- Tour of Go: https://go.dev/tour
- Standard library reference: https://pkg.go.dev/std

## Contributing

1. Fork and create a feature branch.
2. Add or improve examples with clear comments where needed.
3. Run `go fmt ./...` and `go test ./...`.
4. Open a PR describing the change and the learning goal it covers.

Happy coding and learning! 🚀
