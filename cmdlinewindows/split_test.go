package cmdlinewindows_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gentlemanautomaton/cmdline/cmdlinewindows"
)

var splitCommandTests = []struct {
	Name         string
	CL           string
	ExpectedName string
	ExpectedArgs []string
}{
	{"empty", ``, ``, nil},
	{"args-0", `a`, `a`, nil},
	{"args-1", `a b`, `a`, []string{`b`}},
	{"args-2", `a b c`, `a`, []string{`b`, `c`}},
	{"args-3", `a b c d`, `a`, []string{`b`, `c`, `d`}},
	{"quoted-program", `"a b c" d`, `a b c`, []string{`d`}},
	{"quoted-program-args", `"a b c" "d e" "f"`, `a b c`, []string{`d e`, `f`}},
}

func TestSplitCommand(t *testing.T) {
	for _, tc := range splitCommandTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			name, args := cmdlinewindows.SplitCommand(tc.CL)
			if name != tc.ExpectedName || !argsEqual(args, tc.ExpectedArgs) {
				e := printArgs(tc.ExpectedArgs)
				a := printArgs(args)
				t.Errorf("input: %s   expected: %s %s   actual: %s %s", tc.CL, tc.ExpectedName, e, name, a)
			}
		})
	}
}

var splitLiteralCommandTests = []struct {
	Name         string
	CL           string
	ExpectedName string
	ExpectedArgs []string
}{
	{"empty", ``, ``, nil},
	{"args-0", `a`, `a`, nil},
	{"args-1", `a b`, `a`, []string{`b`}},
	{"args-2", `a b c`, `a`, []string{`b`, `c`}},
	{"args-3", `a b c d`, `a`, []string{`b`, `c`, `d`}},
	{"quoted-program", `"a b c" d`, `"a b c"`, []string{`d`}},
	{"quoted-program-args", `"a b c" "d e" "f"`, `"a b c"`, []string{`"d e"`, `"f"`}},
}

func TestSplitLiteralCommand(t *testing.T) {
	for _, tc := range splitLiteralCommandTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			name, args := cmdlinewindows.SplitLiteralCommand(tc.CL)
			if name != tc.ExpectedName || !argsEqual(args, tc.ExpectedArgs) {
				e := printArgs(tc.ExpectedArgs)
				a := printArgs(args)
				t.Errorf("input: %s   expected: %s %s   actual: %s %s", tc.CL, tc.ExpectedName, e, name, a)
			}
		})
	}
}

var splitTests = []struct {
	Name     string
	CL       string
	Expected []string
}{
	{"quoted-space", `"a b c" d e`, []string{`a b c`, `d`, `e`}},
	{"escaped-quote", `"ab\"c" "\\" d`, []string{`ab"c`, `\`, `d`}},
	{"unescaped-backslash", `a\\\b d"e f"g h`, []string{`a\\\b`, `de fg`, `h`}},
	{"escaped-backslash-odd", `a\\\"b c d`, []string{`a\"b`, `c`, `d`}},
	{"escaped-backslash-even", `a\\\\"b c" d e`, []string{`a\\b c`, `d`, `e`}},
	{"empty-arg-1", `a "" c`, []string{`a`, ``, `c`}},
	{"empty-arg-2", `a "" c ""`, []string{`a`, ``, `c`, ``}},
	{"empty-arg-3", `a b "" ""`, []string{`a`, `b`, ``, ``}},
	{"uneven-quote-1", `a"bc`, []string{`abc`}},
	{"uneven-quote-2", `a"b c d`, []string{`ab c d`}},
	{"uneven-quote-3", `ab c d"`, []string{`ab`, `c`, `d`}},
	{"uneven-quote-4", `a "b c d`, []string{`a`, `b c d`}},
	{"command-echo", `cmd /C echo test`, []string{`cmd`, `/C`, `echo`, `test`}},
}

func TestSplit(t *testing.T) {
	for _, tc := range splitTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			args := cmdlinewindows.Split(tc.CL)
			if !argsEqual(args, tc.Expected) {
				e := printArgs(tc.Expected)
				a := printArgs(args)
				t.Errorf("input: %s   expected: %s   actual: %s", tc.CL, e, a)
			}
		})
	}
}

var splitLiteralTests = []struct {
	Name     string
	CL       string
	Expected []string
}{
	{"quoted-space", `"a b c" d e`, []string{`"a b c"`, `d`, `e`}},
	{"escaped-quote", `"ab\"c" "\\" d`, []string{`"ab\"c"`, `"\\"`, `d`}},
	{"unescaped-backslash", `a\\\b d"e f"g h`, []string{`a\\\b`, `d"e f"g`, `h`}},
	{"escaped-backslash-odd", `a\\\"b c d`, []string{`a\\\"b`, `c`, `d`}},
	{"escaped-backslash-even", `a\\\\"b c" d e`, []string{`a\\\\"b c"`, `d`, `e`}},
	{"empty-arg-1", `a "" c`, []string{`a`, `""`, `c`}},
	{"empty-arg-2", `a "" c ""`, []string{`a`, `""`, `c`, `""`}},
	{"empty-arg-3", `a b "" ""`, []string{`a`, `b`, `""`, `""`}},
	{"uneven-quote-1", `a"bc`, []string{`a"bc`}},
	{"uneven-quote-2", `a"b c d`, []string{`a"b c d`}},
	{"uneven-quote-3", `ab c d"`, []string{`ab`, `c`, `d"`}},
	{"uneven-quote-4", `a "b c d`, []string{`a`, `"b c d`}},
	{"command-echo", `cmd /C echo test`, []string{`cmd`, `/C`, `echo`, `test`}},
}

func TestSplitLiteral(t *testing.T) {
	for _, tc := range splitLiteralTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			args := cmdlinewindows.SplitLiteral(tc.CL)
			if !argsEqual(args, tc.Expected) {
				e := printArgs(tc.Expected)
				a := printArgs(args)
				t.Errorf("input: %s   expected: %s   actual: %s", tc.CL, e, a)
			}
		})
	}
}

func argsEqual(a []string, b []string) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func printArgs(args []string) string {
	return strings.TrimPrefix(fmt.Sprintf("%#v", args), "[]string")
}
