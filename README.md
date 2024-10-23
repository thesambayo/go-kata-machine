# Golang's DSA version of [Frontend masters JavaScript algorithms course](https://frontendmasters.com/courses/algorithms)

## WARNING
I have just started to add algorithms, as I have not completed the course.

## How to get started
Clone the repo, of course, you should have golang installed on your computer.

Create a day of katas, this will use the files in `helpers/*`.
You can check the `helpers/dsa-details.go` file to see list of ready algorithms that will have its respective bare files and tests generated.

```go
go run . -generate
// or
go run . -generate=true
```

This will progressively create folders named
```shell
day1/*
day2/*
day3/*
day4/*
```

## To clear all generated days of katas, run:
```go
go run . -clear
// or
go run . -clear=true
```

## To see most recent generated day run:
```go
go run . -day
// or
go run . -day=true
```

## How to test
### You run go test `path to day/package` like so to test a single test:
```shell
go test ./day1/bubblesort
```


### Or run all tests for a day workspace:
```shell
# open terminal
cd day* # into desired day workspace

# then run
go test ./...
```
