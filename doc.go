// Package cmdline provides shell argument parsing functions.
//
// Split can be used to extract a slice of arguments from a string that contains
// an entire command line. The ability to split shell arguments is particularly
// useful when invoking exec.Command.
//
// The cmdline package will use typical windows shell parsing rules when
// GOOS=windows. On non-windows systems posix shell parsing rules will be used.
//
// To use posix or windows shell parsing regardless of GOOS, use the
// cmdlineposix and cmdlinewindows sub-packages directly.
package cmdline
