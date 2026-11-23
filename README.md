# Trimmed Mean Example

This repository contains a Go application that imports the [`trimmedmean`](https://github.com/gummybeardream/trimmed-mean) package and
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
- Compiled executable instructions
- Whole-application execution-time and memory measurements
- Saved application performance evidence

## Requirements

- Go 1.27.1 or later
- R 4.3.2 or later for R validation
- Git for cloning the repository

The Go application imports:

```text
github.com/gummybeardream/trimmed-mean
```

Go downloads the required package version recorded in `go.mod`.

## Installation

Clone the repository using its `.git` URL:

```bash
git clone https://github.com/gummybeardream/trimmed-mean-example.git
cd trimmed-mean-example
```

Download the recorded Go dependencies:

```bash
go mod download
```

The application uses relative paths to the files in `data`, so commands should be run from the repository root.

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
│   ├── application_performance.txt
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

The Go application reads the shared CSV files and calculates 5% symmetric trimmed means using the external `trimmedmean` package.

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

## Building the executable

Create the local build directory and compile the application:

```bash
mkdir -p bin
go build -o bin/trimmed-mean-example .
```

Run the compiled executable from the repository root:

```bash
./bin/trimmed-mean-example
```

Expected output:

```text
Go trimmed mean validation
Integer trimmed mean (5%): 49.5000000000
Floating-point trimmed mean (5%): 49.7500000000
```

The `bin` directory is ignored by Git because compiled executables are generated artifacts and can be rebuilt from the source code.

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

The results agree to the ten decimal places displayed. This provides evidence that the Go package reproduces the base R trimmed-mean results for these datasets. It does not prove that every possible dataset will produce identical results.

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

## Application performance

The compiled application was measured on November 22, 2026, using:

- Apple M1 processor
- macOS on ARM64
- Go 1.27.1

Five individual executions were measured with macOS `/usr/bin/time -l` for memory use. Because each execution completed faster than the timer's displayed precision, five additional batches of 1,000 executions were used to estimate execution time.

### Results

| Measurement | Result |
|---|---:|
| Average execution time | 2.548 milliseconds |
| Execution-time range | 2.54–2.57 milliseconds |
| Average maximum resident memory | approximately 3.94 MiB |
| Compiled executable size | 2,588,530 bytes (approximately 2.47 MiB) |

The execution-time measurement includes process startup, reading and parsing both CSV files, calculating both trimmed means, and formatting the output. It therefore measures a broader workflow than the package benchmarks, which measure only calls to the trimmed-mean function.

The raw measurements are available in [`results/application_performance.txt`](results/application_performance.txt).

## Logging and error handling

The application uses Go's `log.Fatalf` function when it cannot read a data file or calculate a trimmed mean. The message explains the failed operation and includes the underlying error. The application then exits with a nonzero status.

Normal results are written to standard output. A separate logging framework was not needed because the application performs one small, noninteractive workflow.

## Limitations

- Statistical validation uses two synthetic datasets containing 100 observations each.
- Performance measurements were collected on one Apple M1 computer.
- The application must be run from the repository root because it uses relative paths to the data files.
- Batch timing includes process-startup and file-reading costs.
- The comparison demonstrates agreement for the included datasets but does not prove identical behavior for every possible dataset.

## GenAI Tools

Generative AI tools were used during development to:

- Break the assignment into smaller tasks
- Explain Go syntax and programming concepts
- Review the package design
- Develop and explain unit tests
- Review documentation

The application code and tests were reviewed and executed locally by the repository author.
