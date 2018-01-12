// +build !windows

package cmdline

import "github.com/gentlemanautomaton/cmdline/cmdlineposix"

// Split breaks the given command line into arguments. The arguments are split
// according to posix shell parsing rules.
func Split(cl string) (args []string) {
	return cmdlineposix.Split(cl)
}
