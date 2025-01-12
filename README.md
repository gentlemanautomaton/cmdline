cmdline
[![Go Reference](https://pkg.go.dev/badge/github.com/gentlemanautomaton/cmdline.svg)](https://pkg.go.dev/github.com/gentlemanautomaton/cmdline)
====

Package cmdline provides shell argument parsing functions.

Split can be used to extract a slice of arguments from a string that contains
an entire command line. The ability to split shell arguments is particularly
useful when invoking [exec.Command](https://pkg.go.dev/os/exec#Command).

The cmdline package will use typical windows shell parsing rules when
`GOOS=windows`. On non-windows systems posix shell parsing rules will be used.

To use posix or windows shell parsing regardless of `GOOS`, use the
[cmdlineposix](cmdlineposix) and [cmdlinewindows](cmdlinewindows) sub-packages
directly.
