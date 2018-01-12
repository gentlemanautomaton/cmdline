// +build windows

package cmdline

import "github.com/gentlemanautomaton/cmdline/cmdlinewindows"

// Split breaks the given command line into arguments. The arguments are split
// according to command prompt parsing rules.
//
// For details on command line parsing in Windows see:
// https://docs.microsoft.com/en-us/cpp/c-language/parsing-c-command-line-arguments
func Split(cl string) (args []string) {
	return cmdlinewindows.Split(cl)
}
