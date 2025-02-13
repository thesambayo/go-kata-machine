
<!-- TO DO -->
## Create a make file to make running the program even easier
with the following commands:
- `make generate` or `go run . -generate`
<!-- -generate a boolean flag -->
- `make delete -day day1` or `go run . -clear -day day3`
<!--
  -clear a boolean flag,
  -day an optional flag to clear only a given day,
  if -day is absent, all days will be cleared
  if present, either dayX is present or not,
    program will try to clear only that day
-->
- `make tests -day=dayX -v -file` or `go test ./tests -v -day=dayX`
<!--
  -day dayX, a flag for day to run tests against,
  otherwise day1 is tested against.
  -v a optional boolean flag for verbose
-->
- To test against only one file:
 You will run `make tests -day=dayX -v -file Test{Name of File in day}` or
 ` go test ./tests -v -day day4 -run ^{File}$`
 eg: `go test ./tests -v -day day4 -run ^TestBubbleSort$`
