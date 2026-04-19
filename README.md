# go-benchmark-lab

A collection of benchmark experiments in Go focused on understanding performance characteristics, trade-offs, and runtime behavior.

## 🎯 Purpose

This repository is intended as a practical lab to explore how different approaches in Go impact:

* CPU performance
* Memory allocation
* Concurrency behavior
* Garbage collection

Each benchmark is designed to be simple, reproducible, and focused on a specific concept.

## 📁 Structure

```
/benchmarks
  /<topic>
    benchmark_test.go
    README.md
```

Examples of topics:

* slices vs arrays
* sync vs channels
* allocation patterns
* string operations
* pointer vs value semantics
* etc.

## 🚀 Running benchmarks

```bash
go test -bench=. -benchmem ./...
```

## 📌 Notes

* All benchmarks follow Go's standard testing package (`testing.B`)
* Results may vary depending on machine and Go version
* Focus is on learning and experimentation, not absolute numbers
