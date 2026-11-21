# Trimmed Mean Example

This repository contains a Go application that imports the
[`trimmedmean`](https://github.com/gummybeardream/trimmed-mean) package and
compares its results with the base R `mean` function.

## Status

Currently implemented:

- Installation of the external `trimmedmean` Go package
- Initial integer and floating-point examples
- Reproducible datasets containing 100 observations each
- R validation using 5% symmetric trimming
- Saved R validation results

Still in progress:

- Reading the shared CSV files in Go
- Calculating 5% trimmed means
- Application tests and performance measurements
- Executable build instructions

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