// +build !windows

package cmdline

import "github.com/gentlemanautomaton/cmdline/cmdlineposix"

// Split breaks the given command line into arguments. The arguments are split
// according to posix shell parsing rules.
//
// For an introduction to command line parsing in unix shells, see:
// http://www.grymoire.com/Unix/Quote.html
//
// For details on posix shell parsing requirements see:
// http://pubs.opengroup.org/onlinepubs/9699919799/utilities/V3_chap02.html#tag_18
func Split(cl string) (args []string) {
	return cmdlineposix.Split(cl)
}

// SplitCommand is like Split, but the first argument is returned separately.
func SplitCommand(cl string) (name string, args []string) {
	return cmdlineposix.SplitCommand(cl)
}
