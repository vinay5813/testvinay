# testvinay

A simple Go utility project by [vinay5813](https://github.com/vinay5813).

## Functions

### `isEven(n int) bool`
Returns `true` if `n` is even, `false` if odd.

### `add(a, b int) int`
Returns the sum of two integers.

### `checkString(s string) string`
Classifies a string as:
- `"numeric"` — all characters are digits
- `"alphanumeric"` — all characters are letters or digits
- `"neither"` — contains special characters
- `"empty"` — string has no characters

## Running

```bash
go run main.go
```

## Testing

```bash
go test -v ./...
```

## CI

GitHub Actions runs lint, build, and tests on every push and pull request to `main`.
