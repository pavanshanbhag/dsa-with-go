# Numerical Optimization

One-dimensional optimization and root-finding methods.

## Files

- `numerical_methods.go` – implementations (package `optimization`)
- `numerical_methods_test.go` – tests

## Contents

- **GoldenSectionSearch** – minimize unimodal function (golden ratio)
- **TernarySearch** – discrete or continuous unimodal minimization
- **BisectionMethod** – root finding (continuous function, sign change)
- **NewtonRaphson** – root finding using derivative

## How to verify (from repo root)

```bash
go test ./04-advanced-topics/optimization/ -v
go test ./04-advanced-topics/optimization/ -bench=. -benchmem
```

## Example

```bash
go run examples/advanced-topics/main.go
```

(Demo includes optimization examples.)
