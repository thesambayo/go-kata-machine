# Go Kata Machine

A Golang implementation of Data Structures and Algorithms from the [Frontend Masters JavaScript Algorithms Course](https://frontendmasters.com/courses/algorithms).

## Overview
This project provides a structured way to practice implementing various data structures and algorithms in Go. It automatically generates daily kata exercises and includes test suites to verify your implementations.

## Prerequisites
- Go
- Git

## Installation
No extra library to install, as long as you've got Golang and a terminal/IDE, you are good to go.
```

## Usage

### Generating Kata Exercises
Generate a new day of kata exercises:
```bash
go run . -generate
```
This will create a new directory (e.g., `day1`, `day2`, etc.) with skeleton implementations and test files.

### Clearing Generated Katas
To remove all generated kata directories:
```bash
go run . -clear
```

### Checking Latest Generated Day
To see the most recently generated kata day:
```bash
go run . -lastDay
```

## Testing Your Implementations


### Verbose Testing
For detailed test output:
```bash
go test ./tests -v -day=dayN
```

### Testing Specific Algorithm
To test a specific algorithm with verbose output:
```bash
go test ./tests -v -day=dayN -run ^TestAlgorithmName$
# Example:
go test ./tests -v -day=day4 -run ^TestBubbleSort$
```

## Project Structure
```
go-kata-machine/
├── helpers/         # Helper functions and DSA details
├── dayN/           # Generated kata directories
│   ├── algorithm1/
│   ├── algorithm2/
│   └── ...
├── tests/          # Test suites
└── README.md
```

## Upcoming Features
- Makefile implementation with the following commands:
  - `make generate` - Generate new kata day
  - `make delete -day dayN` - Delete specific day
  - `make tests -day=dayN -v -file` - Run tests with options
  - `make tests -day=dayN -v -file TestAlgorithm` - Test specific algorithm

## Contributing
Feel free to contribute by:
1. Forking the repository
2. Creating a feature branch
3. Committing your changes
4. Creating a pull request

## Acknowledgments
- Frontend Masters for the original JavaScript algorithms course
- [Add any other acknowledgments]

## Note
This project is under active development. New algorithms and features are being added regularly.
