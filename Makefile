.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## run: builds the binary and test it with nvim
.PHONY: run
run:
	go build -o bin/main && nvim ./README.md

## log: watch the changes on a log file
.PHONY: log
log:
	tail -f log.txt
