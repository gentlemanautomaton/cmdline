// +build windows

package cmdline

import "github.com/gentlemanautomaton/cmdline/cmdlinewindows"

// Split breaks the given command line into arguments. The split is performed
// according to the standard windows shell parsing rules as implemented by the
// Microsoft C compiler.
//
// For details on command line parsing in Windows see:
// https://docs.microsoft.com/en-us/cpp/c-language/parsing-c-command-line-arguments
func Split(cl string) (args []string) {
	return cmdlinewindows.Split(cl)
}

// SplitCommand is like Split, but the first argument is returned separately.
func SplitCommand(cl string) (name string, args []string) {
	return cmdlinewindows.SplitCommand(cl)
}
