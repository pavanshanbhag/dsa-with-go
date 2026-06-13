# Advanced Search

Binary search variants and related search algorithms.

## Files

- `binary_search_variants.go` – implementations
- `binary_search_variants_test.go` – tests and benchmarks

## Contents

- **ClassicBinarySearch** – standard binary search on sorted slice
- **LowerBound** / **UpperBound** – first index ≥ target, first index > target
- **Rotated array search** – search in rotated sorted array
- **ExponentialSearch** – unbounded / infinite sequences
- **InterpolationSearch** – uniformly distributed data
- **TernarySearch** – unimodal function minimization

## How to verify (from repo root)

```bash
go test ./04-advanced-topics/search/ -v
go test ./04-advanced-topics/search/ -bench=. -benchmem
```

## Example

```bash
go run examples/advanced-topics/main.go
```

(Demo includes search and number-theory examples.)
