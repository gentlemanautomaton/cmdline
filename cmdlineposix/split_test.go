package cmdlineposix_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gentlemanautomaton/cmdline/cmdlineposix"
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
}

func TestSplitCommand(t *testing.T) {
	for _, tc := range splitCommandTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			name, args := cmdlineposix.SplitCommand(tc.CL)
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
	{"simple", `abcd`, []string{`abcd`}},
	{"simple-multi-1", `ab cd`, []string{`ab`, `cd`}},
	{"simple-multi-2", `a b cd efg h`, []string{`a`, `b`, `cd`, `efg`, `h`}},
	{"escaped-backslash", `ab\\cd`, []string{`ab\cd`}},
	{"escaped-single-quote", `ab\'cd`, []string{`ab'cd`}},
	{"escaped-double-quote", `ab\"cd`, []string{`ab"cd`}},
	{"escaped-space", `ab\ cd`, []string{`ab cd`}},
	{"escaped-tab", "ab\\\tcd", []string{"ab\tcd"}},
	{"single-arg-quoted", `a"bc"d`, []string{`abcd`}},
	{"single-quote", `'ab cd'`, []string{`ab cd`}},
	{"single-quote-multi", `ab 'c d' ef`, []string{`ab`, `c d`, `ef`}},
	{"single-quote-inline", `ab'c d'ef`, []string{`abc def`}},
	{"double-quote", `"ab cd"`, []string{`ab cd`}},
	{"double-quote-multi", `ab "c d" ef`, []string{`ab`, `c d`, `ef`}},
	{"double-quote-inline", `ab"c d"ef`, []string{`abc def`}},
	{"nested-double-quote", `"ab\"cd\"ef"`, []string{`ab"cd"ef`}},
	{"single-quoted-backslash", `'ab\cd\ef'`, []string{`ab\cd\ef`}},
	{"single-quoted-double-quote", `'ab"cd"ef'`, []string{`ab"cd"ef`}},
	{"single-quoted-backslash-double-quote", `'ab\"cd\"ef'`, []string{`ab\"cd\"ef`}},
	{"single-quoted-backslash-back-quote", "'ab\\`cd\\`ef'", []string{"ab\\`cd\\`ef"}},
	{"single-quoted-backslash-dollar-sign", `'ab\$cd\$ef'`, []string{`ab\$cd\$ef`}},
	{"double-quoted-backslash", `"ab\cd\ef"`, []string{`ab\cd\ef`}},
	{"double-quoted-single-quote", `"ab'cd'ef"`, []string{`ab'cd'ef`}},
	{"double-quoted-backslash-single-quote", `"ab\'cd\'ef"`, []string{`ab\'cd\'ef`}},
	{"double-quoted-backslash-back-quote", "\"ab\\`cd\\`ef\"", []string{"ab`cd`ef"}},
	{"double-quoted-backslash-dollar-sign", `"ab\$cd\$ef"`, []string{`ab$cd$ef`}},
	{"empty-arg-1", `a "" c`, []string{`a`, ``, `c`}},
	{"empty-arg-2", `a "" c ""`, []string{`a`, ``, `c`, ``}},
	{"empty-arg-3", `a b "" ""`, []string{`a`, `b`, ``, ``}},
	{"uneven-quote-1", `a"bc`, []string{`abc`}},
	{"uneven-quote-2", `a"b c d`, []string{`ab c d`}},
	{"uneven-quote-3", `ab c d"`, []string{`ab`, `c`, `d`}},
	{"uneven-quote-4", `a "b c d`, []string{`a`, `b c d`}},
	{"trailing-newline", "abcd\n", []string{`abcd`}},
	{"trailing-backslash", `abcd\`, []string{`abcd\`}},
	{"continuing-backslash", "abcd\\\nef", []string{`abcdef`}},
	//{"quoted-space", `"a b c" d e`, []string{`a b c`, `d`, `e`}},
	//{"escaped-backslash-odd", `a\\\"b c d`, []string{`a\"b`, `c`, `d`}},
	//{"escaped-backslash-even", `a\\\\"b c" d e`, []string{`a\\b c`, `d`, `e`}},
}

func TestSplit(t *testing.T) {
	for _, tc := range splitTests {
		tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			args := cmdlineposix.Split(tc.CL)
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
