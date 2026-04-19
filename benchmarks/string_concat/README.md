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

### Metrics

* **iterations**: number of benchmark iterations
* **ns/op**: average time per operation (nanoseconds)
* **B/op**: bytes allocated per operation
* **allocs/op**: allocations per operation

---

## 📊 Results

### Run 1

| Benchmark                      | iterations | ns/op | B/op | allocs/op |
| ------------------------------ | ---------: | ----: | ---: | --------: |
| BenchmarkConcatPlus            |     218046 |  4843 | 5664 |        99 |
| BenchmarkConcatBuilderPrealloc |    5417446 | 195.1 |  112 |         1 |

---

### Run 2

| Benchmark                      | iterations | ns/op | B/op | allocs/op |
| ------------------------------ | ---------: | ----: | ---: | --------: |
| BenchmarkConcatPlus            |     258368 |  4408 | 5664 |        99 |
| BenchmarkConcatBuilderPrealloc |    6048504 | 200.9 |  112 |         1 |

---

### Run 3

| Benchmark                      | iterations | ns/op | B/op | allocs/op |
| ------------------------------ | ---------: | ----: | ---: | --------: |
| BenchmarkConcatPlus            |     232026 |  4929 | 5664 |        99 |
| BenchmarkConcatBuilderPrealloc |    5623080 | 210.2 |  112 |         1 |

---

## 📊 Average Results (3 runs)

| Benchmark                      | avg iterations | avg ns/op | avg B/op | avg allocs/op |
| ------------------------------ | -------------: | --------: | -------: | ------------: |
| BenchmarkConcatPlus            |         236813 |    4726.7 |     5664 |            99 |
| BenchmarkConcatBuilderPrealloc |        5696343 |     202.1 |      112 |             1 |

---

## 🖥️ Environment

Benchmarks were executed on:

* OS: linux
* Architecture: amd64
* CPU: Intel(R) Core(TM) i5-8250U CPU @ 1.60GHz
* 12 GB RAM DDR4 1400 MHz

---

## 📌 Analysis

### Performance

`strings.Builder` with preallocation is ~23x faster on average in this benchmark.

### Memory usage

* `+` operator: high memory usage due to repeated allocations and copies
* `Builder`: minimal memory usage due to buffer reuse

### Allocations

* Naive approach: **~99 allocations per operation** (not 1 per iteration)
* Builder (preallocated): **1 allocation total**

---

## 🧠 Explanation

Strings in Go are immutable. Each concatenation using `+`:

* allocates a new string
* copies previous content
* increases total cost progressively

This leads to:

* multiple allocations (~N)
* growing copy cost
* overall **O(n²)** behavior in loops

`strings.Builder` avoids this by:

* using a mutable internal buffer
* reducing copies
* performing a single allocation when `Grow()` is used

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
* Preallocation has a major impact
* This is a microbenchmark, not a real-world workload
