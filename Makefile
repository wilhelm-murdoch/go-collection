
BIN=go1.18

build:
	${BIN} build -v ./...

test:
	${BIN} test -race -v ./...

bench:
	${BIN} test -benchmem -count 3 -bench ./...

coverage:
	${BIN} test -v -coverprofile cover.out .