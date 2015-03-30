This directory contains end-to-end test cases for godebug.

**Organization**

For each test `X`, there is a file named `X-in.go` and a file named `X-out.go`. `X-in.go` is a sample program. `X-out.go` is the expected output of running `godebug` on `X-in.go`. There may also be session files for `X`. Session files are named `X-session<name>.txt`.

**Session files**

A session file is a record of a debugging session that steps through and examines the execution of an input file `X-in.go`. The file may begin with a comment, which is any number of blank lines or lines that start with "//".

**endtoend_test**

`endtoend_test.go`, two directories up, checks that `godebug output` produces each `X-out.go`, byte-for-byte, given the `X-in.go` inputs. It also checks that `X-out.go` can run successfully with an exit code of 0. For each test with an `X-session.txt` file, `endtoend_test.go` also runs a debugging session on the output program and checks that it is identical to the session recorded in `X-session.txt`.
