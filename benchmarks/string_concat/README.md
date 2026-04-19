# String Concatenation Benchmark

This benchmark analyzes different approaches to string concatenation in Go, focusing on performance, memory usage, and allocation behavior.

---

## 🎯 Goal

Evaluate how string concatenation strategies impact:

* Execution time
* Memory allocation
* Number of allocations

---

## 🧪 Scenarios

### 1. Naive concatenation (`+`)

Repeated string concatenation using the `+` operator inside a loop.

### 2. `strings.Builder` with preallocation

Using `strings.Builder` with `Grow()` to preallocate the buffer size and avoid resizing.

---

## ▶️ How to run

```bash
go test -bench=. -benchmem
```

---

- Iterations: Number of times the benchmark loop was executed
- ns/op: Average time per operation in nanoseconds
- B/op: Average bytes allocated per operation
- allocs/op: Average number of allocations per operation

## 📊 Results

| Benchmark                      | iterations | ns/op | B/op | allocs/op |
| ------------------------------ | ---------: | ----: | ---: | --------: |
| BenchmarkConcatPlus            |     228008 |  4897 | 5664 |        99 |
| BenchmarkConcatBuilderPrealloc |    5483650 |   208 |  112 |         1 |

---

## 🖥️ Environment

Benchmarks were executed on:

* OS: linux
* Architecture: amd64
* CPU: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
* 12 GB RAM DDR4 1400 MHz
* Go version: (fill with `go version`)

---

## 📌 Analysis

### Performance

`strings.Builder` with preallocation is significantly faster (~20x improvement in this run).

### Memory usage

* `+` operator: high memory usage due to repeated allocations
* `Builder`: minimal memory usage due to buffer reuse

### Allocations

* Naive approach: ~1 allocation per iteration
* Builder (preallocated): ~1 allocation total

---

## 🧠 Explanation

Strings in Go are immutable. Each concatenation using `+` creates a new string and copies previous data, leading to:

* repeated allocations
* increasing copy cost
* quadratic time complexity (O(n²))

`strings.Builder` avoids this by:

* using a mutable internal buffer
* minimizing copying
* reducing allocations

Calling `Grow()` ensures the buffer is allocated once, avoiding resizing during writes.

---

## ✅ Key Takeaway

Avoid repeated string concatenation with `+` inside loops.

Prefer:

```go
var b strings.Builder
b.Grow(n)
```

---

## 🔬 Notes

* Results are machine-dependent
* Focus on relative differences, not absolute values
* Preallocation has a significant impact on performance
* This is a microbenchmark, not a real-world workload
