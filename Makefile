
#============================================================================== #
# HELPERS
#============================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

# Create the new confirm target.
.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

#============================================================================== #
# DEVELOPMENT
#============================================================================== #

## args: show list of args needs by the app
.PHONY: args
args:
	go run . -help

## generate: generate new day directory
.PHONY: generate
generate:
	go run . -generate

## clear: clear all days directories
.PHONY: clear
clear:
	go run . -clear

## lastDay: get last created day
.PHONY: lastDay
lastDay:
	go run . -lastDay
