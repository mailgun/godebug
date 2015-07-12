# Comments starting with #: below are remake GNU Makefile comments. See
# https://github.com/rocky/remake/wiki/Rake-tasks-for-gnu-make

.PHONY: all eval test check gentests install

#: Same as repl
all: repl

#: The REPL front-end to the evaluator
repl: lib
	go build -o go-repl demo/repl.go

#: The evaluator library
lib:
	go build

#: Same as "check"
test: check

#: Install code
install:
	go install

#: Automated generate of the massive number of type checking tests used
gentests:
	make -C testgen


#: Run all tests (quick and interpreter)
check:
	go test -i && go test
