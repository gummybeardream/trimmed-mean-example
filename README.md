# Trimmed Mean Example

This repository contains a Go application that imports the
[`trimmedmean`](https://github.com/gummybeardream/trimmed-mean) package and
compares its results with the base R `mean` function.

## Status

Currently implemented:

- Installation of the external `trimmedmean` Go package
- Reproducible datasets containing 100 observations each
- CSV reading and input validation in Go
- Go calculations using 5% symmetric trimming
- R validation using 5% symmetric trimming
- Unit tests for CSV reading and error handling
- Saved Go and R validation results
- Comparison of Go and R results

Still in progress:

- Application performance measurements
- Executable build instructions
- Final documentation and verification

## Project structure

```text
trimmed-mean-example/
├── data/
│   ├── integers.csv
│   └── floats.csv
├── r/
│   ├── generate_data.R
│   └── validate.R
├── results/
│   ├── go_results.txt
│   └── r_results.txt
├── data.go
├── data_test.go
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## Data

The `data` directory contains two shared CSV files:

- `integers.csv`: 100 integers from 0 through 99
- `floats.csv`: 100 floating-point values from 0.25 through 99.25

The observations are stored in an unsorted but reproducible order. The Go application and R validation script read the same files.

Regenerate the datasets from the repository root with:

```bash
Rscript r/generate_data.R
```

## Go validation

The Go application reads the shared CSV files and calculates 5% symmetric
trimmed means using the external `trimmedmean` package.

Run the application from the repository root:

```bash
go run .
```

Results:

| Dataset | Go trimmed mean |
|---|---:|
| 100 integers | 49.5000000000 |
| 100 floating-point values | 49.7500000000 |

The saved output is available in
[`results/go_results.txt`](results/go_results.txt).

## R validation

The R validation script reads the same CSV files that will be used by the Go application. It calculates 5% symmetric trimmed means using the base R `mean` function.

Run it from the repository root:

```bash
Rscript r/validate.R
```

Results:

| Dataset | R trimmed mean |
|---|---:|
| 100 integers | 49.5000000000 |
| 100 floating-point values | 49.7500000000 |

The saved output is available in
[`results/r_results.txt`](results/r_results.txt).

## Go and R comparison

Both implementations read the same CSV files and apply 5% symmetric trimming.

| Dataset | Go | R | Absolute difference |
|---|---:|---:|---:|
| 100 integers | 49.5000000000 | 49.5000000000 | 0.0000000000 |
| 100 floating-point values | 49.7500000000 | 49.7500000000 | 0.0000000000 |

The results agree to the ten decimal places displayed. This provides evidence
that the Go package reproduces the base R trimmed-mean results for these
datasets. It does not prove that every possible dataset will produce identical
results.

## Testing

Run all application tests:

```bash
go test ./...
```

Display every test and subtest:

```bash
go test -v ./...
```

Run Go's static analysis tool:

```bash
go vet ./...
```

The tests cover:

- Valid integer CSV input
- Valid floating-point CSV input
- Invalid integer values
- Invalid floating-point values
- Missing files
- Files containing only a header
- Incorrect column headers
- Empty values
- Unexpected additional columns