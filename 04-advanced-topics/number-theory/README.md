# Number Theory

Prime generation, primality testing, modular arithmetic, GCD, factorization, and related algorithms.

## Files

- `primes.go` – implementations (package `number_theory`)
- `primes_test.go` – tests and benchmarks

## Contents

- **SieveOfEratosthenes** – list primes up to n
- **IsPrime** (trial division, Miller-Rabin)
- **Modular arithmetic** – exponentiation, inverse
- **GCD** – Euclidean, extended Euclidean, binary GCD
- **Factorization** – trial division, Pollard's rho
- **Chinese Remainder Theorem**

## How to verify (from repo root)

```bash
go test ./04-advanced-topics/number-theory/ -v
go test ./04-advanced-topics/number-theory/ -bench=. -benchmem
```

## Example

```bash
go run examples/advanced-topics/main.go
```
